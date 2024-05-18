package main

import (
	"log"
	"os"

	"github.com/hathora/ci/internal/commands"
	"github.com/hathora/ci/internal/setup"
	"github.com/urfave/cli/v2"
)

func main() {
	var cleanup []func()
	app := &cli.App{
		EnableBashCompletion:   true,
		Suggest:                true,
		UseShortOptionHandling: true,
		SliceFlagSeparator:     ",",
		Before: func(c *cli.Context) error {
			cfg := commands.GlobalConfigFrom(c)
			_, cleanupLogger := setup.Logger(cfg.Verbosity)
			cleanup = append(cleanup, cleanupLogger)
			return nil
		},
		Commands: []*cli.Command{
			commands.Build,
			commands.Deployment,
		},
		After: func(c *cli.Context) error {
			for _, fn := range cleanup {
				fn()
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
