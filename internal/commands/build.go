package commands

import (
	"context"
	"github.com/hathora/ci/internal/sdk"
	"github.com/hathora/ci/internal/sdk/pkg/models/operations"
	"github.com/hathora/ci/internal/sdk/pkg/models/shared"
	"github.com/hathora/ci/internal/setup"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func logBuild(build shared.Build) {
	zap.L().Info("",	zap.Any("Build", build))
}

func logBuildP(build *shared.Build) {
	logBuild(*build)
}

var Build = &cli.Command{
	Name:  "build",
	Usage: "options for builds",
	Subcommands: []*cli.Command{
		{
			Name:  "get-build-info",
			Usage: "get a build",
			Flags: append([]cli.Flag{
				buildIDFlag,
			}, globalFlags...),
			Action: func(cCtx *cli.Context) error {
				token, baseUrl, appID, _ := getCommonFlagValues(cCtx)
				buildID := cCtx.Int(buildIDFlag.Name)
				sdk := setup.SDK(token, baseUrl)

				res, err := sdk.BuildV2.GetBuildInfo(context.Background(), buildID, appID)
				if err != nil {
					return err
				}

				logBuildP(res.Build)

				return nil
			},
		},
		{
			Name:  "get-builds",
			Usage: "get all builds",
			Flags: append([]cli.Flag{}, globalFlags...),
			Action: func(cCtx *cli.Context) error {
				token, baseUrl, appID, _ := getCommonFlagValues(cCtx)
				sdk := setup.SDK(token, baseUrl)

				res, err := sdk.BuildV2.GetBuilds(context.Background(), appID)
				if err != nil {
					return err
				}

				if len(res.Builds) == 0 {
					zap.L().Info("You have no builds")
					return nil
				}

				for _, build := range res.Builds {
					logBuild(build)
				}

				return nil
			},
		},
		{
			Name:  "create-build",
			Usage: "create a build",
			Flags: append([]cli.Flag{
				buildTagFlag,
			}, globalFlags...),
			Action: func(cCtx *cli.Context) error {
				token, baseUrl, appID, _ := getCommonFlagValues(cCtx)
				buildTag := sdk.String(cCtx.String(buildTagFlag.Name))
				sdk := setup.SDK(token, baseUrl)

				res, err := sdk.BuildV2.CreateBuild(
					context.Background(),
					shared.CreateBuildParams{BuildTag: buildTag},
					appID)
				if err != nil {
					return err
				}

				logBuildP(res.Build)
				return nil
			},
		},
		{
			Name:  "run-build",
			Usage: "run a build by id",
			Flags: append([]cli.Flag{
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
				token, baseUrl, appID, _ := getCommonFlagValues(cCtx)
				buildID := cCtx.Int(buildIDFlag.Name)
				filePath := cCtx.String("binary-path")
				sdk := setup.SDK(token, baseUrl)

				file, err := EnforceTar(filePath)
				if err != nil {
					return err
				}

				res, err := sdk.BuildV2.RunBuild(
					context.Background(),
					buildID,
					operations.RunBuildRequestBody{
						File: operations.RunBuildFile{
							FileName: "build.tgz",
							Content: file,
						},
					},
					appID)
				if err != nil {
					return err
				}

				// TODO stream response to zap logger
				zap.L().Info("Status Code", zap.Int("code", res.StatusCode))
				return nil
			},
		},
		{
			Name:  "delete-build",
			Usage: "delete a build",
			Flags: append([]cli.Flag{
				buildIDFlag,
			}, globalFlags...),
			Action: func(cCtx *cli.Context) error {
				token, baseUrl, appID, _ := getCommonFlagValues(cCtx)
				buildId := cCtx.Int(buildIDFlag.Name)
				sdk := setup.SDK(token, baseUrl)

				_, err := sdk.BuildV2.DeleteBuild(context.Background(), buildId, appID)
				if err != nil {
					return err
				}
				zap.L().Info("Build successfully deleted")

				return nil
			},
		},
	},
}
