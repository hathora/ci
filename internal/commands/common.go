package commands

import (
	"fmt"
	"math"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/hathora/ci/internal/output"
	"github.com/hathora/ci/internal/sdk/models/shared"
	"github.com/urfave/cli/v3"
	"go.uber.org/zap"
)

var (
	infoCommandName   = "info"
	listCommandName   = "list"
	createCommandName = "create"
	deleteCommandName = "delete"
	latestCommandName = "latest"
)

type LoadableConfig interface {
	New() LoadableConfig
	Load(cmd *cli.Command) error
}

func configFromMetadata[T LoadableConfig](key string, md map[string]any) (T, bool) {
	var nilT T
	untyped, hasKey := md[key]
	if !hasKey {
		return nilT, false
	}
	typed, isType := untyped.(T)
	if isType {
		return typed, true
	}

	return nilT, false
}

func ConfigFromCLI[T LoadableConfig](key string, cmd *cli.Command) (T, error) {
	cfg, hasCfg := configFromMetadata[T](key, cmd.Metadata)
	if !hasCfg {
		cfg = cfg.New().(T)
	}
	err := cfg.Load(cmd)
	if err != nil {
		return cfg, err
	}

	cmd.Metadata[key] = cfg
	return cfg, nil
}

type GlobalConfig struct {
	Token     string
	BaseURL   string
	AppID     *string
	Output    output.FormatWriter
	Verbosity int
	Log       *zap.Logger
}

func (c *GlobalConfig) Load(cmd *cli.Command) error {
	c.Token = cmd.String(tokenFlag.Name)
	c.BaseURL = cmd.String(hathoraCloudEndpointFlag.Name)
	appID := cmd.String(appIDFlag.Name)
	if appID == "" {
		c.AppID = nil
	} else {
		c.AppID = &appID
	}

	outputType := cmd.String(outputTypeFlag.Name)
	switch output.ParseOutputType(outputType) {
	case output.JSON:
		c.Output = output.JSONFormat(cmd.Bool(outputPrettyFlag.Name))
	case output.Text:
		c.Output = BuildTextFormatter()
	default:
		return fmt.Errorf("unsupported output type: %s", outputType)
	}

	verboseCount := cmd.Count(verboseFlag.Name)
	verbosity := cmd.Int(verbosityFlag.Name)
	c.Verbosity = int(math.Max(float64(verbosity), float64(verboseCount)))
	c.Log = zap.L().With(zap.String("app.id", appID))
	return nil
}

func (c *GlobalConfig) New() LoadableConfig {
	return &GlobalConfig{}
}

var (
	globalConfigKey = "commands.GlobalConfig.DI"
)

func GlobalConfigFrom(cmd *cli.Command) (*GlobalConfig, error) {
	cfg, err := ConfigFromCLI[*GlobalConfig](globalConfigKey, cmd)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func isCallForHelp(cmd *cli.Command) bool {
	for _, arg := range cmd.Args().Slice() {
		if arg == "--help" || arg == "-h" {
			return true
		}
	}
	return false
}

func BuildTextFormatter() output.FormatWriter {
	// TODO: Allow commands to register their own formatters so that this one function doesn't have to know the desired format for every type
	var build shared.Build
	var deployment shared.DeploymentV2
	var envVar shared.DeploymentV2Env
	var containerPort shared.ContainerPort
	var timestamp time.Time
	var float float64
	var long int64
	return output.TextFormat(
		output.WithFieldOrder(build,
			"BuildID",
			"BuildTag",
			"CreatedAt",
			"Status",
			"ImageSize",
			"StartedAt",
			"FinishedAt",
		),
		output.WithoutFields(build, "AppID", "RegionalContainerTags", "DeletedAt", "CreatedBy"),
		output.WithFieldOrder(deployment,
			"DeploymentID",
			"BuildID",
			"CreatedAt",
			"IdleTimeoutEnabled",
			"RoomsPerProcess",
			"DefaultContainerPort",
			"AdditionalContainerPorts",
			"Env",
		),
		output.WithoutFields(deployment, "AppID", "CreatedBy"),
		output.WithFormatter(envVar,
			func(e shared.DeploymentV2Env) string {
				return fmt.Sprintf("%s=%s", e.Name, e.Value)
			},
		),
		output.WithFormatter(containerPort,
			func(cp shared.ContainerPort) string {
				return fmt.Sprintf("%s:%d/%s", cp.Name, cp.Port, cp.TransportType)
			},
		),
		output.WithFormatter(timestamp,
			func(t time.Time) string {
				// TODO: consider using this human-friendly time format
				// return humanize.Time(t)
				return t.Format(time.RFC3339)
			},
		),
		output.WithFormatter(float,
			func(f float64) string {
				return humanize.Ftoa(f)
			},
		),
		// TODO: this should not be how we generally handle all int64s, so we may want support for targeted formatters by field name
		output.WithFormatter(long,
			func(l int64) string {
				return humanize.IBytes(uint64(l))
			},
		),
	)
}
