package commands

import (
	"github.com/urfave/cli/v2"
)

var (
	// common flags
	outputTypeFlag = &cli.StringFlag{
		Name:     "output",
		Aliases:  []string{"o"},
		EnvVars:  globalFlagEnvVar("OUTPUT"),
		Usage:    "the format of the output",
		Value:    allowedOutputTypes[0],
		Category: "Global:",
		Action: func(ctx *cli.Context, v string) error {
			return requireValidEnumValue(v, allowedOutputTypes, "output")
		},
	}

	appIDFlag = &cli.StringFlag{
		Name:     "app-id",
		Aliases:  []string{"a"},
		EnvVars:  globalFlagEnvVar("APP_ID"),
		Usage:    "the ID of the app in Hathora",
		Required: true,
		Category: "Global:",
	}

	verboseFlag = &cli.BoolFlag{
		Name:     "verbose",
		Aliases:  []string{"v"},
		Usage:    "enable verbose logging",
		Category: "Global:",
	}

	verbosityFlag = &cli.IntFlag{
		Name:     "verbosity",
		EnvVars:  globalFlagEnvVar("VERBOSITY"),
		Usage:    "set the logging verbosity level",
		Value:    0,
		Category: "Global:",
	}

	hathoraCloudEndpointFlag = &cli.StringFlag{
		Name:        "hathora-cloud-endpoint",
		EnvVars:     globalFlagEnvVar("CLOUD_ENDPOINT"),
		Usage:       "override the default API base url",
		DefaultText: "https://api.hathora.dev",
		Category:    "Global:",
	}

	tokenFlag = &cli.StringFlag{
		Name:     "token",
		Aliases:  []string{"t"},
		EnvVars:  globalFlagEnvVar("TOKEN"),
		Usage:    "the access token for authenticating with the API",
		Category: "Global:",
		Required: true,
	}

	GlobalFlags = []cli.Flag{
		appIDFlag,
		hathoraCloudEndpointFlag,
		tokenFlag,
		outputTypeFlag,
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
	return append(GlobalFlags, flags...)
}
