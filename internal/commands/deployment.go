package commands

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/urfave/cli/v3"
	"go.uber.org/zap"

	"github.com/hathora/ci/internal/commands/altsrc"
	"github.com/hathora/ci/internal/output"
	"github.com/hathora/ci/internal/setup"
	"github.com/hathora/ci/internal/shorthand"
	"github.com/hathora/ci/internal/workaround"
	sdk "github.com/hathora/cloud-sdk-go/hathoracloud"
	"github.com/hathora/cloud-sdk-go/hathoracloud/models/components"
)

const deploymentFlagEnvVarPrefix = globalFlagEnvVarPrefix + "DEPLOYMENT_"

var (
	allowedTransportTypes = []string{"tcp", "udp", "tls"}
	minRoomsPerProcess    = 1
	maxRoomsPerProcess    = 10000
	minPort               = 1
	maxPort               = 65535
	minCPU                = 0.25
	maxCPUDecimalPlaces   = 2
	maxCPU                = float64(8)
	minMemoryMB           = float64(512)
	maxMemoryMB           = float64(16384)
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

				if deployment.AppID == nil || *deployment.AppID == "" {
					err = errors.Join(err, missingRequiredFlag(appIDFlag.Name))
				}

				if err != nil {
					//nolint:errcheck
					cli.ShowSubcommandHelp(cmd)
					return err
				}
				deployment.Log.Debug("getting deployment info...")

				res, err := deployment.SDK.DeploymentsV3.GetDeployment(
					ctx,
					deployment.DeploymentID,
					deployment.AppID,
				)
				if err != nil {
					return fmt.Errorf("failed to get deployment info: %w", err)
				}

				return deployment.Output.Write(res, os.Stdout)
			},
		},
		{
			Name:    latestCommandName,
			Aliases: []string{"get-latest-deployment"},
			Usage:   "get the latest deployment",
			Flags:   subcommandFlags(),
			Action: func(ctx context.Context, cmd *cli.Command) error {
				deployment, err := DeploymentConfigFrom(cmd)
				if deployment.AppID == nil || *deployment.AppID == "" {
					err = errors.Join(err, missingRequiredFlag(appIDFlag.Name))
				}
				if err != nil {
					//nolint:errcheck
					cli.ShowSubcommandHelp(cmd)
					return err
				}
				deployment.Log.Debug("getting the latest deployment...")

				res, err := deployment.SDK.DeploymentsV3.GetLatestDeployment(ctx, deployment.AppID)
				if err != nil {
					return fmt.Errorf("failed to get the latest deployment: %w", err)
				}

				return deployment.Output.Write(res, os.Stdout)
			},
		},
		{
			Name:    listCommandName,
			Aliases: []string{"get-deployments"},
			Usage:   "get all deployments",
			Flags:   subcommandFlags(),
			Action: func(ctx context.Context, cmd *cli.Command) error {
				deployment, err := DeploymentConfigFrom(cmd)
				if deployment.AppID == nil || *deployment.AppID == "" {
					err = errors.Join(err, missingRequiredFlag(appIDFlag.Name))
				}
				if err != nil {
					//nolint:errcheck
					cli.ShowSubcommandHelp(cmd)
					return err
				}
				deployment.Log.Debug("getting all deployments...")

				res, err := deployment.SDK.DeploymentsV3.GetDeployments(ctx, deployment.AppID, nil)
				if err != nil {
					return fmt.Errorf("failed to get deployments: %w", err)
				}

				if len(res.Deployments) == 0 {
					return fmt.Errorf("no deployments found")
				}

				return deployment.Output.Write(res.Deployments, os.Stdout)
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
				fromLatestFlag,
				deploymentTagFlag,
			),
			Action: func(ctx context.Context, cmd *cli.Command) error {
				zap.L().Debug("creating a deployment...")
				deployment, err := CreateDeploymentConfigFrom(cmd)
				if err != nil {
					//nolint:errcheck
					cli.ShowSubcommandHelp(cmd)
					return err
				}

				useLatest := cmd.Bool(fromLatestFlag.Name)
				if useLatest {
					res, err := deployment.SDK.DeploymentsV3.GetLatestDeployment(ctx, deployment.AppID)
					if err != nil {
						return fmt.Errorf("unable to retrieve latest deployment: %w", err)
					}

					deployment.Merge(res)
				}

				if err := deployment.Validate(); err != nil {
					//nolint:errcheck
					cli.ShowSubcommandHelp(cmd)
					return err
				}

				var deploymentTag *string
				if deployment.DeploymentTag != "" {
					deploymentTag = &deployment.DeploymentTag
				}

				res, err := deployment.SDK.DeploymentsV3.CreateDeployment(
					ctx,
					components.DeploymentConfigV3{
						BuildID:                  deployment.BuildID,
						IdleTimeoutEnabled:       *deployment.IdleTimeoutEnabled,
						RoomsPerProcess:          deployment.RoomsPerProcess,
						TransportType:            deployment.TransportType,
						ContainerPort:            deployment.ContainerPort,
						RequestedMemoryMB:        deployment.RequestedMemoryMB,
						RequestedCPU:             deployment.RequestedCPU,
						AdditionalContainerPorts: deployment.AdditionalContainerPorts,
						Env:                      deployment.Env,
						DeploymentTag:            deploymentTag,
					},
					deployment.AppID,
				)
				if err != nil {
					return fmt.Errorf("failed to create a deployment: %w", err)
				}

				return deployment.Output.Write(res, os.Stdout)
			},
		},
	},
}

func deploymentEnvVar(name string) string {
	return fmt.Sprintf("%s%s", deploymentFlagEnvVarPrefix, name)
}

var (
	deploymentIDFlag = &cli.StringFlag{
		Name:     "deployment-id",
		Aliases:  []string{"d"},
		Sources:  cli.EnvVars(deploymentEnvVar("ID")),
		Usage:    "the `<id>` of the deployment in Hathora",
		Required: true,
		Category: "Deployment:",
	}

	idleTimeoutFlag = &cli.BoolFlag{
		Name: "idle-timeout-enabled",
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(deploymentEnvVar("IDLE_TIMEOUT_ENABLED")),
			altsrc.ConfigFile(configFlag.Name, "deployment.idle-timeout-enabled"),
		),
		Usage:      "whether to shut down processes that have had no new connections or rooms for five minutes",
		Persistent: true,
		Category:   "Deployment:",
	}

	roomsPerProcessFlag = &workaround.IntFlag{
		Name: "rooms-per-process",
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(deploymentEnvVar("ROOMS_PER_PROCESS")),
			altsrc.ConfigFile(configFlag.Name, "deployment.rooms-per-process"),
		),
		Usage:      "`<count>` of the rooms that can be scheduled in a process",
		Persistent: true,
		Category:   "Deployment:",
	}

	transportTypeFlag = &cli.StringFlag{
		Name: "transport-type",
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(deploymentEnvVar("TRANSPORT_TYPE")),
			altsrc.ConfigFile(configFlag.Name, "deployment.transport-type"),
		),
		Usage:      "`<protocol>` for the exposed port to use (tcp, udp, tls)",
		Persistent: true,
		Category:   "Deployment:",
	}

	containerPortFlag = &workaround.IntFlag{
		Name: "container-port",
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(deploymentEnvVar("CONTAINER_PORT")),
			altsrc.ConfigFile(configFlag.Name, "deployment.container-port"),
		),
		Usage:      "`<port>` to expose on the deployed container",
		Persistent: true,
		Category:   "Deployment:",
	}

	additionalContainerPortsFlag = &cli.StringSliceFlag{
		Name: "additional-container-ports",
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(deploymentEnvVar("ADDITIONAL_CONTAINER_PORTS")),
			altsrc.ConfigFile(configFlag.Name, "deployment.additional-container-ports"),
		),
		Usage:    "additional `<port>` to expose; format: name:port/protocol",
		Category: "Deployment:",
	}

	envVarsFlag = &cli.StringSliceFlag{
		Name: "env",
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(deploymentEnvVar("ENV")),
			altsrc.ConfigFile(configFlag.Name, "deployment.env"),
		),
		Usage:    "`<KEY=VALUE>` formatted environment variables (use quotes for values with spaces or commas)",
		Category: "Deployment:",
	}

	requestedMemoryFlag = &workaround.FloatFlag{
		Name: "requested-memory-mb",
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(deploymentEnvVar("REQUESTED_MEMORY_MB")),
			altsrc.ConfigFile(configFlag.Name, "deployment.requested-memory-mb"),
		),
		Usage:      "`<memory-in-mb>` to allocate to your process",
		Persistent: true,
		Category:   "Deployment:",
	}

	requestedCPUFlag = &workaround.FloatFlag{
		Name: "requested-cpu",
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(deploymentEnvVar("REQUESTED_CPU")),
			altsrc.ConfigFile(configFlag.Name, "deployment.requested-cpu"),
		),
		Usage:      "`<cores>` to allocate to your process",
		Persistent: true,
		Category:   "Deployment:",
	}

	fromLatestFlag = &cli.BoolFlag{
		Name:     "from-latest",
		Sources:  cli.EnvVars(deploymentEnvVar("FROM_LATEST")),
		Usage:    "whether to use settings from the latest deployment; if true other flags and config file values act as overrides",
		Category: "Deployment:",
	}

	deploymentTagFlag = &cli.StringFlag{
		Name: "deployment-tag",
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(deploymentEnvVar("DEPLOYMENT_TAG")),
			altsrc.ConfigFile(configFlag.Name, "deployment.deployment-tag"),
		),
		Usage:    "arbitrary metadata associated with a deployment",
		Category: "Deployment:",
	}
)

func parseContainerPorts(ports []string) ([]components.ContainerPort, error) {
	// this converts a string representation of the slice from a config file into a real string slice
	if len(ports) == 1 && strings.HasPrefix(ports[0], "[") && strings.HasSuffix(ports[0], "]") {
		ports = strings.Split(strings.TrimPrefix(strings.TrimSuffix(ports[0], "]"), "["), " ")
	}
	output := make([]components.ContainerPort, 0, len(ports))
	for _, port := range ports {
		p, err := shorthand.ParseContainerPort(port)
		if err != nil {
			return nil, err
		}
		output = append(output, *p)
	}
	return output, nil
}

func parseEnvVars(envVars []string) ([]components.DeploymentConfigV3Env, error) {
	// Envs from a Config File are parsed from urfave/cli as a single-element
	// string slice of all the values like:
	// []string{`[KEY1=VAL1 KEY2=VAL2 KEY3="QUOTED VAL3"]`}
	// This stanza parses those into a proper slice of one Env per slice element
	if len(envVars) == 1 && strings.HasPrefix(envVars[0], "[") && strings.HasSuffix(envVars[0], "]") {
		var err error
		envVars, err = shorthand.ParseConfigFileVars(envVars[0])
		if err != nil {
			return nil, err
		}
	}
	envVars = fixOverZealousCommaSplitting(envVars)
	output := make([]components.DeploymentConfigV3Env, 0, len(envVars))
	for _, envVar := range envVars {
		env, err := shorthand.ParseDeploymentEnvVar(envVar)
		if err != nil {
			return nil, err
		}
		output = append(output, *env)
	}
	return output, nil
}

func fixOverZealousCommaSplitting(envVars []string) []string {
	var fixedEnvVars []string
	var currentPair string
	for _, val := range envVars {
		if strings.Contains(val, "=") && !strings.HasSuffix(currentPair, "=") {
			// Start a new key-value pair
			if currentPair != "" {
				fixedEnvVars = append(fixedEnvVars, currentPair)
			}
			currentPair = val
		} else {
			// Append to the current key-value pair
			currentPair += "," + val
		}
	}
	if currentPair != "" {
		fixedEnvVars = append(fixedEnvVars, currentPair)
	}

	return fixedEnvVars
}

var (
	deploymentConfigKey = "commands.DeploymentConfig.DI"
)

type DeploymentConfig struct {
	*GlobalConfig
	SDK    *sdk.HathoraCloud
	Output output.FormatWriter
}

var _ LoadableConfig = (*DeploymentConfig)(nil)

func (c *DeploymentConfig) Load(cmd *cli.Command) error {
	global, err := GlobalConfigFrom(cmd)
	if err != nil {
		return err
	}
	c.GlobalConfig = global

	c.SDK = setup.SDK(c.Token, c.BaseURL, c.Verbosity)
	var deployment components.DeploymentV3
	output, err := OutputFormatterFor(cmd, deployment)
	if err != nil {
		return err
	}
	c.Output = output
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
	DeploymentID string
}

var _ LoadableConfig = (*OneDeploymentConfig)(nil)

func (c *OneDeploymentConfig) Load(cmd *cli.Command) error {
	deployment, err := DeploymentConfigFrom(cmd)
	if err != nil {
		return err
	}
	c.DeploymentConfig = deployment
	c.DeploymentID = cmd.String(deploymentIDFlag.Name)
	c.Log = c.Log.With(zap.String("deployment.id", c.DeploymentID))
	return nil
}

func (c *OneDeploymentConfig) New() LoadableConfig {
	return &OneDeploymentConfig{}
}

func OneDeploymentConfigFrom(cmd *cli.Command) (*OneDeploymentConfig, error) {
	return ConfigFromCLI[*OneDeploymentConfig](oneDeploymentConfigKey, cmd)
}

var (
	createDeploymentConfigKey = "commands.CreateDeploymentConfig.DI"
)

type CreateDeploymentConfig struct {
	*DeploymentConfig
	BuildID                  string
	IdleTimeoutEnabled       *bool
	RoomsPerProcess          int
	TransportType            components.TransportType
	ContainerPort            int
	RequestedMemoryMB        float64
	RequestedCPU             float64
	AdditionalContainerPorts []components.ContainerPort
	Env                      []components.DeploymentConfigV3Env
	DeploymentTag            string
}

var _ LoadableConfig = (*CreateDeploymentConfig)(nil)

func (c *CreateDeploymentConfig) Load(cmd *cli.Command) error {
	deployment, err := DeploymentConfigFrom(cmd)
	if err != nil {
		return err
	}

	c.DeploymentConfig = deployment
	c.BuildID = cmd.String(buildIDFlag.Name)

	// Value of the idleTimeoutFlag by priority, high to low
	// Passed in as an argument
	// From latest deployment config (if from-latest is true)
	// Default true
	if cmd.IsSet(idleTimeoutFlag.Name) {
		idleTimeoutEnabled := cmd.Bool(idleTimeoutFlag.Name)
		c.IdleTimeoutEnabled = &idleTimeoutEnabled
	} else {
		idleTimeoutEnabled := true
		c.IdleTimeoutEnabled = &idleTimeoutEnabled
	}

	c.RoomsPerProcess = int(cmd.Int(roomsPerProcessFlag.Name))
	c.TransportType = components.TransportType(cmd.String(transportTypeFlag.Name))
	c.ContainerPort = int(cmd.Int(containerPortFlag.Name))
	c.RequestedMemoryMB = cmd.Float(requestedMemoryFlag.Name)
	c.RequestedCPU = cmd.Float(requestedCPUFlag.Name)
	c.DeploymentTag = cmd.String(deploymentTagFlag.Name)

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

func (c *CreateDeploymentConfig) Merge(latest *components.DeploymentV3) {
	if latest == nil {
		return
	}

	if c.BuildID == "" {
		c.BuildID = latest.BuildID
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

func (c *CreateDeploymentConfig) Validate() error {
	var err error

	if c.AppID == nil || *c.AppID == "" {
		err = errors.Join(err, missingRequiredFlag(appIDFlag.Name))
	}

	if c.BuildID == "" {
		err = errors.Join(err, missingRequiredFlag(buildIDFlag.Name))
	}

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

	if c.RequestedMemoryMB < minMemoryMB {
		err = errors.Join(err, fmt.Errorf("flag %s must be greater than %s",
			requestedMemoryFlag.Name,
			strconv.FormatFloat(float64(minMemoryMB), 'f', -1, 64),
		))
	}

	if c.RequestedCPU == 0 {
		err = errors.Join(err, missingRequiredFlag(requestedCPUFlag.Name))
	}

	if c.RequestedCPU < minCPU {
		err = errors.Join(err, fmt.Errorf("flag %s must be greater than %s",
			requestedCPUFlag.Name,
			strconv.FormatFloat(float64(minCPU), 'f', -1, 64),
		))
	}

	err = errors.Join(err, requireMaxDecimals(c.RequestedCPU, maxCPUDecimalPlaces, requestedCPUFlag.Name))

	return err
}

func (c *CreateDeploymentConfig) New() LoadableConfig {
	return &CreateDeploymentConfig{}
}

func CreateDeploymentConfigFrom(cmd *cli.Command) (*CreateDeploymentConfig, error) {
	return ConfigFromCLI[*CreateDeploymentConfig](createDeploymentConfigKey, cmd)
}
