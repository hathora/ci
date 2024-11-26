package commands

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
	"go.uber.org/zap"

	"github.com/hathora/ci/internal/sdk/models/shared"
	"github.com/hathora/ci/internal/shorthand"
)

func Deploy() *cli.Command {
	return &cli.Command{
		Name:  "deploy",
		Usage: "create a build and a deployment in a combined flow",
		Flags: subcommandFlags(
			buildIDFlag(),
			buildTagFlag(),
			fileFlag(),
			hideUploadProgressFlag(),
			fromLatestFlag(),
			roomsPerProcessFlag(),
			transportTypeFlag(),
			containerPortFlag(),
			requestedMemoryFlag(),
			requestedCPUFlag(),
			additionalContainerPortsFlag(),
			envVarsFlag(),
			idleTimeoutFlag(),
		),
		UsageText: `hathora deploy [options]`,
		Action: func(ctx context.Context, cmd *cli.Command) error {
			deploy, err := DeployConfigFrom(cmd)
			if err != nil {
				//nolint:errcheck
				cli.ShowSubcommandHelp(cmd)
				return err
			}

			useLatest := cmd.Bool(fromLatestFlagName)
			if useLatest {
				res, err := deploy.SDK.DeploymentsV3.GetLatestDeployment(ctx, deploy.AppID)
				if err != nil {
					return fmt.Errorf("unable to retrieve latest deployment: %w", err)
				}

				deploy.Merge(res.DeploymentV3, cmd.IsSet(idleTimeoutFlagName))
			}

			if err := deploy.Validate(); err != nil {
				//nolint:errcheck
				cli.ShowSubcommandHelp(cmd)
				return err
			}

			createdBuild, err := doBuildCreate(ctx, deploy.SDK, deploy.BuildTag, deploy.BuildID, deploy.FilePath, deploy.HideUploadProgress)
			if err != nil {
				return err
			}

			res, err := deploy.SDK.DeploymentsV3.CreateDeployment(
				ctx,
				shared.DeploymentConfigV3{
					BuildID:                  createdBuild.BuildID,
					IdleTimeoutEnabled:       *deploy.IdleTimeoutEnabled,
					RoomsPerProcess:          deploy.RoomsPerProcess,
					TransportType:            deploy.TransportType,
					ContainerPort:            deploy.ContainerPort,
					RequestedMemoryMB:        deploy.RequestedMemoryMB,
					RequestedCPU:             deploy.RequestedCPU,
					AdditionalContainerPorts: deploy.AdditionalContainerPorts,
					Env:                      deploy.Env,
				},
				deploy.AppID,
			)
			if err != nil {
				return fmt.Errorf("failed to create a deployment: %w", err)
			}

			return deploy.Output.Write(res.DeploymentV3, os.Stdout)
		},
	}
}

var (
	deployConfigKey = "commands.Deploy.DI"
)

type DeployConfig struct {
	*CreateDeploymentConfig
	BuildTag           string
	FilePath           string
	HideUploadProgress bool
}

var _ LoadableConfig = (*DeployConfig)(nil)

func (c *DeployConfig) Load(cmd *cli.Command) error {
	deployment, err := CreateDeploymentConfigFrom(cmd)
	if err != nil {
		return err
	}
	c.CreateDeploymentConfig = deployment

	c.BuildTag = cmd.String(buildTagFlagName)
	c.BuildID = cmd.String(buildIDFlagName)
	c.FilePath = cmd.String(fileFlagName)
	c.HideUploadProgress = cmd.Bool(hideUploadProgressFlagName)
	c.Log = c.Log.With(zap.String("build.tag", c.BuildTag)).With(zap.String("build.id", c.BuildID))

	return nil
}

func (c *DeployConfig) Merge(latest *shared.DeploymentV3, isIdleTimeoutDefault bool) {
	if latest == nil {
		return
	}

	if !isIdleTimeoutDefault {
		c.IdleTimeoutEnabled = &latest.IdleTimeoutEnabled
	}

	if c.RoomsPerProcess == 0 {
		c.RoomsPerProcess = latest.RoomsPerProcess
	}

	if c.TransportType == "" {
		c.TransportType = latest.DefaultContainerPort.TransportType
	}

	if c.ContainerPort == 0 {
		c.ContainerPort = latest.DefaultContainerPort.Port
	}

	if c.RequestedMemoryMB == 0 {
		c.RequestedMemoryMB = latest.RequestedMemoryMB
	}

	if c.RequestedCPU == 0 {
		c.RequestedCPU = latest.RequestedCPU
	}

	if len(c.AdditionalContainerPorts) == 0 {
		c.AdditionalContainerPorts = latest.AdditionalContainerPorts
	}

	if len(c.Env) == 0 {
		c.Env = shorthand.MapEnvToEnvConfig(latest.Env)
	}
}

func (c *DeployConfig) Validate() error {
	var err error

	if c.AppID == nil || *c.AppID == "" {
		err = errors.Join(err, missingRequiredFlag(appIDFlagName))
	}

	if c.RoomsPerProcess == 0 {
		err = errors.Join(err, missingRequiredFlag(roomsPerProcessFlagName))
	}

	err = errors.Join(err, requireIntInRange(c.RoomsPerProcess, minRoomsPerProcess, maxRoomsPerProcess, roomsPerProcessFlagName))

	if c.TransportType == "" {
		err = errors.Join(err, missingRequiredFlag(transportTypeFlagName))
	}
	err = errors.Join(err, requireValidEnumValue(c.TransportType, allowedTransportTypes, transportTypeFlagName))

	if c.ContainerPort == 0 {
		err = errors.Join(err, missingRequiredFlag(containerPortFlagName))
	}
	err = errors.Join(err, requireIntInRange(c.ContainerPort, minPort, maxPort, containerPortFlagName))

	if c.RequestedMemoryMB == 0 {
		err = errors.Join(err, missingRequiredFlag(requestedMemoryFlagName))
	}
	err = errors.Join(err, requireFloatInRange(c.RequestedMemoryMB, minMemoryMB, maxMemoryMB, requestedMemoryFlagName))
	if c.RequestedCPU == 0 {
		err = errors.Join(err, missingRequiredFlag(requestedCPUFlagName))
	}

	err = errors.Join(err, requireFloatInRange(c.RequestedCPU, minCPU, maxCPU, requestedCPUFlagName))
	err = errors.Join(err, requireMaxDecimals(c.RequestedCPU, maxCPUDecimalPlaces, requestedCPUFlagName))

	return err
}

func (c *DeployConfig) New() LoadableConfig {
	return &DeployConfig{}
}

func DeployConfigFrom(cmd *cli.Command) (*DeployConfig, error) {
	return ConfigFromCLI[*DeployConfig](deployConfigKey, cmd)
}
