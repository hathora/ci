package commands

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
	"go.uber.org/zap"

	"github.com/hathora/ci/internal/commands/altsrc"
	"github.com/hathora/ci/internal/output"
	"github.com/hathora/ci/internal/sdk"
	"github.com/hathora/ci/internal/setup"
	"github.com/hathora/ci/internal/workaround"
)

var (
	minTailLines = 1
	maxTailLines = 5000
)

var Log = &cli.Command{
	Name:  "log",
	Usage: "view live process logs",
	Flags: subcommandFlags(followFlag, processIDFlag, tailLinesFlag),
	Action: func(ctx context.Context, cmd *cli.Command) error {
		log, err := ProcessLogConfigFrom(cmd)
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

		res, err := log.SDK.LogsV1.GetLogsForProcess(
			ctx,
			log.ProcessID,
			log.AppID,
			sdk.Bool(log.Follow),
			sdk.Int(log.TailLines))

		if err != nil {
			return fmt.Errorf("failed to get logs: %w", err)
		}

		err = output.StreamOutput(res.Stream, os.Stdout)
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
		Usage:      "whether to stream logs in real time",
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
		Usage:    "`<id>` of the runtime instance of your game server",
	}

	tailLinesFlag = &workaround.IntFlag{
		Name: "tail-lines",
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(buildFlagEnvVar("TAIL_LINES")),
			altsrc.ConfigFile(configFlag.Name, "log.tail-lines"),
		),
		Usage:      "`<number>` of lines to return from the most recent log history (1-5000)",
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
	processLogConfigKey = "commands.ProcessLogConfig.DI"
)

type ProcessLogConfig struct {
	*LogConfig
	Follow    bool
	TailLines int
	ProcessID string
}

var _ LoadableConfig = (*ProcessLogConfig)(nil)

func (c *ProcessLogConfig) Load(cmd *cli.Command) error {
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

func (c *ProcessLogConfig) New() LoadableConfig { return &ProcessLogConfig{} }

func ProcessLogConfigFrom(cmd *cli.Command) (*ProcessLogConfig, error) {
	return ConfigFromCLI[*ProcessLogConfig](processLogConfigKey, cmd)
}

func (c *ProcessLogConfig) Validate() error {
	var err error
	err = errors.Join(err, requireIntInRange(c.TailLines, minTailLines, maxTailLines, tailLinesFlag.Name))

	if c.ProcessID == "" {
		err = errors.Join(err, missingRequiredFlag(processIDFlag.Name))
	}

	return err
}
