package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"go/version"
	"math"
	"net/http"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/urfave/cli/v3"
	"go.uber.org/zap"

	"github.com/hathora/ci/internal/output"
	"github.com/hathora/ci/internal/sdk/models/shared"
)

var (
	infoCommandName   = "info"
	listCommandName   = "list"
	createCommandName = "create"
	deleteCommandName = "delete"
	latestCommandName = "latest"
)

type LoadableConfig interface {
	New() LoadableConfig
	Load(cmd *cli.Command) error
}

func configFromMetadata[T LoadableConfig](key string, md map[string]any) (T, bool) {
	var nilT T
	untyped, hasKey := md[key]
	if !hasKey {
		return nilT, false
	}
	typed, isType := untyped.(T)
	if isType {
		return typed, true
	}

	return nilT, false
}

func ConfigFromCLI[T LoadableConfig](key string, cmd *cli.Command) (T, error) {
	cfg, hasCfg := configFromMetadata[T](key, cmd.Metadata)
	if !hasCfg {
		cfg = cfg.New().(T)
	}
	err := cfg.Load(cmd)
	if err != nil {
		return cfg, err
	}

	cmd.Metadata[key] = cfg
	return cfg, nil
}

type VerbosityConfig struct {
	Verbosity int
	Log       *zap.Logger
}

func (c *VerbosityConfig) Load(cmd *cli.Command) error {
	// we subtract 1 because the flag is counted an additional time for the
	// --verbose alias
	verboseCount := cmd.Count(verboseFlag.Name) - 1
	verbosity := cmd.Int(verbosityFlag.Name)
	c.Verbosity = int(math.Max(float64(verbosity), float64(verboseCount)))
	return nil
}

func (c *VerbosityConfig) New() LoadableConfig {
	return &VerbosityConfig{}
}

var (
	verbosityConfigKey = "commands.VerbosityConfig.DI"
)

func VerbosityConfigFrom(cmd *cli.Command) (*VerbosityConfig, error) {
	cfg, err := ConfigFromCLI[*VerbosityConfig](verbosityConfigKey, cmd)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

type GlobalConfig struct {
	*VerbosityConfig
	Token   string
	BaseURL string
	AppID   *string
}

func (c *GlobalConfig) Load(cmd *cli.Command) error {
	verbosityConfig, err := VerbosityConfigFrom(cmd)
	if err != nil {
		return err
	}
	c.VerbosityConfig = verbosityConfig
	c.Token = cmd.String(tokenFlag.Name)
	if c.Token == "" {
		err = errors.Join(err, missingRequiredFlag(tokenFlag.Name))
	}
	c.BaseURL = cmd.String(hathoraCloudEndpointFlag.Name)
	if c.BaseURL == "" {
		err = errors.Join(err, missingRequiredFlag(hathoraCloudEndpointFlag.Name))
	}
	appID := cmd.String(appIDFlag.Name)
	if appID == "" {
		c.AppID = nil
	} else {
		c.AppID = &appID
	}
	c.Log = zap.L().With(zap.String("app.id", appID))

	return err
}

func (c *GlobalConfig) New() LoadableConfig {
	return &GlobalConfig{}
}

var (
	globalConfigKey = "commands.GlobalConfig.DI"
)

func GlobalConfigFrom(cmd *cli.Command) (*GlobalConfig, error) {
	cfg, err := ConfigFromCLI[*GlobalConfig](globalConfigKey, cmd)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func isCallForHelp(cmd *cli.Command) bool {
	for _, arg := range cmd.Args().Slice() {
		if arg == "--help" || arg == "-h" || arg == "help" {
			return true
		}
	}
	return false
}

func OutputFormatterFor(cmd *cli.Command, outputType any) (output.FormatWriter, error) {
	outputFmt := cmd.String(outputFlag.Name)
	switch output.ParseOutputType(outputFmt) {
	case output.JSON:
		return output.JSONFormat(cmd.Bool(outputPrettyFlag.Name)), nil
	case output.Text:
		return BuildTextFormatter(), nil
	case output.Value:
		fieldName := strings.TrimSuffix(outputFmt, "Value")
		if len(fieldName) == 0 {
			return nil, fmt.Errorf("invalid value format: %s", outputType)
		}
		return output.ValueFormat(outputType, fieldName)
	default:
		return nil, fmt.Errorf("unsupported output type: %s", outputType)
	}
}

func BuildTextFormatter() output.FormatWriter {
	// TODO: Allow commands to register their own formatters so that this one function doesn't have to know the desired format for every type
	var build shared.BuildV3
	var deployment shared.DeploymentV3
	var envVar shared.DeploymentV3Env
	var containerPort shared.ContainerPort
	var timestamp time.Time
	var long int64
	return output.TextFormat(
		output.WithFieldOrder(build,
			"BuildID",
			"BuildTag",
			"CreatedAt",
			"Status",
			"ImageSize",
			"StartedAt",
			"FinishedAt",
		),
		output.WithoutFields(build, "AppID", "RegionalContainerTags", "DeletedAt", "CreatedBy"),
		output.WithFieldOrder(deployment,
			"DeploymentID",
			"BuildID",
			"CreatedAt",
			"IdleTimeoutEnabled",
			"RoomsPerProcess",
			"RequestedCPU",
			"RequestedMemoryMB",
			"DefaultContainerPort",
			"AdditionalContainerPorts",
		),
		output.RenameField(deployment, "RequestedMemoryMB", "RequestedMemory"),
		output.WithPropertyFormatter(deployment, "RequestedMemoryMB", func(f float64) string {
			return humanize.IBytes((uint64)(f * 1024 * 1024))
		}),
		output.WithoutFields(deployment, "AppID", "CreatedBy", "Env"),
		output.WithFormatter(envVar,
			func(e shared.DeploymentV3Env) string {
				return fmt.Sprintf("%s=%s", e.Name, e.Value)
			},
		),
		output.WithFormatter(containerPort,
			func(cp shared.ContainerPort) string {
				return fmt.Sprintf("%s:%d/%s", cp.Name, cp.Port, cp.TransportType)
			},
		),
		output.WithFormatter(timestamp,
			func(t time.Time) string {
				// TODO: consider using this human-friendly time format
				// return humanize.Time(t)
				return t.Format(time.RFC3339)
			},
		),
		// TODO: this should not be how we generally handle all int64s, so we may want support for targeted formatters by field name
		output.WithFormatter(long,
			func(l int64) string {
				return humanize.IBytes(uint64(l))
			},
		),
	)
}

type Release struct {
	TagName     string `json:"tag_name"`
	Name        string `json:"name"`
	PublishedAt string `json:"published_at"`
	URL         string `json:"html_url"`
}

func handleNewVersionAvailable(currentVersion string) {
	url := "https://api.github.com/repos/hathora/ci/releases/latest"

	resp, err := http.Get(url)
	if err != nil {
		zap.L().Warn("unable to fetch latest version number")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		zap.L().Warn("unable to fetch latest version number")
	}

	var release Release
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		zap.L().Warn("unable to decode the latest version number")
	}

	prefixedVersion := "go" + currentVersion
	showLatest := false
	if !version.IsValid(prefixedVersion) {
		zap.L().Warn("You are using a development version of the hathora cli.")
		showLatest = true
	} else if version.Compare("go"+release.TagName, prefixedVersion) > 0 {
		zap.L().Warn("You are using an outdated version of the hathora cli.")
		showLatest = true
	}

	if showLatest {
		zap.L().Warn("Version " + release.TagName + " is available for download.")
	}
}

// NormalizeArgs accepts an args slice (e.g., os.Args) and transforms it to
// collapse any Boolean flags with separate value parameters into a single
// parameter=value arg. This allows users to provide arg lists like:
// --idle-timeout-enabled false
// which would normally ignore the provided value and always register as True,
// as Flag's required syntax for bools is strict: --idle-timeout-enabled=false
// In all cases, a new slice is returned.
func NormalizeArgs(cmd *cli.Command, args []string) []string {
	l := len(args)
	if l == 0 {
		return make([]string, 0)
	}
	// this returns a map of boolean flag names from the command tree
	bfl := getBooleanFlags(cmd)
	if len(bfl) == 0 {
		out := make([]string, l)
		copy(out, args)
		return out
	}
	// this checks if there are any boolean flags provided by the user where
	// the next flag is a value instead of a flag. in those cases, this will
	// merge the elements. so ["--my-boolean", "true"] becomes ["--my-boolean=true"]
	var skipNext bool
	out := make([]string, 0, l)
	for i, arg := range args {
		// this skips whenever the previous arg just merged the current into it
		if skipNext {
			skipNext = false
			continue
		}
		// this adds the current arg to the output
		out = append(out, arg)
		// this moves on when the current arg is not a flag name
		if !strings.HasPrefix(arg, "-") {
			continue
		}
		// this provides the flag name without a leading - or --
		flagName := strings.TrimPrefix(strings.TrimPrefix(arg, "-"), "-")
		// this skips any flag names that are not boolean
		if _, ok := bfl[flagName]; !ok {
			continue
		}
		// this skips when the iterator is on the last arg or whenever it's not
		// the last but another flag name arg immediately follows this one
		if i >= l-1 || strings.HasPrefix(args[i+1], "-") {
			continue
		}
		// the current arg must be a boolean flag name followed by a value arg.
		// this therefore merges them into a single element with an = delimiter
		// by updating the just-appended element and skipping the next element
		out[len(out)-1] = fmt.Sprintf("%s=%s", arg, args[i+1])
		skipNext = true
	}
	return out
}

// getBooleanFlags iterates the command tree and returns a lookup table of all
// boolean flags
func getBooleanFlags(cmd *cli.Command) map[string]any {
	out := make(map[string]any)
	getBooleanFlagsIter(cmd, out)
	return out
}

func getBooleanFlagsIter(cmd *cli.Command, dst map[string]any) {
	// this skips over help, as it is allowed to have a value in the next arg
	// (e.g., --help deploy)
	omit := func(names []string) bool {
		for _, name := range names {
			if name == "help" {
				return true
			}
		}
		return false
	}
	// this iterates and registers the current Command's boolean Flag names
	for _, flag := range cmd.Flags {
		// this skips non-bools
		if _, ok := flag.(*cli.BoolFlag); !ok {
			continue
		}
		// this skips any other explicit omissions
		fns := flag.Names()
		if omit(fns) {
			continue
		}
		// otherwise this add the Flag's names to the output map
		for _, name := range fns {
			dst[name] = nil
		}
	}
	// this sends the passed command's subcommands through this iterator func
	for _, sub := range cmd.Commands {
		getBooleanFlagsIter(sub, dst)
	}
}
