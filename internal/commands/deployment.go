package commands

import (
	"context"
	"fmt"
	"github.com/hathora/ci/internal/commands/altsrc"
	"os"
	"strconv"

	"github.com/hathora/ci/internal/sdk"
	"github.com/hathora/ci/internal/sdk/models/shared"
	"github.com/hathora/ci/internal/setup"
	"github.com/hathora/ci/internal/shorthand"
	"github.com/urfave/cli/v3"
	"go.uber.org/zap"
)

var (
	allowedTransportTypes = []string{"tcp", "udp", "tls"}
	maxRoomsPerProcess    = int64(10000)
	maxPort               = int64(65535)
)

var Deployment = &cli.Command{
	Name:  "deployment",
	Usage: "options for deployments",
	Commands: []*cli.Command{
		{
			Name:    infoCommandName,
			Aliases: []string{"get-deployment-info"},
			Usage:   "get a deployment by id",
			Flags:   subcommandFlags(deploymentIDFlag),
			Action: func(ctx context.Context, cmd *cli.Command) error {
				deployment, err := OneDeploymentConfigFrom(cmd)
				if err != nil {
					return err
				}
				deployment.Log.Debug("getting deployment info...")

				res, err := deployment.SDK.DeploymentV2.GetDeploymentInfo(
					ctx,
					deployment.DeploymentID,
					deployment.AppID,
				)
				if err != nil {
					return fmt.Errorf("failed to get deployment info: %w", err)
				}

				return deployment.Output.Write(res.DeploymentV2, os.Stdout)
			},
		},
		{
			Name:    latestCommandName,
			Aliases: []string{"get-latest-deployment"},
			Usage:   "get the latest deployment",
			Flags:   subcommandFlags(),
			Action: func(ctx context.Context, cmd *cli.Command) error {
				deployment, err := DeploymentConfigFrom(cmd)
				if err != nil {
					return err
				}
				deployment.Log.Debug("getting the latest deployment...")

				res, err := deployment.SDK.DeploymentV2.GetLatestDeployment(ctx, deployment.AppID)
				if err != nil {
					return fmt.Errorf("failed to get the latest deployment: %w", err)
				}

				return deployment.Output.Write(res.DeploymentV2, os.Stdout)
			},
		},
		{
			Name:    listCommandName,
			Aliases: []string{"get-deployments"},
			Usage:   "get all deployments",
			Flags:   subcommandFlags(),
			Action: func(ctx context.Context, cmd *cli.Command) error {
				deployment, err := DeploymentConfigFrom(cmd)
				if err != nil {
					return err
				}
				deployment.Log.Debug("getting all deployments...")

				res, err := deployment.SDK.DeploymentV2.GetDeployments(ctx, deployment.AppID)
				if err != nil {
					return fmt.Errorf("failed to get deployments: %w", err)
				}

				if len(res.DeploymentV2s) == 0 {
					return fmt.Errorf("no deployments found")
				}

				return deployment.Output.Write(res.DeploymentV2s, os.Stdout)
			},
		},
		{
			Name:    createCommandName,
			Aliases: []string{"create-deployment"},
			Usage:   "create a deployment",
			Flags: subcommandFlags(
				buildIDFlag,
				idleTimeoutFlag,
				roomsPerProcessFlag,
				transportTypeFlag,
				containerPortFlag,
				requestedMemoryFlag,
				requestedCPUFlag,
				additionalContainerPortsFlag,
				envVarsFlag,
			),
			Action: func(ctx context.Context, cmd *cli.Command) error {
				deployment, err := DeploymentConfigFrom(cmd)
				if err != nil {
					return err
				}

				buildID := cmd.Int(buildIDFlag.Name)
				deployment.Log = deployment.Log.With(zap.Int64("build.id", buildID))
				deployment.Log.Debug("creating a deployment...")

				idleTimeoutEnabled := cmd.Bool(idleTimeoutFlag.Name)
				roomsPerProcess := cmd.Int(roomsPerProcessFlag.Name)
				transportType := shared.TransportType(cmd.String(transportTypeFlag.Name))
				containerPort := cmd.Int(containerPortFlag.Name)
				requestedMemory := cmd.Float(requestedMemoryFlag.Name)
				requestedCPU := cmd.Float(requestedCPUFlag.Name)
				addlPorts := cmd.StringSlice(additionalContainerPortsFlag.Name)
				envVars := cmd.StringSlice(envVarsFlag.Name)

				if requestedMemory != (requestedCPU * 2048) {
					return fmt.Errorf("invalid memory: %s and cpu: %s requested-memory-mb must be a 2048:1 ratio to requested-cpu",
						strconv.FormatFloat(requestedMemory, 'f', -1, 64),
						strconv.FormatFloat(requestedCPU, 'f', -1, 64))
				}

				additionalContainerPorts, err := parseContainerPorts(addlPorts)
				if err != nil {
					return fmt.Errorf("invalid additional container ports: %w", err)
				}

				env, err := parseEnvVars(envVars)
				if err != nil {
					return fmt.Errorf("invalid environment variables: %w", err)
				}

				res, err := deployment.SDK.DeploymentV2.CreateDeployment(
					ctx,
					int(buildID),
					shared.DeploymentConfigV2{
						IdleTimeoutEnabled:       idleTimeoutEnabled,
						RoomsPerProcess:          int(roomsPerProcess),
						TransportType:            transportType,
						ContainerPort:            int(containerPort),
						RequestedMemoryMB:        requestedMemory,
						RequestedCPU:             requestedCPU,
						AdditionalContainerPorts: additionalContainerPorts,
						Env:                      env,
					},
					deployment.AppID,
				)
				if err != nil {
					return fmt.Errorf("failed to create a deployment: %w", err)
				}

				return deployment.Output.Write(res.DeploymentV2, os.Stdout)
			},
		},
	},
}

func deploymentEnvVar(name string) string {
	return fmt.Sprintf("%s%s", deploymentFlagEnvVarPrefix, name)
}

var (
	deploymentFlagEnvVarPrefix = fmt.Sprintf("%s%s", globalFlagEnvVarPrefix, "DEPLOYMENT_")

	deploymentIDFlag = &cli.IntFlag{
		Name:     "deployment-id",
		Aliases:  []string{"d"},
		Sources:  cli.EnvVars(deploymentEnvVar("ID")),
		Usage:    "the ID of the deployment in Hathora",
		Required: true,
	}

	idleTimeoutFlag = &cli.BoolFlag{
		Name: "idle-timeout-enabled",
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(deploymentEnvVar("IDLE_TIMEOUT_ENABLED")),
			altsrc.File(configFlag.Name, "deployment.idle-timeout-enabled"),
		),
		Usage:    "option to shut down processes that have had no new connections or rooms for five minutes",
		Persistent: true,
	}

	roomsPerProcessFlag = &cli.IntFlag{
		Name: "rooms-per-process",
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(deploymentEnvVar("ROOMS_PER_PROCESS")),
			altsrc.File(configFlag.Name, "deployment.rooms-per-process"),
		),
		Usage:    "how many rooms can be scheduled in a process",
		Persistent: true,
		Action: func(ctx context.Context, cmd *cli.Command, v int64) error {
			return requireIntInRange(v, 1, maxRoomsPerProcess, "rooms-per-process")
		},
	}

	transportTypeFlag = &cli.StringFlag{
		Name: "transport-type",
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(deploymentEnvVar("TRANSPORT_TYPE")),
			altsrc.File(configFlag.Name, "deployment.transport-type"),
		),
		Usage:    "the underlying communication protocol to the exposed port",
		Persistent: true,
		Action: func(ctx context.Context, cmd *cli.Command, v string) error {
			return requireValidEnumValue(v, allowedTransportTypes, "transport-type")
		},
	}

	containerPortFlag = &cli.IntFlag{
		Name: "container-port",
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(deploymentEnvVar("CONTAINER_PORT")),
			altsrc.File(configFlag.Name, "deployment.container-port"),
		),
		Usage:    "default server port",
		Persistent: true,
		Action: func(ctx context.Context, cmd *cli.Command, v int64) error {
			return requireIntInRange(v, 1, maxPort, "container-port")
		},
	}

	additionalContainerPortsFlag = &cli.StringSliceFlag{
		Name:    "additional-container-ports",
		Aliases: []string{"additional-container-port"},
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(deploymentEnvVar("ADDITIONAL_CONTAINER_PORTS")),
			altsrc.File(configFlag.Name, "deployment.additional-container-ports"),
		),
		Usage: "additional server ports",
	}

	envVarsFlag = &cli.StringSliceFlag{
		Name:    "env",
		Sources: cli.EnvVars(deploymentEnvVar("ENV")),
		Usage:   "environment variables",
	}

	requestedMemoryFlag = &cli.FloatFlag{
		Name: "requested-memory-mb",
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(deploymentEnvVar("REQUESTED_MEMORY_MB")),
			altsrc.File(configFlag.Name, "deployment.requested-memory-mb"),
		),
		Usage:    "the amount of memory allocated to your process in MB",
		Persistent: true,
		Action: func(ctx context.Context, cmd *cli.Command, v float64) error {
			return requireFloatInRange(v, 1024, 8192, "requested-memory-mb")
		},
	}

	requestedCPUFlag = &cli.FloatFlag{
		Name: "requested-cpu",
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(deploymentEnvVar("REQUESTED_CPU")),
			altsrc.File(configFlag.Name, "deployment.requested-cpu"),
		),
		Usage:    "the number of cores allocated to your process",
		Persistent: true,
		Action: func(ctx context.Context, cmd *cli.Command, v float64) error {
			rangeErr := requireFloatInRange(v, 0.5, 4, "requested-cpu")
			if rangeErr != nil {
				return rangeErr
			}
			decimalErr := requireMaxDecimals(v, 1, "requested-cpu")
			if decimalErr != nil {
				return decimalErr
			}

			return nil
		},
	}
)

func parseContainerPorts(ports []string) ([]shared.ContainerPort, error) {
	output := make([]shared.ContainerPort, 0, len(ports))
	for _, port := range ports {
		p, err := shorthand.ParseContainerPort(port)
		if err != nil {
			return nil, err
		}
		output = append(output, *p)
	}
	return output, nil
}

func parseEnvVars(envVars []string) ([]shared.DeploymentConfigV2Env, error) {
	output := make([]shared.DeploymentConfigV2Env, 0, len(envVars))
	for _, envVar := range envVars {
		env, err := shorthand.ParseDeploymentEnvVar(envVar)
		if err != nil {
			return nil, err
		}
		output = append(output, *env)
	}
	return output, nil
}

var (
	deploymentConfigKey = "commands.DeploymentConfig.DI"
)

type DeploymentConfig struct {
	*GlobalConfig
	SDK *sdk.SDK
}

var _ LoadableConfig = (*DeploymentConfig)(nil)

func (c *DeploymentConfig) Load(cmd *cli.Command) error {
	global, err := GlobalConfigFrom(cmd)
	if err != nil {
		return err
	}
	c.GlobalConfig = global
	c.SDK = setup.SDK(c.Token, c.BaseURL, c.Verbosity)
	return nil
}

func (c *DeploymentConfig) New() LoadableConfig {
	return &DeploymentConfig{}
}

func DeploymentConfigFrom(cmd *cli.Command) (*DeploymentConfig, error) {
	return ConfigFromCLI[*DeploymentConfig](deploymentConfigKey, cmd)
}

var (
	oneDeploymentConfigKey = "commands.OneDeploymentConfig.DI"
)

type OneDeploymentConfig struct {
	*DeploymentConfig
	DeploymentID int
}

var _ LoadableConfig = (*OneDeploymentConfig)(nil)

func (c *OneDeploymentConfig) Load(cmd *cli.Command) error {
	deployment, err := DeploymentConfigFrom(cmd)
	if err != nil {
		return err
	}
	c.DeploymentConfig = deployment
	c.DeploymentID = int(cmd.Int(deploymentIDFlag.Name))
	c.Log = c.Log.With(zap.Int("deployment.id", c.DeploymentID))
	return nil
}

func (c *OneDeploymentConfig) New() LoadableConfig {
	return &OneDeploymentConfig{}
}

func OneDeploymentConfigFrom(cmd *cli.Command) (*OneDeploymentConfig, error) {
	return ConfigFromCLI[*OneDeploymentConfig](oneDeploymentConfigKey, cmd)
}
