package commands

import (
	"github.com/urfave/cli/v3"

	"github.com/hathora/ci/internal/commands/altsrc"
)

var (
	outputFlag = &cli.StringFlag{
		Name:    "output",
		Aliases: []string{"o"},
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(globalFlagEnvVar("OUTPUT")),
			altsrc.ConfigFile(configFlag.Name, "global.output"),
		),
		Usage:       "the `<format>` of the output. Supported values: (json, text, buildIdValue)",
		Value:       "text",
		DefaultText: "text",
		Persistent:  true,
		Category:    "Global:",
	}

	outputPrettyFlag = &cli.BoolFlag{
		Name:       "pretty",
		Usage:      "enable pretty output (json only)",
		Value:      true,
		Category:   "Global:",
		Persistent: true,
	}

	verboseFlag = &cli.BoolFlag{
		Name:       "verbose",
		Aliases:    []string{"v"},
		Usage:      "enable verbose logging",
		Category:   "Global:",
		Persistent: true,
	}

	verbosityFlag = &cli.IntFlag{
		Name: "verbosity",
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(globalFlagEnvVar("VERBOSITY")),
			altsrc.ConfigFile(configFlag.Name, "global.verbosity"),
		),
		Usage:      "set the logging verbosity `<level>` (0-3)",
		Value:      0,
		Category:   "Global:",
		Persistent: true,
	}

	hathoraCloudEndpointFlag = &cli.StringFlag{
		Name: "hathora-cloud-endpoint",
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(globalFlagEnvVar("CLOUD_ENDPOINT")),
			altsrc.ConfigFile(configFlag.Name, "global.cloud-endpoint"),
		),
		Usage:       "override the default API base `<url>`",
		Value:       "https://api.hathora.dev",
		DefaultText: "https://api.hathora.dev",
		Category:    "Global:",
		Persistent:  true,
	}

	tokenFlag = &cli.StringFlag{
		Name:       "token",
		Aliases:    []string{"t"},
		Sources:    cli.EnvVars(globalFlagEnvVar("TOKEN")),
		Usage:      "`<access-token>` for authenticating with the API",
		Category:   "Global:",
		Persistent: true,
	}

	configFlag = &cli.StringFlag{
		Name:       "config",
		Aliases:    []string{"c"},
		Usage:      "`<path>` to the configuration file",
		Category:   "Global:",
		Persistent: true,
	}

	GlobalFlags = []cli.Flag{
		hathoraCloudEndpointFlag,
		tokenFlag,
		outputFlag,
		outputPrettyFlag,
		verboseFlag,
		verbosityFlag,
		configFlag,
	}
)

var (
	globalFlagEnvVarPrefix = "HATHORA_"
)

func globalFlagEnvVar(name string) string {
	return globalFlagEnvVarPrefix + name
}

func subcommandFlags(flags ...cli.Flag) []cli.Flag {
	return append(flags, GlobalFlags...)
}
