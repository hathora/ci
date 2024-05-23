package commands

import (
	"context"

	"github.com/urfave/cli/v3"
)

var (
	outputTypeFlag = &cli.StringFlag{
		Name:       "output",
		Aliases:    []string{"o"},
		Sources:    CombineSources(EnvVarSource("OUTPUT"), ConfigFileSource("output")),
		Usage:      "the format of the output",
		Value:      allowedOutputTypes[0],
		Persistent: true,
		Category:   "Global:",
		Action: func(ctx context.Context, cmd *cli.Command, v string) error {
			return requireValidEnumValue(v, allowedOutputTypes, "output")
		},
	}

	outputPrettyFlag = &cli.BoolFlag{
		Name:       "pretty",
		Usage:      "enable pretty output, if relevant for the output type",
		Value:      true,
		Category:   "Global:",
		Persistent: true,
	}

	appIDFlag = &cli.StringFlag{
		Name:       "app-id",
		Aliases:    []string{"a"},
		Sources:    CombineSources(EnvVarSource("APP_ID"), ConfigFileSource("app-id")),
		Usage:      "the ID of the app in Hathora",
		Category:   "Global:",
		Required:   true,
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
		Name:       "verbosity",
		Sources:    CombineSources(EnvVarSource("VERBOSITY"), ConfigFileSource("verbosity")),
		Usage:      "set the logging verbosity level",
		Value:      0,
		Category:   "Global:",
		Persistent: true,
	}

	hathoraCloudEndpointFlag = &cli.StringFlag{
		Name:        "hathora-cloud-endpoint",
		Sources:     CombineSources(EnvVarSource("CLOUD_ENDPOINT"), ConfigFileSource("hathora-cloud-endpoint")),
		Usage:       "override the default API base url",
		DefaultText: "https://api.hathora.dev",
		Category:    "Global:",
		Persistent:  true,
	}

	tokenFlag = &cli.StringFlag{
		Name:       "token",
		Aliases:    []string{"t"},
		Sources:    CombineSources(EnvVarSource("TOKEN"), ConfigFileSource("token")),
		Usage:      "the access token for authenticating with the API",
		Category:   "Global:",
		Required:   true,
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
	}

	allowedOutputTypes = []string{"text", "json"}
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
