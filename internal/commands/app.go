package commands

import (
	"github.com/hathora/ci/internal/setup"
	"github.com/urfave/cli/v2"
)

func App() *cli.App {
	var cleanup []func()
	return &cli.App{
		Name:                   "cloud-ci",
		EnableBashCompletion:   true,
		Suggest:                true,
		UseShortOptionHandling: true,
		SliceFlagSeparator:     ",",
		Flags:                  GlobalFlags,
		Before: func(c *cli.Context) error {
			if isCallForHelp(c) {
				return nil
			}
			cfg, err := GlobalConfigFrom(c)
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
		After: func(c *cli.Context) error {
			for _, fn := range cleanup {
				fn()
			}
			return nil
		},
	}
}
