package commands

import (
	"context"
	"github.com/hathora/ci/internal/cloudapi/sdk"
	"github.com/hathora/ci/internal/setup"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

var Build = &cli.Command{
	Name:  "build",
	Usage: "options for builds",
	Subcommands: []*cli.Command{
		{
			Name:  "get-build-info",
			Usage: "get a build",
			Flags: append([]cli.Flag{
				appIDFlag,
				buildIDFlag,
			}, globalFlags...),
			Action: func(cCtx *cli.Context) error {
				zap.L().Info("Getting a build...")

				// init the sdk
				token := cCtx.String(tokenFlag.Name)
				url := cCtx.String(hathoraCloudEndpointFlag.Name)
				sdk := setup.SDK(token, url)

				// call the API
				// TODO figure out the correct context
				build, err := sdk.BuildV2.GetBuildInfo(context.Background(), cCtx.Int(buildIDFlag.Name), To(cCtx.String(appIDFlag.Name)))
				if err != nil {
					return err
				}

				// handle the response
				zap.L().Debug("logging build")

				return nil
			},
		},
		{
			Name:  "get-builds",
			Usage: "get all builds",
			Flags: append([]cli.Flag{
				appIDFlag,
			}, globalFlags...),
			Action: func(cCtx *cli.Context) error {
				zap.L().Info("Getting all builds...")
				return nil
			},
		},
		{
			Name:  "create-build",
			Usage: "create a build",
			Flags: append([]cli.Flag{
				appIDFlag,
				buildTagFlag,
			}, globalFlags...),
			Action: func(cCtx *cli.Context) error {
				zap.L().Info("Creating a build...")
				return nil
			},
		},
		{
			Name:  "run-build",
			Usage: "run a build by id",
			Flags: append([]cli.Flag{
				appIDFlag,
				buildIDFlag,
				&cli.StringFlag{
					Name:     "binary-path",
					Aliases:  []string{"bp"},
					EnvVars:  []string{"HATHORA_BUILD_BINARY_PATH"},
					Usage:    "path to the built game server binary",
					Required: true,
				},
			}, globalFlags...),
			Action: func(cCtx *cli.Context) error {
				zap.L().Info("Running a build...")
				return nil
			},
		},
		{
			Name:  "delete-build",
			Usage: "delete a build",
			Flags: append([]cli.Flag{
				appIDFlag,
				buildTagFlag,
			}, globalFlags...),
			Action: func(cCtx *cli.Context) error {
				zap.L().Info("Deleting a build...")
				return nil
			},
		},
	},
}
