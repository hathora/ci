package commands

import (
	"github.com/urfave/cli/v3"

	"github.com/hathora/ci/internal/commands/altsrc"
)

const (
	appIDFlagName                = "app-id"
	configFlagName               = "config"
	outputFlagName               = "output"
	outputPrettyFlagName         = "pretty"
	hathoraCloudEndpointFlagName = "hathora-cloud-endpoint"
	tokenFlagName                = "token"
	verboseFlagName              = "verbose"
	verbosityFlagName            = "verbosity"
)

func outputFlag() *cli.StringFlag {
	return &cli.StringFlag{
		Name:    outputFlagName,
		Aliases: []string{"o"},
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(globalFlagEnvVar("OUTPUT")),
			altsrc.ConfigFile(configFlagName, "global.output"),
		),
		Usage:       "the `<format>` of the output. Supported values: (json, text, buildIdValue)",
		Value:       "text",
		DefaultText: "text",
		Category:    "Global:",
	}
}
func outputPrettyFlag() *cli.BoolFlag {
	return &cli.BoolFlag{
		Name:     outputPrettyFlagName,
		Usage:    "enable pretty output (json only)",
		Value:    true,
		Category: "Global:",
	}
}

func appIDFlag() *cli.StringFlag {
	return &cli.StringFlag{
		Name:    appIDFlagName,
		Aliases: []string{"a"},
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(globalFlagEnvVar("APP_ID")),
			altsrc.ConfigFile(configFlagName, "app.id"),
		),
		Usage:    "the `<id>` of the app in Hathora",
		Category: "Global:",
	}
}

func verboseFlag() *cli.BoolFlag {
	return &cli.BoolFlag{
		Name:     verboseFlagName,
		Aliases:  []string{"v"},
		Usage:    "enable verbose logging",
		Category: "Global:",
	}

}

func verbosityFlag() *cli.IntFlag {
	return &cli.IntFlag{
		Name: verbosityFlagName,
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(globalFlagEnvVar("VERBOSITY")),
			altsrc.ConfigFile(configFlagName, "global.verbosity"),
		),
		Usage:    "set the logging verbosity `<level>` (0-3)",
		Value:    0,
		Category: "Global:",
	}
}

func hathoraCloudEndpointFlag() *cli.StringFlag {
	return &cli.StringFlag{
		Name: hathoraCloudEndpointFlagName,
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(globalFlagEnvVar("CLOUD_ENDPOINT")),
			altsrc.ConfigFile(configFlagName, "global.cloud-endpoint"),
		),
		Usage:       "override the default API base `<url>`",
		Value:       "https://api.hathora.dev",
		DefaultText: "https://api.hathora.dev",
		Category:    "Global:",
	}
}

func tokenFlag() *cli.StringFlag {
	return &cli.StringFlag{
		Name:     tokenFlagName,
		Aliases:  []string{"t"},
		Sources:  cli.EnvVars(globalFlagEnvVar("TOKEN")),
		Usage:    "`<access-token>` for authenticating with the API",
		Category: "Global:",
	}
}

func configFlag() *cli.StringFlag {
	return &cli.StringFlag{
		Name:     configFlagName,
		Aliases:  []string{"c"},
		Usage:    "`<path>` to the configuration file",
		Category: "Global:",
	}
}

func GlobalFlags() []cli.Flag {
	return []cli.Flag{
		appIDFlag(),
		hathoraCloudEndpointFlag(),
		tokenFlag(),
		outputFlag(),
		outputPrettyFlag(),
		verboseFlag(),
		verbosityFlag(),
		configFlag(),
	}
}

var (
	globalFlagEnvVarPrefix = "HATHORA_"
)

func globalFlagEnvVar(name string) string {
	return globalFlagEnvVarPrefix + name
}

func subcommandFlags(flags ...cli.Flag) []cli.Flag {
	return append(flags, GlobalFlags()...)
}
