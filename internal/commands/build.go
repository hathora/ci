package commands

import (
	"fmt"

	"github.com/hathora/ci/internal/archive"
	"github.com/hathora/ci/internal/output"
	"github.com/hathora/ci/internal/sdk"
	"github.com/hathora/ci/internal/sdk/models/operations"
	"github.com/hathora/ci/internal/sdk/models/shared"
	"github.com/hathora/ci/internal/setup"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

var Build = &cli.Command{
	Name:  "build",
	Usage: "options for builds",
	Subcommands: []*cli.Command{
		{
			Name:    infoCommandName,
			Aliases: []string{"get-build-info"},
			Usage:   "get a build",
			Flags:   subcommandFlags(buildIDFlag),
			Action: func(cCtx *cli.Context) error {
				zap.L().Debug("getting build info...")
				build := OneBuildConfigFrom(cCtx)

				res, err := build.SDK.BuildV2.GetBuildInfo(build.Context, build.BuildID, build.AppID)
				if err != nil {
					return fmt.Errorf("failed to get build info: %w", err)
				}

				return output.As(build.OutputType, res.Build)
			},
		},
		{
			Name:    listCommandName,
			Aliases: []string{"get-builds"},
			Usage:   "get all builds",
			Flags:   subcommandFlags(),
			Action: func(cCtx *cli.Context) error {
				zap.L().Debug("getting builds...")
				build := BuildConfigFrom(cCtx)

				res, err := build.SDK.BuildV2.GetBuilds(cCtx.Context, build.AppID)
				if err != nil {
					return fmt.Errorf("failed to get builds: %w", err)
				}

				if len(res.Builds) == 0 {
					return fmt.Errorf("no builds found")
				}

				return output.As(build.OutputType, res.Builds)
			},
		},
		{
			Name:    createCommandName,
			Aliases: []string{"create-build"},
			Usage:   "create a build",
			Flags:   subcommandFlags(buildTagFlag),
			Action: func(cCtx *cli.Context) error {
				zap.L().Debug("creating a build...")
				build := BuildConfigFrom(cCtx)
				buildTag := sdk.String(cCtx.String(buildTagFlag.Name))

				res, err := build.SDK.BuildV2.CreateBuild(
					cCtx.Context,
					shared.CreateBuildParams{
						BuildTag: buildTag,
					},
					build.AppID,
				)
				if err != nil {
					return fmt.Errorf("failed to create a build: %w", err)
				}

				return output.As(build.OutputType, res.Build)
			},
		},
		{
			Name:    "run",
			Aliases: []string{"run-build"},
			Usage:   "run a build by id",
			Flags:   subcommandFlags(buildIDFlag, binaryPathFlag),
			Action: func(cCtx *cli.Context) error {
				zap.L().Debug("running a build...")
				build := OneBuildConfigFrom(cCtx)

				filePath := cCtx.String("binary-path")
				file, err := archive.RequireTGZ(filePath)
				if err != nil {
					return fmt.Errorf("no tgz file available for run: %w", err)
				}

				zap.L().Debug("using archive file", zap.Any("file", file))

				res, err := build.SDK.BuildV2.RunBuild(
					cCtx.Context,
					build.BuildID,
					operations.RunBuildRequestBody{
						File: operations.RunBuildFile{
							FileName: file.Name,
							Content:  file.Content,
						},
					},
					build.AppID,
				)
				if err != nil {
					return fmt.Errorf("failed to run build: %w", err)
				}

				return output.As(build.OutputType, &DefaultResult{
					Success: true,
					Message: "Build ran successfully",
					Code:    res.StatusCode,
				})
			},
		},
		{
			Name:    deleteCommandName,
			Aliases: []string{"delete-build"},
			Usage:   "delete a build",
			Flags:   subcommandFlags(buildIDFlag),
			Action: func(cCtx *cli.Context) error {
				zap.L().Debug("deleting a build...")
				build := OneBuildConfigFrom(cCtx)

				res, err := build.SDK.BuildV2.DeleteBuild(cCtx.Context, build.BuildID, build.AppID)
				if err != nil {
					return fmt.Errorf("failed to delete build: %w", err)
				}

				return output.As(build.OutputType, &DefaultResult{
					Success: true,
					Message: "Build deleted successfully",
					Code:    res.StatusCode,
				})
			},
		},
	},
}

func buildFlagEnvVar(name string) []string {
	return []string{fmt.Sprintf("%s%s", buildFlagEnvVarPrefix, name)}
}

var (
	buildFlagEnvVarPrefix = fmt.Sprintf("%s%s", globalFlagEnvVarPrefix, "BUILD_")

	buildIDFlag = &cli.IntFlag{
		Name:     "build-id",
		Aliases:  []string{"b"},
		EnvVars:  buildFlagEnvVar("ID"),
		Usage:    "the ID of the build in Hathora",
		Required: true,
	}

	buildTagFlag = &cli.StringFlag{
		Name:     "build-tag",
		Aliases:  []string{"bt"},
		EnvVars:  buildFlagEnvVar("TAG"),
		Usage:    "tag to associate an external version with a build",
		Required: true,
	}

	binaryPathFlag = &cli.StringFlag{
		Name:     "binary-path",
		Aliases:  []string{"bp"},
		EnvVars:  buildFlagEnvVar("BINARY_PATH"),
		Usage:    "path to the built game server binary",
		Required: true,
	}
)

var (
	buildConfigKey = struct{}{}
)

type BuildConfig struct {
	*GlobalConfig
	SDK   *sdk.SDK
	AppID *string
}

func (c *BuildConfig) Load(cCtx *cli.Context) {
	c.GlobalConfig = GlobalConfigFrom(cCtx)
	c.SDK = setup.SDK(c.Token, c.BaseURL, c.Verbosity)
}

func BuildConfigFrom(cCtx *cli.Context) *BuildConfig {
	return ConfigFromCLI[*BuildConfig](buildConfigKey, cCtx)
}

var (
	oneBuildConfigKey = struct{}{}
)

type OneBuildConfig struct {
	*BuildConfig
	BuildID int
}

func (c *OneBuildConfig) Load(cCtx *cli.Context) {
	c.BuildConfig = BuildConfigFrom(cCtx)
	c.BuildID = cCtx.Int(buildIDFlag.Name)
}

func OneBuildConfigFrom(cCtx *cli.Context) *OneBuildConfig {
	return ConfigFromCLI[*OneBuildConfig](oneBuildConfigKey, cCtx)
}
