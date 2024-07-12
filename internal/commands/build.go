package commands

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/urfave/cli/v3"
	"go.uber.org/zap"

	"github.com/hathora/ci/internal/archive"
	"github.com/hathora/ci/internal/commands/altsrc"
	"github.com/hathora/ci/internal/output"
	"github.com/hathora/ci/internal/sdk"
	"github.com/hathora/ci/internal/sdk/models/operations"
	"github.com/hathora/ci/internal/sdk/models/shared"
	"github.com/hathora/ci/internal/setup"
	"github.com/hathora/ci/internal/workaround"
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
					//nolint:errcheck
					cli.ShowSubcommandHelp(cmd)
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
					//nolint:errcheck
					cli.ShowSubcommandHelp(cmd)
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
			Flags:   subcommandFlags(buildTagFlag, fileFlag),
			Action: func(ctx context.Context, cmd *cli.Command) error {
				build, err := CreateBuildConfigFrom(cmd)
				if err != nil {
					//nolint:errcheck
					cli.ShowSubcommandHelp(cmd)
					return err
				}
				created, err := doBuildCreate(ctx, build.Log, build.SDK, build.AppID, build.BuildTag, build.FilePath)
				if err != nil {
					return err
				}

				return build.Output.Write(created, os.Stdout)
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
					//nolint:errcheck
					cli.ShowSubcommandHelp(cmd)
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

func doBuildCreate(ctx context.Context, logger *zap.Logger, hathora *sdk.SDK, appID *string, buildTag, filePath string) (*shared.Build, error) {
	createRes, err := hathora.BuildV2.CreateBuildWithUploadURL(
		ctx,
		shared.CreateBuildParams{
			BuildTag: sdk.String(buildTag),
		},
		appID,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create a build: %w", err)
	}

	file, err := archive.RequireTGZ(filePath)
	if err != nil {
		return nil, fmt.Errorf("no build file available for run: %w", err)
	}

	if createRes.BuildWithUploadURL == nil {
		return nil, fmt.Errorf("no build object in response")
	}

	err = uploadToUrl(createRes.BuildWithUploadURL.UploadURL, createRes.BuildWithUploadURL.UploadBodyParams, file.Path, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to upload file: %w", err)
	}

	runRes, err := hathora.BuildV2.RunBuild(
		ctx,
		createRes.BuildWithUploadURL.BuildID,
		operations.RunBuildRequestBody{
			File: operations.RunBuildFile{
				FileName: file.Name,
				Content:  file.Content,
			},
		},
		appID,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to run build: %w", err)
	}

	zap.L().Debug("streaming build output to console...")
	err = output.StreamOutput(runRes.Stream, os.Stderr)
	if err != nil {
		zap.L().Error("failed to stream output to console", zap.Error(err))
	}

	infoRes, err := hathora.BuildV2.GetBuildInfo(
		ctx,
		createRes.BuildWithUploadURL.BuildID,
		appID,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve build info: %w", err)
	}

	return infoRes.Build, nil
}

func buildFlagEnvVar(name string) string {
	return buildFlagEnvVarPrefix + name
}

var (
	buildFlagEnvVarPrefix = globalFlagEnvVarPrefix + "BUILD_"

	buildIDFlag = &workaround.IntFlag{
		Name:    "build-id",
		Aliases: []string{"b"},
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(buildFlagEnvVar("ID")),
			altsrc.ConfigFile(configFlag.Name, "build.id"),
		),
		Usage:      "the `<id>` of the build in Hathora",
		Category:   "Build:",
		Persistent: true,
	}

	buildTagFlag = &cli.StringFlag{
		Name:    "build-tag",
		Aliases: []string{"bt"},
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(buildFlagEnvVar("TAG")),
			altsrc.ConfigFile(configFlag.Name, "build.tag"),
		),
		Category: "Build:",
		Usage:    "`<tag>` or external version to associate with the build",
	}

	fileFlag = &cli.StringFlag{
		Name:    "file",
		Aliases: []string{"f"},
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(buildFlagEnvVar("FILE")),
			altsrc.ConfigFile(configFlag.Name, "build.file"),
		),
		Usage:     "`<path>` to the built game server binary or archive",
		Category:  "Build:",
		TakesFile: true,
		Value:     ".", // default to current working directory
	}
)

var (
	buildConfigKey = "commands.BuildConfig.DI"
)

type BuildConfig struct {
	*GlobalConfig
	SDK    *sdk.SDK
	Output output.FormatWriter
}

var _ LoadableConfig = (*BuildConfig)(nil)

func (c *BuildConfig) Load(cmd *cli.Command) error {
	global, err := GlobalConfigFrom(cmd)
	if err != nil {
		return err
	}
	c.GlobalConfig = global
	c.SDK = setup.SDK(c.Token, c.BaseURL, c.Verbosity)
	var build shared.Build
	output, err := OutputFormatterFor(cmd, build)
	if err != nil {
		return err
	}
	c.Output = output
	return nil
}

func (c *BuildConfig) New() LoadableConfig {
	return &BuildConfig{}
}

func BuildConfigFrom(cmd *cli.Command) (*BuildConfig, error) {
	return ConfigFromCLI[*BuildConfig](buildConfigKey, cmd)
}

var (
	createBuildConfigKey = "commands.CreateBuildConfig.DI"
)

type CreateBuildConfig struct {
	*BuildConfig
	BuildTag string
	FilePath string
}

var _ LoadableConfig = (*CreateBuildConfig)(nil)

func (c *CreateBuildConfig) Load(cmd *cli.Command) error {
	build, err := BuildConfigFrom(cmd)
	if err != nil {
		return err
	}
	c.BuildConfig = build
	c.BuildTag = cmd.String(buildTagFlag.Name)
	c.FilePath = cmd.String(fileFlag.Name)
	c.Log = c.Log.With(zap.String("build.tag", c.BuildTag))
	return nil
}

func (c *CreateBuildConfig) New() LoadableConfig {
	return &CreateBuildConfig{}
}

func CreateBuildConfigFrom(cmd *cli.Command) (*CreateBuildConfig, error) {
	return ConfigFromCLI[*CreateBuildConfig](createBuildConfigKey, cmd)
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

func uploadToUrl(uploadUrl string, uploadBodyParams []shared.UploadBodyParams, filePath string, logger *zap.Logger) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	fileSize := fileInfo.Size()

	var requestBody bytes.Buffer
	multipartWriter := multipart.NewWriter(&requestBody)

	for _, param := range uploadBodyParams {
		_ = multipartWriter.WriteField(param.Key, param.Value)
	}

	fileWriter, err := multipartWriter.CreateFormFile("file", fileInfo.Name())
	if err != nil {
		return err
	}

	progressReader := &progressReaderType{
		reader: file,
		total:  fileSize,
		callback: func(percentage float64, loaded int64, total int64) {
			progressLine := fmt.Sprintf("Upload progress: %.2f%% (%d/%d bytes)\r", percentage, loaded, total)
			logger.Info(progressLine)
		},
	}

	_, err = io.Copy(fileWriter, progressReader)
	if err != nil {
		return err
	}
	logger.Info("\n")

	err = multipartWriter.Close()
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", uploadUrl, &requestBody)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("upload failed with status: %s", resp.Status)
	}

	return nil
}

type progressReaderType struct {
	reader   io.Reader
	total    int64
	read     int64
	callback func(percentage float64, loaded int64, total int64)
}

func (pr *progressReaderType) Read(p []byte) (int, error) {
	n, err := pr.reader.Read(p)
	if n > 0 {
		pr.read += int64(n)
		percentage := float64(pr.read) / float64(pr.total) * 100
		pr.callback(percentage, pr.read, pr.total)
	}
	return n, err
}
