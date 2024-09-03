package commands

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync/atomic"

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

				res, err := build.SDK.BuildsV2.GetBuildInfoV2Deprecated(ctx, build.BuildID, build.AppID)
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

				res, err := build.SDK.BuildsV2.GetBuildsV2Deprecated(ctx, build.AppID)
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
				created, err := doBuildCreate(ctx, build.SDK, build.AppID, build.BuildTag, build.FilePath)
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

				res, err := build.SDK.BuildsV2.DeleteBuildV2Deprecated(ctx, build.BuildID, build.AppID)
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

func doBuildCreate(ctx context.Context, hathora *sdk.SDK, appID *string, buildTag, filePath string) (*shared.Build, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to get file info: %w", err)
	}

	createRes, err := hathora.BuildsV2.CreateWithMultipartUploadsV2Deprecated(
		ctx,
		shared.CreateMultipartBuildParams{
			BuildTag:         sdk.String(buildTag),
			BuildSizeInBytes: float64(fileInfo.Size()),
		},
		appID,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create a build: %w", err)
	}

	_, err = archive.RequireTGZ(filePath)
	if err != nil {
		return nil, fmt.Errorf("no build file available for run: %w", err)
	}

	if createRes.BuildWithMultipartUrls == nil {
		return nil, fmt.Errorf("no build object in response")
	}

	etagsForParts := make(map[int64](chan string), len(createRes.BuildWithMultipartUrls.UploadParts))

	globalUploadProgress := atomic.Int64{}

	for _, part := range createRes.BuildWithMultipartUrls.UploadParts {
		etagChan := make(chan string)
		partNumber := int64(part.PartNumber)
		etagsForParts[partNumber] = etagChan
		maxChunkSize := int64(createRes.BuildWithMultipartUrls.MaxChunkSize)

		startByteForPart := (partNumber - 1) * maxChunkSize
		endByteForPart := min(partNumber*maxChunkSize, fileInfo.Size())

		partSize := endByteForPart - startByteForPart
		partBuffer := make([]byte, partSize)

		_, err := file.ReadAt(partBuffer, startByteForPart)
		if err != nil {
			return nil, fmt.Errorf("failed to read file: %w", err)
		}
		go uploadFileToS3(part.PutRequestURL, partBuffer, etagChan, &globalUploadProgress, fileInfo.Size())
	}

	xmlBody := "<CompleteMultipartUpload>"
	for partNumber, etagChan := range etagsForParts {
		xmlBody += fmt.Sprintf(`
			<Part>
				<PartNumber>%d</PartNumber>
				<ETag>%s</ETag>
			</Part>`, int(partNumber), <-etagChan)
	}
	xmlBody += "</CompleteMultipartUpload>"

	resp, err := http.Post(createRes.BuildWithMultipartUrls.CompleteUploadPostRequestURL, "application/xml", bytes.NewBufferString(xmlBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Complete multipart upload failed with status: %s\n", resp.Status)
	} else {
		fmt.Println("Complete multipart upload succeeded.")
	}
	// err = uploadToUrl(createRes.BuildWithMultipartUrls.UploadParts, int(createRes.BuildWithMultipartUrls.MaxChunkSize), createRes.BuildWithMultipartUrls.CompleteUploadPostRequestURL, osFile)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to upload file: %w", err)
	// }

	runRes, err := hathora.BuildsV2.RunBuildV2Deprecated(
		ctx,
		createRes.BuildWithMultipartUrls.BuildID,
		operations.RunBuildV2DeprecatedRequestBody{},
		appID,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to run build: %w", err)
	}

	zap.L().Debug("streaming build output to console...")
	err = output.StreamOutput(runRes.ResponseStream, os.Stderr)
	if err != nil {
		zap.L().Error("failed to stream output to console", zap.Error(err))
	}

	infoRes, err := hathora.BuildsV2.GetBuildInfoV2Deprecated(
		ctx,
		createRes.BuildWithMultipartUrls.BuildID,
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

type progressReaderType struct {
	reader               io.Reader
	total                int64
	globalUploadProgress *atomic.Int64
}

func (pr *progressReaderType) Read(p []byte) (int, error) {
	n, err := pr.reader.Read(p)
	if err != nil && err == io.EOF {
		os.Stderr.WriteString("Upload complete\n")
	}
	if n > 0 {
		pr.globalUploadProgress.Add(int64(n))
		loaded := pr.globalUploadProgress.Load()
		percentage := float64(loaded*100) / float64(pr.total)
		os.Stderr.WriteString(fmt.Sprintf("Upload progress: %.2f%% (%d/%d bytes)\r", percentage, loaded, pr.total))
	}
	return n, err
}

func uploadFileToS3(presignedUrl string, byteBuffer []byte, etagChan chan string, globalUploadProgress *atomic.Int64, globalTotal int64) {
	requestBody := bytes.NewReader(byteBuffer)
	progressReader := &progressReaderType{
		reader:               requestBody,
		total:                globalTotal,
		globalUploadProgress: globalUploadProgress,
	}
	req, err := http.NewRequest("PUT", presignedUrl, progressReader)
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("failed to create request: %v", err))
		etagChan <- "error"
		return
	}
	req.Header.Set("Content-Type", "application/octet-stream")
	req.ContentLength = int64(requestBody.Len())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("failed to upload part: %v", err))
		etagChan <- "error"
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		os.Stderr.WriteString(fmt.Sprintf("upload failed with status code: %d, response: %s", resp.StatusCode, body))
		etagChan <- "error"
		return
	}

	etag := resp.Header.Get("ETag")
	if etag == "" {
		os.Stderr.WriteString("ETag header not found in response")
		etagChan <- "error"
	} else {
		etagChan <- etag
	}
}
