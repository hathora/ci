package commands

import (
	"context"
	"errors"
	"fmt"
	"github.com/urfave/cli/v2/altsrc"
	"math"

	"github.com/hathora/ci/internal/output"
	"github.com/urfave/cli/v2"
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
	Load(cCtx *cli.Context) error
	SetContext(context.Context)
}

func ConfigFromCtx[T LoadableConfig](key any, ctx context.Context) (T, bool) {
	var nilT T
	untyped := ctx.Value(key)
	if untyped == nil {
		return nilT, false
	}
	typed, isType := untyped.(T)
	if isType {
		return typed, true
	}

	return nilT, false
}

func ConfigFromCLI[T LoadableConfig](key any, cCtx *cli.Context) (T, error) {
	cfg, hasCfg := ConfigFromCtx[T](key, cCtx.Context)
	if !hasCfg {
		cfg = cfg.New().(T)
		err := cfg.Load(cCtx)
		if err != nil {
			return cfg, err
		}
		cCtx.Context = context.WithValue(cCtx.Context, key, cfg)
		cfg.SetContext(cCtx.Context)
	}
	return cfg, nil
}

func SetFlagsFromFile (cCtx *cli.Context, flags []cli.Flag) error {
	if cCtx.IsSet(configFileFlag.Name) {
		if err := altsrc.InitInputSourceWithContext(flags, altsrc.NewYamlSourceFromFlagFunc(configFileFlag.Name))(cCtx); err != nil {
			return err
		}
	}
	return nil
}

type GlobalConfig struct {
	Token     string
	BaseURL   string
	AppID     *string
	Output    output.FormatWriter
	Verbosity int
	Context   context.Context
}

func (c *GlobalConfig) Load(cCtx *cli.Context) error {
	c.Token = cCtx.String(tokenFlag.Name)
	c.BaseURL = cCtx.String(hathoraCloudEndpointFlag.Name)
	appID := cCtx.String(appIDFlag.Name)
	if appID == "" {
		c.AppID = nil
	} else {
		c.AppID = &appID
	}

	outputType := cCtx.String(outputTypeFlag.Name)
	switch output.ParseOutputType(outputType) {
	case output.JSON:
		c.Output = output.JSONFormat(cCtx.Bool(outputPrettyFlag.Name))
	case output.Text:
		c.Output = output.TextFormat()
	default:
		return fmt.Errorf("unsupported output type: %s", outputType)
	}

	verboseCount := cCtx.Count(verboseFlag.Name)
	verbosity := cCtx.Int(verbosityFlag.Name)
	c.Verbosity = int(math.Max(float64(verbosity), float64(verboseCount)))
	return nil
}

func (c *GlobalConfig) New() LoadableConfig {
	return &GlobalConfig{}
}

func (c *GlobalConfig) SetContext(ctx context.Context) {
	c.Context = ctx
}

var (
	globalConfigKey        = struct{}{}
	ErrMissingRequiredFlag = errors.New("missing required flag")
)

func GlobalConfigFrom(cCtx *cli.Context) (*GlobalConfig, error) {
	if cCtx.IsSet(configFileFlag.Name) {
		if err := altsrc.InitInputSourceWithContext(GlobalFlags, altsrc.NewYamlSourceFromFlagFunc(configFileFlag.Name))(cCtx); err != nil {
			return nil, err
		}
	}

	cfg, err := ConfigFromCLI[*GlobalConfig](globalConfigKey, cCtx)
	if err != nil {
		return nil, err
	}

	if cfg.AppID == nil || *cfg.AppID == "" {
		return nil, fmt.Errorf("%w %s", ErrMissingRequiredFlag, appIDFlag.Name)
	}
	if cfg.Token == "" {
		return nil, fmt.Errorf("%w %s", ErrMissingRequiredFlag, tokenFlag.Name)
	}

	return cfg, nil
}

func isCallForHelp(ctx *cli.Context) bool {
	for _, arg := range ctx.Args().Slice() {
		if arg == "--help" || arg == "-h" {
			return true
		}
	}
	return false
}
