package commands

import (
	"github.com/hathora/ci/internal/commands/altsrc"

	"github.com/urfave/cli/v3"
)

var (
	outputTypeFlag = &cli.StringFlag{
		Name:    "output",
		Aliases: []string{"o"},
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(globalFlagEnvVar("OUTPUT")),
			altsrc.File(configFlag.Name, "global.output"),
		),
		Usage:      "the format of the output",
		Value:      "json",
		Persistent: true,
		Category:   "Global:",
	}

	outputPrettyFlag = &cli.BoolFlag{
		Name:       "pretty",
		Usage:      "enable pretty output, if relevant for the output type",
		Value:      true,
		Category:   "Global:",
		Persistent: true,
	}

	appIDFlag = &cli.StringFlag{
		Name:    "app-id",
		Aliases: []string{"a"},
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(buildFlagEnvVar("APP_ID")),
			altsrc.File(configFlag.Name, "app.id"),
		),
		Usage:      "the ID of the app in Hathora",
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
			altsrc.File(configFlag.Name, "global.verbosity"),
		),
		Usage:      "set the logging verbosity level",
		Value:      0,
		Category:   "Global:",
		Persistent: true,
	}

	hathoraCloudEndpointFlag = &cli.StringFlag{
		Name: "hathora-cloud-endpoint",
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(globalFlagEnvVar("CLOUD_ENDPOINT")),
			altsrc.File(configFlag.Name, "global.cloud-endpoint"),
		),
		Usage:       "override the default API base url",
		DefaultText: "https://api.hathora.dev",
		Category:    "Global:",
		Persistent:  true,
	}

	tokenFlag = &cli.StringFlag{
		Name:       "token",
		Aliases:    []string{"t"},
		Sources:    cli.EnvVars(globalFlagEnvVar("TOKEN")),
		Usage:      "the access token for authenticating with the API",
		Category:   "Global:",
		Persistent: true,
	}

	configFlag = &cli.StringFlag{
		Name:       "config",
		Aliases:    []string{"c"},
		Usage:      "path to the configuration file",
		Category:   "Global:",
		Persistent: true,
	}

	GlobalFlags = []cli.Flag{
		appIDFlag,
		hathoraCloudEndpointFlag,
		tokenFlag,
		outputTypeFlag,
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
	return flags
}
