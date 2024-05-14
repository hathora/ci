package commands

import (
	"encoding/json"
	"fmt"
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

				res, err := sdk.BuildV2.GetBuildInfo(cCtx.Context, buildID, appID)
				if err != nil {
					return fmt.Errorf("failed to get build info from API: %w", err)
				}

				// TODO this on other commands
				jsonBytes, err := json.Marshal(res.Build)
				if err != nil {
					return fmt.Errorf("failed to marshal build: %w", err)
				}

				return cli.Exit(string(jsonBytes), 0)
			},
		},
		{
			Name:  "get-builds",
			Usage: "get all builds",
			Flags: append([]cli.Flag{}, globalFlags...),
			Action: func(cCtx *cli.Context) error {
				token, baseUrl, appID, _ := getCommonFlagValues(cCtx)
				sdk := setup.SDK(token, baseUrl)

				res, err := sdk.BuildV2.GetBuilds(cCtx.Context, appID)
				if err != nil {
					return fmt.Errorf("failed to get builds from API: %w", err)
				}

				if len(res.Builds) == 0 {
					return cli.Exit("No builds were found", 1)
				}

				jsonBytes, err := json.Marshal(res.Builds)
				if err != nil {
					return fmt.Errorf("failed to marshal builds: %w", err)
				}

				return cli.Exit(string(jsonBytes), 0)
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
					cCtx.Context,
					shared.CreateBuildParams{BuildTag: buildTag},
					appID)
				if err != nil {
					return fmt.Errorf("failed to create a build in the API: %w", err)
				}

				jsonBytes, err := json.Marshal(res.Build)
				if err != nil {
					return fmt.Errorf("failed to marshal build: %w", err)
				}

				return cli.Exit(string(jsonBytes), 0)
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

				file, err := RequireTGZ(filePath)
				if err != nil {
					return fmt.Errorf("failed to get tgz file: %w", err)
				}

				zap.L().Debug("using archive file", zap.Any("file", file))

				_, err = sdk.BuildV2.RunBuild(
					cCtx.Context,
					buildID,
					operations.RunBuildRequestBody{
						File: operations.RunBuildFile{
							FileName: file.Name,
							Content:  file.Content,
						},
					},
					appID)
				if err != nil {
					return fmt.Errorf("failed to run a build in the API: %w", err)
				}

				return cli.Exit("Successfully ran build", 0)
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

				_, err := sdk.BuildV2.DeleteBuild(cCtx.Context, buildId, appID)
				if err != nil {
					return fmt.Errorf("failed to delete a build from API: %w", err)
				}

				return cli.Exit("Successfully deleted build", 0)
			},
		},
	},
}
