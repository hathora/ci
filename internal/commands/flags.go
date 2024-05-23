package commands

import (
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

var (
	// common flags
	outputTypeFlag = altsrc.NewStringFlag(&cli.StringFlag{
		Name:     "output",
		Aliases:  []string{"o"},
		EnvVars:  globalFlagEnvVar("OUTPUT"),
		Usage:    "the format of the output",
		Value:    allowedOutputTypes[0],
		Category: "Global:",
		Action: func(ctx *cli.Context, v string) error {
			return requireValidEnumValue(v, allowedOutputTypes, "output")
		},
	})

	outputPrettyFlag = altsrc.NewBoolFlag(&cli.BoolFlag{
		Name:     "pretty",
		Usage:    "enable pretty output, if relevant for the output type",
		Value:    true,
		Category: "Global:",
	})

	appIDFlag = altsrc.NewStringFlag(&cli.StringFlag{
		Name:     "app-id",
		Aliases:  []string{"a"},
		EnvVars:  globalFlagEnvVar("APP_ID"),
		Usage:    "the ID of the app in Hathora",
		Category: "Global:",
	})

	verboseFlag = altsrc.NewBoolFlag(&cli.BoolFlag{
		Name:     "verbose",
		Aliases:  []string{"v"},
		Usage:    "enable verbose logging",
		Category: "Global:",
	})

	verbosityFlag = altsrc.NewIntFlag(&cli.IntFlag{
		Name:     "verbosity",
		EnvVars:  globalFlagEnvVar("VERBOSITY"),
		Usage:    "set the logging verbosity level",
		Value:    0,
		Category: "Global:",
	})

	hathoraCloudEndpointFlag = altsrc.NewStringFlag(&cli.StringFlag{
		Name:        "hathora-cloud-endpoint",
		EnvVars:     globalFlagEnvVar("CLOUD_ENDPOINT"),
		Usage:       "override the default API base url",
		DefaultText: "https://api.hathora.dev",
		Category:    "Global:",
	})

	tokenFlag = altsrc.NewStringFlag(&cli.StringFlag{
		Name:     "token",
		Aliases:  []string{"t"},
		EnvVars:  globalFlagEnvVar("TOKEN"),
		Usage:    "the access token for authenticating with the API",
		Category: "Global:",
	})

	configFileFlag = &cli.StringFlag{
		Name:     "config",
		Aliases:  []string{"c"},
		EnvVars:  globalFlagEnvVar("CONFIG"),
		Usage:    "path to config file with flag values",
		Category: "Global:",
	}

	GlobalFlags = []cli.Flag{
		appIDFlag,
		configFileFlag,
		hathoraCloudEndpointFlag,
		tokenFlag,
		outputTypeFlag,
		outputPrettyFlag,
		verboseFlag,
		verbosityFlag,
	}

	allowedOutputTypes = []string{"text", "json"}
)

var (
	globalFlagEnvVarPrefix = "HATHORA_"
)

func globalFlagEnvVar(name string) []string {
	return []string{globalFlagEnvVarPrefix + name}
}

func subcommandFlags(flags ...cli.Flag) []cli.Flag {
	return flags
}
