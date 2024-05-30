package commands

import (
	"encoding/json"
	"fmt"
	"go/version"
	"math"
	"net/http"

	"github.com/hathora/ci/internal/output"
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
		return fmt.Errorf("app-id is required")
	} else {
		c.AppID = &appID
	}

	outputType := cmd.String(outputTypeFlag.Name)
	switch output.ParseOutputType(outputType) {
	case output.JSON:
		c.Output = output.JSONFormat(cmd.Bool(outputPrettyFlag.Name))
	case output.Text:
		c.Output = output.TextFormat()
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

	versionDiff := version.Compare(release.TagName, currentVersion)

	if versionDiff > 0 {
		zap.L().Warn("A new version of hathora-ci is available for download.")
	}
}
