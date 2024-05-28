package commands

import (
	"context"
	"fmt"
	"os"

	"github.com/hathora/ci/internal/archive"
	"github.com/hathora/ci/internal/commands/altsrc"
	"github.com/hathora/ci/internal/sdk"
	"github.com/hathora/ci/internal/sdk/models/operations"
	"github.com/hathora/ci/internal/sdk/models/shared"
	"github.com/hathora/ci/internal/setup"
	"github.com/urfave/cli/v3"
	"go.uber.org/zap"
)

var Build = &cli.Command{
	Name:  "build",
	Usage: "options for builds",
	Commands: []*cli.Command{
		{
			Name:    infoCommandName,
			Aliases: []string{"get-build-info"},
			Usage:   "get a build",
			Flags:   subcommandFlags(buildIDFlag),
			Action: func(ctx context.Context, cmd *cli.Command) error {
				build, err := OneBuildConfigFrom(cmd)
				if err != nil {
					return err
				}
				build.Log.Debug("getting build info...")

				res, err := build.SDK.BuildV2.GetBuildInfo(ctx, build.BuildID, build.AppID)
				if err != nil {
					return fmt.Errorf("failed to get build info: %w", err)
				}

				return build.Output.Write(res.Build, os.Stdout)
			},
		},
		{
			Name:    listCommandName,
			Aliases: []string{"get-builds"},
			Usage:   "get all builds",
			Flags:   subcommandFlags(),
			Action: func(ctx context.Context, cmd *cli.Command) error {
				build, err := BuildConfigFrom(cmd)
				if err != nil {
					return err
				}
				build.Log.Debug("getting all builds...")

				res, err := build.SDK.BuildV2.GetBuilds(ctx, build.AppID)
				if err != nil {
					return fmt.Errorf("failed to get builds: %w", err)
				}

				if len(res.Builds) == 0 {
					return fmt.Errorf("no builds found")
				}

				return build.Output.Write(res.Builds, os.Stdout)
			},
		},
		{
			Name:    createCommandName,
			Aliases: []string{"create-build"},
			Usage:   "create a build",
			Flags:   subcommandFlags(buildTagFlag),
			Action: func(ctx context.Context, cmd *cli.Command) error {
				build, err := BuildConfigFrom(cmd)
				if err != nil {
					return err
				}
				build.Log.Debug("creating a build...")

				buildTag := sdk.String(cmd.String(buildTagFlag.Name))

				res, err := build.SDK.BuildV2.CreateBuild(
					ctx,
					shared.CreateBuildParams{
						BuildTag: buildTag,
					},
					build.AppID,
				)
				if err != nil {
					return fmt.Errorf("failed to create a build: %w", err)
				}

				return build.Output.Write(res.Build, os.Stdout)
			},
		},
		{
			Name:    "run",
			Aliases: []string{"run-build"},
			Usage:   "run a build by id",
			Flags:   subcommandFlags(buildIDFlag, fileFlag),
			Action: func(ctx context.Context, cmd *cli.Command) error {
				build, err := OneBuildConfigFrom(cmd)
				if err != nil {
					return err
				}
				build.Log.Debug("running a build...")

				filePath := cmd.String(fileFlag.Name)
				file, err := archive.RequireTGZ(filePath)
				if err != nil {
					return fmt.Errorf("no tgz file available for run: %w", err)
				}

				build.Log.Debug("using archive file", zap.Any("file", file))

				res, err := build.SDK.BuildV2.RunBuild(
					ctx,
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

				return build.Output.Write(&DefaultResult{
					Success: true,
					Message: "Build ran successfully",
					Code:    res.StatusCode,
				}, os.Stdout)
			},
		},
		{
			Name:    deleteCommandName,
			Aliases: []string{"delete-build"},
			Usage:   "delete a build",
			Flags:   subcommandFlags(buildIDFlag),
			Action: func(ctx context.Context, cmd *cli.Command) error {
				build, err := OneBuildConfigFrom(cmd)
				if err != nil {
					return err
				}
				build.Log.Debug("deleting a build...")

				res, err := build.SDK.BuildV2.DeleteBuild(ctx, build.BuildID, build.AppID)
				if err != nil {
					return fmt.Errorf("failed to delete build: %w", err)
				}

				return build.Output.Write(&DefaultResult{
					Success: true,
					Message: "Build deleted successfully",
					Code:    res.StatusCode,
				}, os.Stdout)
			},
		},
	},
}

func buildFlagEnvVar(name string) string {
	return buildFlagEnvVarPrefix + name
}

var (
	buildFlagEnvVarPrefix = globalFlagEnvVarPrefix + "BUILD_"

	buildIDFlag = &cli.IntFlag{
		Name:    "build-id",
		Aliases: []string{"b"},
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(buildFlagEnvVar("ID")),
			altsrc.File(configFlag.Name, "build.id"),
		),
		Usage:      "the ID of the build in Hathora",
		Persistent: true,
	}

	buildTagFlag = &cli.StringFlag{
		Name:    "build-tag",
		Aliases: []string{"bt"},
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(buildFlagEnvVar("TAG")),
		),
		Usage: "tag to associate an external version with a build",
	}

	fileFlag = &cli.StringFlag{
		Name:    "file",
		Aliases: []string{"f"},
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(buildFlagEnvVar("FILE")),
		),
		Usage:    "filepath of the built game server binary or archive",
		Required: true,
	}
)

var (
	buildConfigKey = "commands.BuildConfig.DI"
)

type BuildConfig struct {
	*GlobalConfig
	SDK *sdk.SDK
}

var _ LoadableConfig = (*BuildConfig)(nil)

func (c *BuildConfig) Load(cmd *cli.Command) error {
	global, err := GlobalConfigFrom(cmd)
	if err != nil {
		return err
	}
	c.GlobalConfig = global
	c.SDK = setup.SDK(c.Token, c.BaseURL, c.Verbosity)
	return nil
}

func (c *BuildConfig) New() LoadableConfig {
	return &BuildConfig{}
}

func BuildConfigFrom(cmd *cli.Command) (*BuildConfig, error) {
	return ConfigFromCLI[*BuildConfig](buildConfigKey, cmd)
}

var (
	oneBuildConfigKey = "commands.OneBuildConfig.DI"
)

type OneBuildConfig struct {
	*BuildConfig
	BuildID int
}

var _ LoadableConfig = (*OneBuildConfig)(nil)

func (c *OneBuildConfig) Load(cmd *cli.Command) error {
	build, err := BuildConfigFrom(cmd)
	if err != nil {
		return err
	}
	c.BuildConfig = build
	c.BuildID = int(cmd.Int(buildIDFlag.Name))
	c.Log = c.Log.With(zap.Int("build.id", c.BuildID))
	return nil
}

func (c *OneBuildConfig) New() LoadableConfig {
	return &OneBuildConfig{}
}

func OneBuildConfigFrom(cmd *cli.Command) (*OneBuildConfig, error) {
	return ConfigFromCLI[*OneBuildConfig](oneBuildConfigKey, cmd)
}
