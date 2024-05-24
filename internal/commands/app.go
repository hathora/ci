package commands

import (
	"context"

	"github.com/hathora/ci/internal/commands/altsrc"
	"github.com/hathora/ci/internal/setup"
	"github.com/urfave/cli/v3"
)

func App() *cli.Command {
	var cleanup []func()
	return &cli.Command{
		Name:                   "cloud-ci",
		Aliases:                []string{"hathora-ci", "ci"},
		EnableShellCompletion:  true,
		Suggest:                true,
		UseShortOptionHandling: true,
		SliceFlagSeparator:     ",",
		Usage:                  "a CLI tool for for CI/CD workflows to manage deployments and builds in hathora.dev",
		Flags:                  GlobalFlags,
		Before: func(ctx context.Context, cmd *cli.Command) error {
			if isCallForHelp(cmd) {
				return nil
			}
			err := altsrc.InitializeValueSourcesFromFlags(ctx, cmd)
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
