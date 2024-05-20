package commands

import (
	"fmt"
	"os"

	"github.com/hathora/ci/internal/sdk"
	"github.com/hathora/ci/internal/sdk/models/shared"
	"github.com/hathora/ci/internal/setup"
	"github.com/hathora/ci/internal/shorthand"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

var (
	allowedTransportTypes = []string{"tcp", "udp", "tls"}
	maxRoomsPerProcess    = 10000
	maxPort               = 65535
)

var Deployment = &cli.Command{
	Name:  "deployment",
	Usage: "options for deployments",
	Subcommands: []*cli.Command{
		{
			Name:    infoCommandName,
			Aliases: []string{"get-deployment-info"},
			Usage:   "get a deployment by id",
			Flags:   subcommandFlags(deploymentIDFlag),
			Action: func(cCtx *cli.Context) error {
				zap.L().Debug("getting deployment info...")
				deployment := OneDeploymentConfigFrom(cCtx)

				res, err := deployment.SDK.DeploymentV2.GetDeploymentInfo(
					deployment.Context,
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
			Action: func(cCtx *cli.Context) error {
				zap.L().Debug("getting the latest deployment...")
				deployment := DeploymentConfigFrom(cCtx)

				res, err := deployment.SDK.DeploymentV2.GetLatestDeployment(deployment.Context, deployment.AppID)
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
			Action: func(cCtx *cli.Context) error {
				zap.L().Debug("getting all deployments...")
				deployment := DeploymentConfigFrom(cCtx)

				res, err := deployment.SDK.DeploymentV2.GetDeployments(deployment.Context, deployment.AppID)
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
			Action: func(cCtx *cli.Context) error {
				zap.L().Debug("creating a deployment...")
				deployment := DeploymentConfigFrom(cCtx)
				buildID := cCtx.Int(buildIDFlag.Name)
				idleTimeoutEnabled := cCtx.Bool(idleTimeoutFlag.Name)
				roomsPerProcess := cCtx.Int(roomsPerProcessFlag.Name)
				transportType := shared.TransportType(cCtx.String(transportTypeFlag.Name))
				containerPort := cCtx.Int(containerPortFlag.Name)
				requestedMemory := cCtx.Float64(requestedMemoryFlag.Name)
				requestedCPU := cCtx.Float64(requestedCPUFlag.Name)
				addlPorts := cCtx.StringSlice(additionalContainerPortsFlag.Name)
				envVars := cCtx.StringSlice(envVarsFlag.Name)

				additionalContainerPorts, err := parseContainerPorts(addlPorts)
				if err != nil {
					return fmt.Errorf("invalid additional container ports: %w", err)
				}

				env, err := parseEnvVars(envVars)
				if err != nil {
					return fmt.Errorf("invalid environment variables: %w", err)
				}

				res, err := deployment.SDK.DeploymentV2.CreateDeployment(
					deployment.Context,
					buildID,
					shared.DeploymentConfigV2{
						IdleTimeoutEnabled:       idleTimeoutEnabled,
						RoomsPerProcess:          roomsPerProcess,
						TransportType:            transportType,
						ContainerPort:            containerPort,
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

func deploymentEnvVar(name string) []string {
	return []string{fmt.Sprintf("%s%s", deploymentFlagEnvVarPrefix, name)}
}

var (
	deploymentFlagEnvVarPrefix = fmt.Sprintf("%s%s", globalFlagEnvVarPrefix, "DEPLOYMENT_")

	deploymentIDFlag = &cli.IntFlag{
		Name:     "deployment-id",
		Aliases:  []string{"d"},
		EnvVars:  deploymentEnvVar("ID"),
		Usage:    "the ID of the deployment in Hathora",
		Required: true,
	}

	idleTimeoutFlag = &cli.BoolFlag{
		Name:     "idle-timeout-enabled",
		EnvVars:  deploymentEnvVar("IDLE_TIMEOUT_ENABLED"),
		Usage:    "option to shut down processes that have had no new connections or rooms for five minutes",
		Required: true,
	}

	roomsPerProcessFlag = &cli.IntFlag{
		Name:     "rooms-per-process",
		EnvVars:  deploymentEnvVar("ROOMS_PER_PROCESS"),
		Usage:    "how many rooms can be scheduled in a process",
		Required: true,
		Action: func(ctx *cli.Context, v int) error {
			return requireIntInRange(v, 1, maxRoomsPerProcess, "rooms-per-process")
		},
	}

	transportTypeFlag = &cli.StringFlag{
		Name:     "transport-type",
		EnvVars:  deploymentEnvVar("TRANSPORT_TYPE"),
		Usage:    "the underlying communication protocol to the exposed port",
		Required: true,
		Action: func(ctx *cli.Context, v string) error {
			return requireValidEnumValue(v, allowedTransportTypes, "transport-type")
		},
	}

	containerPortFlag = &cli.IntFlag{
		Name:     "container-port",
		EnvVars:  deploymentEnvVar("CONTAINER_PORT"),
		Usage:    "default server port",
		Required: true,
		Action: func(ctx *cli.Context, v int) error {
			return requireIntInRange(v, 1, maxPort, "container-port")
		},
	}

	additionalContainerPortsFlag = &cli.StringSliceFlag{
		Name:    "additional-container-ports",
		Aliases: []string{"additional-container-port"},
		EnvVars: append(deploymentEnvVar("ADDITIONAL_CONTAINER_PORTS"), deploymentEnvVar("ADDITIONAL_CONTAINER_PORT")...),
		Usage:   "additional server ports",
	}

	envVarsFlag = &cli.StringSliceFlag{
		Name:    "env",
		EnvVars: deploymentEnvVar("ENV"),
		Usage:   "environment variables",
	}

	requestedMemoryFlag = &cli.Float64Flag{
		Name:     "requested-memory-mb",
		EnvVars:  deploymentEnvVar("REQUESTED_MEMORY_MB"),
		Usage:    "the amount of memory allocated to your process in MB",
		Required: true,
	}

	requestedCPUFlag = &cli.Float64Flag{
		Name:     "requested-cpu",
		EnvVars:  deploymentEnvVar("REQUESTED_CPU"),
		Usage:    "the number of cores allocated to your process",
		Required: true,
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
	deploymentConfigKey = struct{}{}
)

type DeploymentConfig struct {
	*GlobalConfig
	SDK   *sdk.SDK
	AppID *string
}

func (c *DeploymentConfig) Load(cCtx *cli.Context) {
	c.GlobalConfig = GlobalConfigFrom(cCtx)
	c.SDK = setup.SDK(c.Token, c.BaseURL, c.Verbosity)
}

func DeploymentConfigFrom(cCtx *cli.Context) *DeploymentConfig {
	return ConfigFromCLI[*DeploymentConfig](deploymentConfigKey, cCtx)
}

var (
	oneDeploymentConfigKey = struct{}{}
)

type OneDeploymentConfig struct {
	*DeploymentConfig
	DeploymentID int
}

func (c *OneDeploymentConfig) Load(cCtx *cli.Context) {
	c.DeploymentConfig = DeploymentConfigFrom(cCtx)
	c.DeploymentID = cCtx.Int(deploymentIDFlag.Name)
}

func OneDeploymentConfigFrom(cCtx *cli.Context) *OneDeploymentConfig {
	return ConfigFromCLI[*OneDeploymentConfig](oneDeploymentConfigKey, cCtx)
}
