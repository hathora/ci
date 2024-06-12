package commands

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/hathora/ci/internal/sdk"
	"github.com/hathora/ci/internal/sdk/models/shared"
	"github.com/hathora/ci/internal/shorthand"
	"github.com/urfave/cli/v3"
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
	Action: func(ctx context.Context, cmd *cli.Command) error {
		deploy, err := DeployConfigFrom(cmd)
		if err != nil {
			return err
		}
		createdBuild, err := doBuildCreate(ctx, deploy.CreateBuildConfig)
		if err != nil {
			return err
		}

		useLatest := cmd.Bool(fromLatestFlag.Name)
		if useLatest {
			res, err := deploy.SDK.DeploymentV2.GetLatestDeployment(ctx, deploy.AppID)
			if err != nil {
				return fmt.Errorf("unable to retrieve latest deployment: %w", err)
			}

			deploy.Merge(res.DeploymentV2)
		}

		if err := deploy.Validate(); err != nil {
			return err
		}

		res, err := deploy.SDK.DeploymentV2.CreateDeployment(
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
	*CreateBuildConfig
	IdleTimeoutEnabled       *bool
	RoomsPerProcess          int
	TransportType            shared.TransportType
	ContainerPort            int
	RequestedMemoryMB        float64
	RequestedCPU             float64
	AdditionalContainerPorts []shared.ContainerPort
	Env                      []shared.DeploymentConfigV2Env
}

var _ LoadableConfig = (*DeployConfig)(nil)

func (c *DeployConfig) Load(cmd *cli.Command) error {
	build, err := CreateBuildConfigFrom(cmd)
	if err != nil {
		return err
	}
	c.CreateBuildConfig = build

	c.RoomsPerProcess = int(cmd.Int(roomsPerProcessFlag.Name))
	c.TransportType = shared.TransportType(cmd.String(transportTypeFlag.Name))
	c.ContainerPort = int(cmd.Int(containerPortFlag.Name))
	c.RequestedMemoryMB = cmd.Float(requestedMemoryFlag.Name)
	c.RequestedCPU = cmd.Float(requestedCPUFlag.Name)
	c.IdleTimeoutEnabled = sdk.Bool(cmd.Bool(idleTimeoutFlag.Name))

	addlPorts := cmd.StringSlice(additionalContainerPortsFlag.Name)
	parsedAddlPorts, err := parseContainerPorts(addlPorts)
	if err != nil {
		return fmt.Errorf("invalid additional container ports: %w", err)
	}
	c.AdditionalContainerPorts = parsedAddlPorts

	envVars := cmd.StringSlice(envVarsFlag.Name)
	env, err := parseEnvVars(envVars)
	if err != nil {
		return fmt.Errorf("invalid environment variables: %w", err)
	}
	c.Env = env

	return nil
}

func (c *DeployConfig) Merge(latest *shared.DeploymentV2) {
	if latest == nil {
		return
	}

	if c.IdleTimeoutEnabled == nil {
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

	if c.IdleTimeoutEnabled == nil {
		err = errors.Join(err, fmt.Errorf("idle timeout enabled is required"))
	}

	if c.RoomsPerProcess == 0 {
		err = errors.Join(err, fmt.Errorf("rooms per process is required"))
	}

	err = errors.Join(err, requireIntInRange(c.RoomsPerProcess, minRoomsPerProcess, maxRoomsPerProcess, roomsPerProcessFlag.Name))

	if c.TransportType == "" {
		err = errors.Join(err, fmt.Errorf("transport type is required"))
	}
	err = errors.Join(err, requireValidEnumValue(c.TransportType, allowedTransportTypes, transportTypeFlag.Name))

	if c.ContainerPort == 0 {
		err = errors.Join(err, fmt.Errorf("container port is required"))
	}
	err = errors.Join(err, requireIntInRange(c.ContainerPort, minPort, maxPort, containerPortFlag.Name))

	if c.RequestedMemoryMB == 0 {
		err = errors.Join(err, fmt.Errorf("requested memory is required"))
	}
	err = errors.Join(err, requireFloatInRange(c.RequestedMemoryMB, minMemoryMB, maxMemoryMB, requestedMemoryFlag.Name))
	if c.RequestedCPU == 0 {
		err = errors.Join(err, fmt.Errorf("requested CPU is required"))
	}

	err = errors.Join(err, requireFloatInRange(c.RequestedCPU, minCPU, maxCPU, requestedCPUFlag.Name))
	err = errors.Join(err, requireMaxDecimals(c.RequestedCPU, maxCPUDecimalPlaces, requestedCPUFlag.Name))

	if c.RequestedMemoryMB != (c.RequestedCPU * memoryMBPerCPU) {
		err = errors.Join(err,
			fmt.Errorf("invalid memory: %f and cpu: %f requested-memory-mb must be a %f:1 ratio to requested-cpu",
				c.RequestedMemoryMB,
				c.RequestedCPU,
				memoryMBPerCPU,
			))
	}

	return err
}

func (c *DeployConfig) New() LoadableConfig {
	return &DeployConfig{}
}

func DeployConfigFrom(cmd *cli.Command) (*DeployConfig, error) {
	return ConfigFromCLI[*DeployConfig](deployConfigKey, cmd)
}
