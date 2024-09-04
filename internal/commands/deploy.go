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

var Deploy = &cli.Command{
	Name:  "deploy",
	Usage: "create a build and a deployment in a combined flow",
	Flags: subcommandFlags(
		buildTagFlag,
		fileFlag,
		fromLatestFlag,
		roomsPerProcessFlag,
		transportTypeFlag,
		containerPortFlag,
		requestedMemoryFlag,
		requestedCPUFlag,
		additionalContainerPortsFlag,
		envVarsFlag,
		idleTimeoutFlag,
	),
	UsageText: `hathora deploy [options]`,
	Action: func(ctx context.Context, cmd *cli.Command) error {
		deploy, err := DeployConfigFrom(cmd)
		if err != nil {
			//nolint:errcheck
			cli.ShowSubcommandHelp(cmd)
			return err
		}

		useLatest := cmd.Bool(fromLatestFlag.Name)
		if useLatest {
			res, err := deploy.SDK.DeploymentsV3.GetLatestDeployment(ctx, deploy.AppID)
			if err != nil {
				return fmt.Errorf("unable to retrieve latest deployment: %w", err)
			}

			deploy.Merge(res.DeploymentV3, cmd.IsSet(idleTimeoutFlag.Name))
		}

		if err := deploy.Validate(); err != nil {
			//nolint:errcheck
			cli.ShowSubcommandHelp(cmd)
			return err
		}

		createdBuild, err := doBuildCreate(ctx, deploy.SDK, deploy.AppID, deploy.BuildTag, deploy.FilePath)
		if err != nil {
			return err
		}

		res, err := deploy.SDK.DeploymentsV2.CreateDeploymentV3(
			ctx,
			createdBuild.BuildID,
			shared.DeploymentConfigV2{
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

		return deploy.Output.Write(res.DeploymentV2, os.Stdout)
	},
}

var (
	deployConfigKey = "commands.Deploy.DI"
)

type DeployConfig struct {
	*CreateDeploymentConfig
	BuildTag string
	FilePath string
}

var _ LoadableConfig = (*DeployConfig)(nil)

func (c *DeployConfig) Load(cmd *cli.Command) error {
	deployment, err := CreateDeploymentConfigFrom(cmd)
	if err != nil {
		return err
	}
	c.CreateDeploymentConfig = deployment

	c.BuildTag = cmd.String(buildTagFlag.Name)
	c.FilePath = cmd.String(fileFlag.Name)
	c.Log = c.Log.With(zap.String("build.tag", c.BuildTag))

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

	if c.RoomsPerProcess == 0 {
		err = errors.Join(err, missingRequiredFlag(roomsPerProcessFlag.Name))
	}

	err = errors.Join(err, requireIntInRange(c.RoomsPerProcess, minRoomsPerProcess, maxRoomsPerProcess, roomsPerProcessFlag.Name))

	if c.TransportType == "" {
		err = errors.Join(err, missingRequiredFlag(transportTypeFlag.Name))
	}
	err = errors.Join(err, requireValidEnumValue(c.TransportType, allowedTransportTypes, transportTypeFlag.Name))

	if c.ContainerPort == 0 {
		err = errors.Join(err, missingRequiredFlag(containerPortFlag.Name))
	}
	err = errors.Join(err, requireIntInRange(c.ContainerPort, minPort, maxPort, containerPortFlag.Name))

	if c.RequestedMemoryMB == 0 {
		err = errors.Join(err, missingRequiredFlag(requestedMemoryFlag.Name))
	}
	err = errors.Join(err, requireFloatInRange(c.RequestedMemoryMB, minMemoryMB, maxMemoryMB, requestedMemoryFlag.Name))
	if c.RequestedCPU == 0 {
		err = errors.Join(err, missingRequiredFlag(requestedCPUFlag.Name))
	}

	err = errors.Join(err, requireFloatInRange(c.RequestedCPU, minCPU, maxCPU, requestedCPUFlag.Name))
	err = errors.Join(err, requireMaxDecimals(c.RequestedCPU, maxCPUDecimalPlaces, requestedCPUFlag.Name))

	return err
}

func (c *DeployConfig) New() LoadableConfig {
	return &DeployConfig{}
}

func DeployConfigFrom(cmd *cli.Command) (*DeployConfig, error) {
	return ConfigFromCLI[*DeployConfig](deployConfigKey, cmd)
}
