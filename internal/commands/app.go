package commands

import (
	"context"
	"os"

	"github.com/urfave/cli/v3"

	"github.com/hathora/ci/internal/commands/altsrc"
	"github.com/hathora/ci/internal/setup"
)

var (
	BuildVersion = "unknown"
)

func App() *cli.Command {
	cli.VersionFlag = &cli.BoolFlag{
		Name:  "version",
		Usage: "print the version",
	}

	var cleanup []func()
	return &cli.Command{
		Name:                   "hathora",
		EnableShellCompletion:  true,
		Suggest:                true,
		UseShortOptionHandling: true,
		SliceFlagSeparator:     ",",
		Usage:                  "a CLI tool for for CI/CD workflows to manage deployments and builds in hathora.dev",
		Flags:                  GlobalFlags,
		Version:                BuildVersion,
		Before: func(ctx context.Context, cmd *cli.Command) error {
			handleNewVersionAvailable(BuildVersion)

			if isCallForHelp(cmd) {
				return nil
			}
			err := altsrc.InitializeValueSourcesFromFlags(ctx, cmd, os.Args[1:])
			if err != nil {
				return err
			}
			cfg, err := GlobalConfigFrom(cmd)
			if err != nil {
				return err
			}
			_, cleanupLogger := setup.Logger(cfg.Verbosity)
			cleanup = append(cleanup, cleanupLogger)
			return nil
		},
		Commands: []*cli.Command{
			Build,
			Deploy,
			Deployment,
		},
		After: func(ctx context.Context, c *cli.Command) error {
			for _, fn := range cleanup {
				fn()
			}
			return nil
		},
	}
}
