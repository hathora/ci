package main

import (
	"log"
	"os"

	"github.com/hathora/ci/internal/commands"
	"github.com/hathora/ci/internal/setup"
	"github.com/urfave/cli/v2"
)

func main() {
	_, cleanupLogger := setup.Logger()
	defer cleanupLogger()

	app := &cli.App{
		EnableBashCompletion: true,
		Suggest:              true,
		Commands: []*cli.Command{
			commands.Build,
			commands.Deployment,
		},
	}

	if err := app.Run(os.Args); err != nil {
		cleanupLogger()
		log.Fatal(err)
	}
}
