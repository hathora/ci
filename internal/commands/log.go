package commands

import (
	"context"
	"errors"
	"fmt"
	"github.com/hathora/ci/internal/commands/altsrc"
	"github.com/hathora/ci/internal/output"
	"github.com/hathora/ci/internal/sdk"
	"github.com/hathora/ci/internal/setup"
	"github.com/hathora/ci/internal/workaround"
	"github.com/urfave/cli/v3"
	"go.uber.org/zap"
	"os"
)

var (
	minTailLines = 0
	maxTailLines = 5000
)

var Log = &cli.Command{
	Name:  "log",
	Usage: "view live process logs",
	Flags: subcommandFlags(followFlag, processIDFlag, tailLinesFlag),
	Action: func(ctx context.Context, cmd *cli.Command) error {
		log, err := OneLogConfigFrom(cmd)
		if err != nil {
			//nolint:errcheck
			cli.ShowSubcommandHelp(cmd)
			return err
		}

		if err := log.Validate(); err != nil {
			//nolint:errcheck
			cli.ShowSubcommandHelp(cmd)
			return err
		}

		log.Log.Debug("Getting build server logs...")

		// validate

		res, err := log.SDK.LogV1.GetLogsForProcess(
			ctx,
			log.ProcessID,
			log.AppID,
			sdk.Bool(log.Follow),
			sdk.Int(log.TailLines))

		if err != nil {
			return fmt.Errorf("failed to get logs: %w", err)
		}

		err = output.StreamOutput(res.Stream, os.Stderr)
		if err != nil {
			zap.L().Error("failed to stream output to console", zap.Error(err))
		}

		return nil
	},
}

func logFlagEnvVar(name string) string {
	return logFlagEnvVarPrefix + name
}

var (
	logFlagEnvVarPrefix = globalFlagEnvVarPrefix + "LOG_"

	followFlag = &cli.BoolFlag{
		Name: "follow",
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(logFlagEnvVar("FOLLOW")),
			altsrc.ConfigFile(configFlag.Name, "log.follow")),
		Usage:      "to streams logs in real time",
		Value:      false,
		Category:   "Log:",
		Persistent: true,
	}

	processIDFlag = &cli.StringFlag{
		Name:    "process-id",
		Aliases: []string{"p"},
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(logFlagEnvVar("PROCESS_ID")),
			altsrc.ConfigFile(configFlag.Name, "log.process-id"),
		),
		Category: "Log:",
		Usage:    "System generated unique identifier to a runtime instance of your game server",
	}

	tailLinesFlag = &workaround.IntFlag{
		Name: "tail-lines",
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(buildFlagEnvVar("TAIL_LINES")),
			altsrc.ConfigFile(configFlag.Name, "log.tail-lines"),
		),
		Usage:      "Number of lines to return from most recent logs history.",
		Value:      100,
		Category:   "Log:",
		Persistent: true,
	}
)

var (
	logConfigKey = "commands.LogConfig.DI"
)

type LogConfig struct {
	*GlobalConfig
	SDK *sdk.SDK
}

var _ LoadableConfig = (*LogConfig)(nil)

func (c *LogConfig) Load(cmd *cli.Command) error {
	global, err := GlobalConfigFrom(cmd)
	if err != nil {
		return err
	}
	c.GlobalConfig = global
	c.SDK = setup.SDK(c.Token, c.BaseURL, c.Verbosity)
	return nil
}

func (c *LogConfig) New() LoadableConfig { return &LogConfig{} }

func LogConfigFrom(cmd *cli.Command) (*LogConfig, error) {
	return ConfigFromCLI[*LogConfig](logConfigKey, cmd)
}

var (
	oneLogConfigKey = "commands.OneLogConfig.DI"
)

type OneLogConfig struct {
	*LogConfig
	Follow    bool
	TailLines int
	ProcessID string
}

var _ LoadableConfig = (*OneLogConfig)(nil)

func (c *OneLogConfig) Load(cmd *cli.Command) error {
	log, err := LogConfigFrom(cmd)
	if err != nil {
		return err
	}
	c.ProcessID = cmd.String(processIDFlag.Name)
	c.Follow = cmd.Bool(followFlag.Name)
	c.TailLines = int(cmd.Int(tailLinesFlag.Name))
	c.LogConfig = log
	return nil
}

func (c *OneLogConfig) New() LoadableConfig { return &OneLogConfig{} }

func OneLogConfigFrom(cmd *cli.Command) (*OneLogConfig, error) {
	return ConfigFromCLI[*OneLogConfig](oneDeploymentConfigKey, cmd)
}

func (c *OneLogConfig) Validate() error {
	var err error
	err = errors.Join(err, requireIntInRange(c.TailLines, minTailLines, maxTailLines, tailLinesFlag.Name))

	return err
}
