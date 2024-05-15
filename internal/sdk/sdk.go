// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"context"
	"fmt"
	"github.com/hathora/ci/internal/sdk/internal/globals"
	"github.com/hathora/ci/internal/sdk/internal/hooks"
	"github.com/hathora/ci/internal/sdk/internal/utils"
	"github.com/hathora/ci/internal/sdk/models/shared"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	"https://api.hathora.dev",
	"https:///",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	Client            HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	Globals           globals.Globals
	RetryConfig       *utils.RetryConfig
	Hooks             *hooks.Hooks
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// SDK - Hathora Cloud API: Welcome to the Hathora Cloud API documentation! Learn how to use the Hathora Cloud APIs to build and scale your game servers globally.
type SDK struct {
	// Operations that allow you manage your [applications](https://hathora.dev/docs/concepts/hathora-entities#application).
	AppV1 *AppV1
	// Operations that allow you to generate a Hathora-signed [JSON web token (JWT)](https://jwt.io/) for [player authentication](https://hathora.dev/docs/lobbies-and-matchmaking/auth-service).
	AuthV1 *AuthV1
	//
	BillingV1 *BillingV1
	// Deprecated. Use [BuildV2](https://hathora.dev/api#tag/BuildV2).
	BuildV1 *BuildV1
	// Operations that allow you create and manage your [builds](https://hathora.dev/docs/concepts/hathora-entities#build).
	BuildV2 *BuildV2
	// Deprecated. Use [DeploymentV2](https://hathora.dev/api#tag/DeploymentV2).
	DeploymentV1 *DeploymentV1
	// Operations that allow you configure and manage an application's [build](https://hathora.dev/docs/concepts/hathora-entities#build) at runtime.
	DeploymentV2 *DeploymentV2
	// Deprecated. Does not include latest Regions (missing Dallas region). Use [DiscoveryV2](https://hathora.dev/api#tag/DiscoveryV2).
	DiscoveryV1 *DiscoveryV1
	// Service that allows clients to directly ping all Hathora regions to get latency information
	DiscoveryV2 *DiscoveryV2
	// Deprecated. Use [LobbyV3](https://hathora.dev/api#tag/LobbyV3).
	LobbyV1 *LobbyV1
	// Deprecated. Use [LobbyV3](https://hathora.dev/api#tag/LobbyV3).
	LobbyV2 *LobbyV2
	// Operations to create and manage lobbies using our [Lobby Service](https://hathora.dev/docs/lobbies-and-matchmaking/lobby-service).
	LobbyV3 *LobbyV3
	// Operations to get logs by [applications](https://hathora.dev/docs/concepts/hathora-entities#application), [processes](https://hathora.dev/docs/concepts/hathora-entities#process), and [deployments](https://hathora.dev/docs/concepts/hathora-entities#deployment). We store 20GB of logs data.
	LogV1 *LogV1
	//
	ManagementV1 *ManagementV1
	// Operations to get metrics by [process](https://hathora.dev/docs/concepts/hathora-entities#process). We store 72 hours of metrics data.
	MetricsV1       *MetricsV1
	OrganizationsV1 *OrganizationsV1
	// Deprecated. Use [ProcessesV2](https://hathora.dev/api#tag/ProcessesV2).
	ProcessesV1 *ProcessesV1
	// Operations to get data on active and stopped [processes](https://hathora.dev/docs/concepts/hathora-entities#process).
	ProcessesV2 *ProcessesV2
	// Deprecated. Use [RoomV2](https://hathora.dev/api#tag/RoomV2).
	RoomV1 *RoomV1
	// Operations to create, manage, and connect to [rooms](https://hathora.dev/docs/concepts/hathora-entities#room).
	RoomV2 *RoomV2
	//
	OrgTokensV1 *OrgTokensV1

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*SDK)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *SDK) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *SDK) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.Client = client
	}
}

func withSecurity(security interface{}) func(context.Context) (interface{}, error) {
	return func(context.Context) (interface{}, error) {
		return security, nil
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.Security = withSecurity(security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (shared.Security, error)) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

// WithAppID allows setting the AppID parameter for all supported operations
func WithAppID(appID string) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.Globals.AppID = &appID
	}
}

func WithRetryConfig(retryConfig utils.RetryConfig) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *SDK {
	sdk := &SDK{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "0.0.1",
			SDKVersion:        "0.2.1",
			GenVersion:        "2.322.5",
			UserAgent:         "speakeasy-sdk/go 0.2.1 2.322.5 0.0.1 github.com/hathora/ci/internal/sdk",
			Globals:           globals.Globals{},
			Hooks:             hooks.New(),
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.Client == nil {
		sdk.sdkConfiguration.Client = &http.Client{Timeout: 60 * time.Second}
	}

	currentServerURL, _ := sdk.sdkConfiguration.GetServerDetails()
	serverURL := currentServerURL
	serverURL, sdk.sdkConfiguration.Client = sdk.sdkConfiguration.Hooks.SDKInit(currentServerURL, sdk.sdkConfiguration.Client)
	if serverURL != currentServerURL {
		sdk.sdkConfiguration.ServerURL = serverURL
	}

	sdk.AppV1 = newAppV1(sdk.sdkConfiguration)

	sdk.AuthV1 = newAuthV1(sdk.sdkConfiguration)

	sdk.BillingV1 = newBillingV1(sdk.sdkConfiguration)

	sdk.BuildV1 = newBuildV1(sdk.sdkConfiguration)

	sdk.BuildV2 = newBuildV2(sdk.sdkConfiguration)

	sdk.DeploymentV1 = newDeploymentV1(sdk.sdkConfiguration)

	sdk.DeploymentV2 = newDeploymentV2(sdk.sdkConfiguration)

	sdk.DiscoveryV1 = newDiscoveryV1(sdk.sdkConfiguration)

	sdk.DiscoveryV2 = newDiscoveryV2(sdk.sdkConfiguration)

	sdk.LobbyV1 = newLobbyV1(sdk.sdkConfiguration)

	sdk.LobbyV2 = newLobbyV2(sdk.sdkConfiguration)

	sdk.LobbyV3 = newLobbyV3(sdk.sdkConfiguration)

	sdk.LogV1 = newLogV1(sdk.sdkConfiguration)

	sdk.ManagementV1 = newManagementV1(sdk.sdkConfiguration)

	sdk.MetricsV1 = newMetricsV1(sdk.sdkConfiguration)

	sdk.OrganizationsV1 = newOrganizationsV1(sdk.sdkConfiguration)

	sdk.ProcessesV1 = newProcessesV1(sdk.sdkConfiguration)

	sdk.ProcessesV2 = newProcessesV2(sdk.sdkConfiguration)

	sdk.RoomV1 = newRoomV1(sdk.sdkConfiguration)

	sdk.RoomV2 = newRoomV2(sdk.sdkConfiguration)

	sdk.OrgTokensV1 = newOrgTokensV1(sdk.sdkConfiguration)

	return sdk
}
