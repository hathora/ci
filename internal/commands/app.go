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

func init() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:     "version",
		Usage:    "print the version",
		Category: "Global:",
	}

	cli.HelpFlag = &cli.BoolFlag{
		Name:     "help",
		Aliases:  []string{"h"},
		Usage:    "show help",
		Category: "Global:",
	}
}

func App() *cli.Command {
	var cleanup []func()
	return &cli.Command{
		Name:                          "hathora",
		EnableShellCompletion:         true,
		Suggest:                       true,
		UseShortOptionHandling:        true,
		SliceFlagSeparator:            ",",
		Usage:                         "a CLI tool for for CI/CD workflows to manage deployments and builds in hathora.dev",
		Flags:                         GlobalFlags,
		Version:                       BuildVersion,
		CustomRootCommandHelpTemplate: cli.SubcommandHelpTemplate,
		Before: func(ctx context.Context, cmd *cli.Command) error {
			cfg, err := VerbosityConfigFrom(cmd)
			if err != nil {
				return err
			}
			_, cleanupLogger := setup.Logger(cfg.Verbosity)
			cleanup = append(cleanup, cleanupLogger)
			handleNewVersionAvailable(BuildVersion)

			if isCallForHelp(cmd) {
				return nil
			}

			err = altsrc.InitializeValueSourcesFromFlags(ctx, cmd, os.Args[1:])
			if err != nil {
				return err
			}
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
