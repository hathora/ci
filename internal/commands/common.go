package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"go/version"
	"math"
	"net/http"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/urfave/cli/v3"
	"go.uber.org/zap"

	"github.com/hathora/ci/internal/output"
	"github.com/hathora/ci/internal/sdk/models/shared"
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

type VerbosityConfig struct {
	Verbosity int
	Log       *zap.Logger
}

func (c *VerbosityConfig) Load(cmd *cli.Command) error {
	// we subtract 1 because the flag is counted an additional time for the
	// --verbose alias
	verboseCount := cmd.Count(verboseFlag.Name) - 1
	verbosity := cmd.Int(verbosityFlag.Name)
	c.Verbosity = int(math.Max(float64(verbosity), float64(verboseCount)))
	return nil
}

func (c *VerbosityConfig) New() LoadableConfig {
	return &VerbosityConfig{}
}

var (
	verbosityConfigKey = "commands.VerbosityConfig.DI"
)

func VerbosityConfigFrom(cmd *cli.Command) (*VerbosityConfig, error) {
	cfg, err := ConfigFromCLI[*VerbosityConfig](verbosityConfigKey, cmd)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

type GlobalConfig struct {
	*VerbosityConfig
	Token   string
	BaseURL string
}

func (c *GlobalConfig) Load(cmd *cli.Command) error {
	verbosityConfig, err := VerbosityConfigFrom(cmd)
	if err != nil {
		return err
	}
	c.VerbosityConfig = verbosityConfig
	c.Token = cmd.String(tokenFlag.Name)
	if c.Token == "" {
		err = errors.Join(err, missingRequiredFlag(tokenFlag.Name))
	}
	c.BaseURL = cmd.String(hathoraCloudEndpointFlag.Name)
	if c.BaseURL == "" {
		err = errors.Join(err, missingRequiredFlag(hathoraCloudEndpointFlag.Name))
	}
	c.Log = zap.L()
	return err
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
		if arg == "--help" || arg == "-h" || arg == "help" {
			return true
		}
	}
	return false
}

func OutputFormatterFor(cmd *cli.Command, outputType any) (output.FormatWriter, error) {
	outputFmt := cmd.String(outputFlag.Name)
	switch output.ParseOutputType(outputFmt) {
	case output.JSON:
		return output.JSONFormat(cmd.Bool(outputPrettyFlag.Name)), nil
	case output.Text:
		return BuildTextFormatter(), nil
	case output.Value:
		fieldName := strings.TrimSuffix(outputFmt, "Value")
		if len(fieldName) == 0 {
			return nil, fmt.Errorf("invalid value format: %s", outputType)
		}
		return output.ValueFormat(outputType, fieldName)
	default:
		return nil, fmt.Errorf("unsupported output type: %s", outputType)
	}
}

func BuildTextFormatter() output.FormatWriter {
	// TODO: Allow commands to register their own formatters so that this one function doesn't have to know the desired format for every type
	var build shared.BuildV3
	var deployment shared.DeploymentV3
	var envVar shared.DeploymentV3Env
	var containerPort shared.ContainerPort
	var timestamp time.Time
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
		output.WithoutFields(build, "RegionalContainerTags", "DeletedAt", "CreatedBy"),
		output.WithFieldOrder(deployment,
			"AppID",
			"DeploymentID",
			"BuildID",
			"CreatedAt",
			"IdleTimeoutEnabled",
			"RoomsPerProcess",
			"RequestedCPU",
			"RequestedMemoryMB",
			"DefaultContainerPort",
			"AdditionalContainerPorts",
			"BuildTag",
		),
		output.RenameField(deployment, "RequestedMemoryMB", "RequestedMemory"),
		output.WithPropertyFormatter(deployment, "RequestedMemoryMB", func(f float64) string {
			return humanize.IBytes((uint64)(f * 1024 * 1024))
		}),
		output.WithoutFields(deployment, "CreatedBy", "Env"),
		output.WithFormatter(envVar,
			func(e shared.DeploymentV3Env) string {
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
		// TODO: this should not be how we generally handle all int64s, so we may want support for targeted formatters by field name
		output.WithFormatter(long,
			func(l int64) string {
				return humanize.IBytes(uint64(l))
			},
		),
	)
}

type Release struct {
	TagName     string `json:"tag_name"`
	Name        string `json:"name"`
	PublishedAt string `json:"published_at"`
	URL         string `json:"html_url"`
}

func handleNewVersionAvailable(currentVersion string) {
	url := "https://api.github.com/repos/hathora/ci/releases/latest"

	resp, err := http.Get(url)
	if err != nil {
		zap.L().Warn("unable to fetch latest version number")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		zap.L().Warn("unable to fetch latest version number")
	}

	var release Release
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		zap.L().Warn("unable to decode the latest version number")
	}

	prefixedVersion := "go" + currentVersion
	showLatest := false
	if !version.IsValid(prefixedVersion) {
		zap.L().Warn("You are using a development version of the hathora cli.")
		showLatest = true
	} else if version.Compare("go"+release.TagName, prefixedVersion) > 0 {
		zap.L().Warn("You are using an outdated version of the hathora cli.")
		showLatest = true
	}

	if showLatest {
		zap.L().Warn("Version " + release.TagName + " is available for download.")
	}
}
