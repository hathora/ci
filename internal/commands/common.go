package commands

import (
	"context"
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
	Load(cCtx *cli.Context)
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

func ConfigFromCLI[T LoadableConfig](key any, cCtx *cli.Context) T {
	cfg, hasCfg := ConfigFromCtx[T](key, cCtx.Context)
	if !hasCfg {
		cfg.Load(cCtx)
		cCtx.Context = context.WithValue(cCtx.Context, key, cfg)
		cfg.SetContext(cCtx.Context)
	}
	return cfg
}

type GlobalConfig struct {
	Token     string
	BaseURL   string
	AppID     *string
	Output    output.FormatWriter
	Verbosity int
	Context   context.Context
}

func (c *GlobalConfig) Load(cCtx *cli.Context) {
	c.Token = cCtx.String(tokenFlag.Name)
	c.BaseURL = cCtx.String(hathoraCloudEndpointFlag.Name)
	appID := cCtx.String(appIDFlag.Name)
	if appID == "" {
		c.AppID = nil
	} else {
		c.AppID = &appID
	}
	switch output.ParseOutputType(cCtx.String(outputTypeFlag.Name)) {
	case output.JSON:
		c.Output = output.JSONFormat(cCtx.Bool(outputPrettyFlag.Name))
	default:
		c.Output = output.TextFormat()
	}
	verboseCount := cCtx.Count(verboseFlag.Name)
	verbosity := cCtx.Int(verbosityFlag.Name)
	c.Verbosity = int(math.Max(float64(verbosity), float64(verboseCount)))
}

func (c *GlobalConfig) SetContext(ctx context.Context) {
	c.Context = ctx
}

var (
	globalConfigKey = struct{}{}
)

func GlobalConfigFrom(cCtx *cli.Context) *GlobalConfig {
	return ConfigFromCLI[*GlobalConfig](globalConfigKey, cCtx)
}
