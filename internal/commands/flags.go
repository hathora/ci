package commands

import (
	"github.com/hathora/ci/internal/sdk"
	"github.com/urfave/cli/v2"
)

var (
	buildIDFlag = &cli.IntFlag{
		Name:     "build-id",
		Aliases:  []string{"b"},
		EnvVars:  []string{"HATHORA_BUILD_ID"},
		Usage:    "the ID of the build in Hathora",
		Required: true,
	}

	buildTagFlag = &cli.StringFlag{
		Name:     "build-tag",
		Aliases:  []string{"bt"},
		EnvVars:  []string{"HATHORA_APP_TAG"},
		Usage:    "tag to associate an external version with a build",
		Required: true,
	}

	deploymentIDFlag = &cli.IntFlag{
		Name:     "deployment-id",
		Aliases:  []string{"d"},
		EnvVars:  []string{"HATHORA_DEPLOYMENT_ID"},
		Usage:    "the ID of the deployment in Hathora",
		Required: true,
	}

	// common flags
	outputTypeFlag = &cli.StringFlag{
		Name:     "output-type",
		Aliases:  []string{"o"},
		Usage:    "the format of the output",
		Value:    allowedOutputTypes[0],
		Category: "Global:",
		Action: func(ctx *cli.Context, v string) error {
			return requireValidEnumValue(v, allowedOutputTypes, "output-type")
		},
	}

	appIDFlag = &cli.StringFlag{
		Name:     "app-id",
		Aliases:  []string{"a"},
		EnvVars:  []string{"HATHORA_APP_ID"},
		Usage:    "the ID of the app in Hathora",
		Required: true,
		Category: "Global:",
	}

	hathoraCloudEndpointFlag = &cli.StringFlag{
		Name:        "hathora-cloud-endpoint",
		EnvVars:     []string{"HATHORA_CLOUD_ENDPOINT"},
		Usage:       "override the default API base url",
		DefaultText: "https://api.hathora.dev",
		Category:    "Global:",
	}

	tokenFlag = &cli.StringFlag{
		Name:     "token",
		Aliases:  []string{"t"},
		EnvVars:  []string{"HATHORA_TOKEN"},
		Usage:    "the access token for authenticating with the API",
		Category: "Global:",
		Required: true,
	}

	globalFlags = []cli.Flag{
		appIDFlag,
		hathoraCloudEndpointFlag,
		tokenFlag,
		outputTypeFlag,
	}

	allowedOutputTypes = []string{"text", "json"}
)

func getCommonFlagValues(context *cli.Context) (token string, baseUrl string, appID *string, outputType string) {
	token = context.String(tokenFlag.Name)
	baseUrl = context.String(hathoraCloudEndpointFlag.Name)
	appID = sdk.String(context.String(appIDFlag.Name))
	outputType = context.String(outputTypeFlag.Name)

	return token, baseUrl, appID, outputType
}