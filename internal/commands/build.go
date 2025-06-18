package commands

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync/atomic"

	"github.com/urfave/cli/v3"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/hathora/ci/internal/archive"
	"github.com/hathora/ci/internal/commands/altsrc"
	"github.com/hathora/ci/internal/httputil"
	"github.com/hathora/ci/internal/output"
	"github.com/hathora/ci/internal/setup"

	sdk "github.com/hathora/cloud-sdk-go/hathoracloud"
	"github.com/hathora/cloud-sdk-go/hathoracloud/models/components"
)

type etagPart struct {
	partNumber int
	etag       string
}

const buildFlagEnvVarPrefix = globalFlagEnvVarPrefix + "BUILD_"

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

				res, err := build.SDK.BuildsV3.GetBuild(ctx, build.BuildID, nil)
				if err != nil {
					return fmt.Errorf("failed to get build info: %w", err)
				}

				return build.Output.Write(res, os.Stdout)
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

				res, err := build.SDK.BuildsV3.GetBuilds(ctx, nil)
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
			Flags:   subcommandFlags(buildTagFlag, buildIDFlag, fileFlag, hideUploadProgressFlag),
			Action: func(ctx context.Context, cmd *cli.Command) error {
				build, err := CreateBuildConfigFrom(cmd)
				if err != nil {
					//nolint:errcheck
					cli.ShowSubcommandHelp(cmd)
					return err
				}
				created, err := doBuildCreate(ctx, build.SDK, build.BuildTag, build.BuildID, build.FilePath, build.HideUploadProgress)
				if err != nil {
					return err
				}

				if created.Status == components.BuildStatusFailed {
					return fmt.Errorf("Build failed")
				}

				return build.Output.Write(created, os.Stdout)
			},
		},
		{
			Name:  dockerCommandName,
			Usage: "create a build using a docker image in an external registry",
			Flags: subcommandFlags(buildTagFlag, buildIDFlag, registryImageFlag, registryFQDNFlag, registryAuthFlag),
			Action: func(ctx context.Context, cmd *cli.Command) error {
				build, err := CreateRegistryBuildConfigFrom(cmd)
				if err != nil {
					//nolint:errcheck
					cli.ShowSubcommandHelp(cmd)
					return err
				}
				created, err := doBuildCreateFromExternalRegstry(ctx, build.SDK, build.BuildTag, build.BuildID, build.Image, build.Registry, build.RegistryAuth)
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

				_, err = build.SDK.BuildsV3.DeleteBuild(ctx, build.BuildID, nil)
				if err != nil {
					return fmt.Errorf("failed to delete build: %w", err)
				}

				return build.Output.Write(&DefaultResult{
					Success: true,
					Message: "Build deleted successfully",
					Code:    200,
				}, os.Stdout)
			},
		},
	},
}

func newRegistryConfig(registry, image, auth string) components.RegistryConfig {
	out := components.RegistryConfig{
		Image: image,
	}
	if registry != "" {
		out.RegistryURL = sdk.String(registry)
	}
	if auth != "" {
		out.Auth = sdk.String(auth)
	}
	return out
}

// ParseRegistryAndImage returns the parsed registry hostname and image reference.
// This allows for the registry to be part of the image or separate, and either work.
func ParseRegistryAndImage(registry, image, auth string) components.RegistryConfig {
	if registry != "" || !strings.Contains(image, "/") {
		return newRegistryConfig(registry, image, auth)
	}
	parts := strings.SplitN(image, "/", 2)
	for _, b := range parts[0] {
		if b == ':' || b == '.' {
			// per docker urlparse rules, this indicates the image reference is
			// actually an image url with parts[0] being the registry hostname
			// and parts[1] being the image reference
			return newRegistryConfig(parts[0], parts[1], auth)
		}
	}
	// at this point, it's a docker hub assumed image with a namespace in the
	return newRegistryConfig(registry, image, auth)
}

func doBuildCreateFromExternalRegstry(ctx context.Context,
	hathora *sdk.HathoraCloud, buildTag, buildId, image, registry,
	token string) (*components.BuildV3, error) {
	cfg := ParseRegistryAndImage(registry, image, token)

	var params components.CreateBuildV3Params
	if buildTag != "" {
		params.BuildTag = sdk.String(buildTag)
	}
	if buildId != "" {
		params.BuildID = sdk.String(buildId)
	}
	createRes, err := hathora.BuildsV3.CreateBuildRegistry(ctx, params, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create build: %w", err)
	}

	runRes, err := hathora.BuildsV3.RunBuildRegistry(
		ctx,
		createRes.BuildID,
		cfg,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to run build: %w", err)
	}

	zap.L().Debug("streaming build output to console...")
	err = output.StreamOutput(runRes, os.Stderr)
	if err != nil {
		zap.L().Error("failed to stream output to console", zap.Error(err))
	}

	infoRes, err := hathora.BuildsV3.GetBuild(
		ctx,
		createRes.BuildID,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve build info: %w", err)
	}

	return infoRes, nil
}

func doBuildCreate(ctx context.Context, hathora *sdk.HathoraCloud, buildTag, buildId, filePath string, hideUploadProgress bool) (*components.BuildV3, error) {
	file, err := archive.RequireTGZ(filePath)
	if err != nil {
		return nil, fmt.Errorf("no build file available for run: %w", err)
	}
	fileSize := int64(len(file.Content))
	params := components.CreateMultipartBuildParams{BuildSizeInBytes: float64(fileSize)}
	if buildTag != "" {
		params.BuildTag = sdk.String(buildTag)
	}
	if buildId != "" {
		params.BuildID = sdk.String(buildId)
	}

	createRes, err := hathora.BuildsV3.CreateBuild(
		ctx,
		params,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create a build: %w", err)
	}

	if createRes == nil {
		return nil, fmt.Errorf("no build object in response")
	}

	globalUploadProgress := atomic.Int64{}

	var etagParts = make([]etagPart, len(createRes.UploadParts))
	var eg errgroup.Group
	for i, uploadPart := range createRes.UploadParts {
		partNum := int64(uploadPart.PartNumber)
		reqURL := uploadPart.PutRequestURL
		index := i
		eg.Go(func() error {
			maxChunkSize := int64(createRes.MaxChunkSize)

			start := maxChunkSize * (partNum - 1)
			end := min(partNum*maxChunkSize, fileSize)
			size := end - start
			buf := make([]byte, size)

			copy(buf, file.Content[start:])
			etag, err := uploadFileToS3(reqURL, buf, &globalUploadProgress, fileSize, hideUploadProgress)
			if err != nil {
				fmt.Printf("failed to upload part %d: %v\n", partNum, err)
			}
			etagParts[index] = etagPart{partNumber: int(partNum), etag: etag}

			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return nil, fmt.Errorf("failed to upload parts: %w", err)
	}

	resp, err := http.Post(createRes.CompleteUploadPostRequestURL, httputil.ValueApplicationXML, bytes.NewBufferString(createEtagXML(etagParts...)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("\nComplete multipart upload failed with status: %s\n", resp.Status)
	} else {
		zap.L().Debug("Complete multiplart upload succeeded")
	}

	runRes, err := hathora.BuildsV3.RunBuild(
		ctx,
		createRes.BuildID,
		nil,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to run build: %w", err)
	}

	zap.L().Debug("streaming build output to console...")
	err = output.StreamOutput(runRes, os.Stderr)
	if err != nil {
		zap.L().Error("failed to stream output to console", zap.Error(err))
	}

	infoRes, err := hathora.BuildsV3.GetBuild(
		ctx,
		createRes.BuildID,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve build info: %w", err)
	}

	return infoRes, nil
}

func createEtagXML(etags ...etagPart) string {
	const header = "<CompleteMultipartUpload>"
	const footer = "</CompleteMultipartUpload>"

	var builder strings.Builder
	builder.WriteString(header)

	for _, etag := range etags {
		builder.WriteString(fmt.Sprintf(`
	<Part>
		<PartNumber>%d</PartNumber>
		<ETag>%s</ETag>
	</Part>`, etag.partNumber, etag.etag))
	}

	builder.WriteString(footer)
	return builder.String()
}
func buildFlagEnvVar(name string) string {
	return buildFlagEnvVarPrefix + name
}

var (
	buildIDFlag = &cli.StringFlag{
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
	hideUploadProgressFlag = &cli.BoolFlag{
		Name:    "hide-upload-progress",
		Aliases: []string{"hup"},
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(buildFlagEnvVar("HIDE_UPLOAD_PROGRESS")),
			altsrc.ConfigFile(configFlag.Name, "build.hide_upload_progress"),
		),
		Usage:      "hide the upload progress percentage from output",
		Category:   "Build:",
		Persistent: true,
	}

	registryImageFlag = &cli.StringFlag{
		Name:     "image",
		Required: true,
		Aliases:  []string{"i"},
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(buildFlagEnvVar("IMAGE")),
			altsrc.ConfigFile(configFlag.Name, "build.image_url"),
		),
		Category: "Build:",
		Usage:    "the docker `<image>` tag or url for the external image",
	}

	registryFQDNFlag = &cli.StringFlag{
		Name:    "registry",
		Aliases: []string{"r"},
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(buildFlagEnvVar("REGISTRY")),
			altsrc.ConfigFile(configFlag.Name, "build.registry"),
		),
		Category: "Build:",
		Usage:    "optional `<registry>` FQDN hosting the server image; can be in the image flag instead",
	}

	registryAuthFlag = &cli.StringFlag{
		Name: "auth",
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(buildFlagEnvVar("REGISTRY_AUTH")),
			altsrc.ConfigFile(configFlag.Name, "build.auth"),
		),
		Category: "Build:",
		Usage:    "optional `<auth>` credential for authenticating with a private registry",
	}
)

var (
	buildConfigKey = "commands.BuildConfig.DI"
)

type BuildConfig struct {
	*GlobalConfig
	SDK    *sdk.HathoraCloud
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
	var build components.BuildV3
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

const (
	createBuildConfigKey = "commands.CreateBuildConfig.DI"
)

type CreateBuildConfig struct {
	*BuildConfig
	BuildTag           string
	BuildID            string
	FilePath           string
	HideUploadProgress bool
}

var _ LoadableConfig = (*CreateBuildConfig)(nil)

func (c *CreateBuildConfig) Load(cmd *cli.Command) error {
	build, err := BuildConfigFrom(cmd)
	if err != nil {
		return err
	}
	c.BuildConfig = build
	c.BuildTag = cmd.String(buildTagFlag.Name)
	c.BuildID = cmd.String(buildIDFlag.Name)
	c.FilePath = cmd.String(fileFlag.Name)
	c.HideUploadProgress = cmd.Bool(hideUploadProgressFlag.Name)
	c.Log = c.Log.With(zap.String("build.tag", c.BuildTag)).With(zap.String("build.id", c.BuildID))
	return nil
}

func (c *CreateBuildConfig) New() LoadableConfig {
	return &CreateBuildConfig{}
}

func CreateBuildConfigFrom(cmd *cli.Command) (*CreateBuildConfig, error) {
	return ConfigFromCLI[*CreateBuildConfig](createBuildConfigKey, cmd)
}

const (
	createRegistryBuildConfigKey = "commands.CreateRegistryBuildConfig.DI"
)

type CreateRegistryBuildConfig struct {
	*BuildConfig
	BuildTag     string
	BuildID      string
	Image        string
	Registry     string
	RegistryAuth string
}

var _ LoadableConfig = (*CreateRegistryBuildConfig)(nil)

func (c *CreateRegistryBuildConfig) Load(cmd *cli.Command) error {
	build, err := BuildConfigFrom(cmd)
	if err != nil {
		return err
	}
	c.BuildConfig = build
	c.BuildTag = cmd.String(buildTagFlag.Name)
	c.BuildID = cmd.String(buildIDFlag.Name)
	c.Image = cmd.String(registryImageFlag.Name)
	c.Registry = cmd.String(registryFQDNFlag.Name)
	c.RegistryAuth = cmd.String(registryAuthFlag.Name)
	c.Log = c.Log.With(zap.String("build.tag", c.BuildTag)).With(zap.String("build.id", c.BuildID))
	return nil
}

func (c *CreateRegistryBuildConfig) New() LoadableConfig {
	return &CreateRegistryBuildConfig{}
}

func CreateRegistryBuildConfigFrom(cmd *cli.Command) (*CreateRegistryBuildConfig, error) {
	return ConfigFromCLI[*CreateRegistryBuildConfig](createRegistryBuildConfigKey, cmd)
}

var (
	oneBuildConfigKey = "commands.OneBuildConfig.DI"
)

type OneBuildConfig struct {
	*BuildConfig
	BuildID string
}

var _ LoadableConfig = (*OneBuildConfig)(nil)

func (c *OneBuildConfig) Load(cmd *cli.Command) error {
	build, err := BuildConfigFrom(cmd)
	if err != nil {
		return err
	}
	c.BuildConfig = build
	c.BuildID = cmd.String(buildIDFlag.Name)
	c.Log = c.Log.With(zap.String("build.id", c.BuildID))
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
	globalTotal          int64
	globalUploadProgress *atomic.Int64
	hideUploadProgress   bool
}

func (pr *progressReaderType) Read(p []byte) (int, error) {
	n, err := pr.reader.Read(p)
	if err == io.EOF {
		return n, err
	}
	if err != nil {
		return n, fmt.Errorf("failed to read data: %w", err)
	}
	if n > 0 && !pr.hideUploadProgress {
		pr.globalUploadProgress.Add(int64(n))
		loaded := pr.globalUploadProgress.Load()
		percentage := float64(loaded*100) / float64(pr.globalTotal)
		os.Stderr.WriteString(fmt.Sprintf("Upload progress: %.2f%% (%d/%d bytes)\r", percentage, loaded, pr.globalTotal))
	}
	return n, err
}

func uploadFileToS3(preSignedURL string, byteBuffer []byte, globalUploadProgress *atomic.Int64, globalTotal int64, hideUploadProgress bool) (string, error) {
	requestBody := bytes.NewReader(byteBuffer)
	progressReader := &progressReaderType{
		reader:               requestBody,
		globalTotal:          globalTotal,
		globalUploadProgress: globalUploadProgress,
		hideUploadProgress:   hideUploadProgress,
	}
	req, err := http.NewRequest(http.MethodPut, preSignedURL, progressReader)
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("failed to create request: %v", err))
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set(httputil.NameContentType, httputil.ValueApplicationOctetStream)
	req.ContentLength = int64(requestBody.Len())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("failed to upload part: %v", err))
		return "", fmt.Errorf("failed to upload part: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		os.Stderr.WriteString(fmt.Sprintf("upload failed with status code: %d, response: %s", resp.StatusCode, body))
		return "", fmt.Errorf("upload failed with status code: %d", resp.StatusCode)
	}

	etag := resp.Header.Get("ETag")
	if etag == "" {
		os.Stderr.WriteString("etag header not found in response")
		return "", fmt.Errorf("etag header not found in response")
	}

	return etag, nil
}
