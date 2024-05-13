# github.com/hathora/ci/internal/sdk

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>


## 🏗 **Welcome to your new SDK!** 🏗

It has been generated successfully based on your OpenAPI spec. However, it is not yet ready for production use. Here are some next steps:
- [ ] 🛠 Make your SDK feel handcrafted by [customizing it](https://www.speakeasyapi.dev/docs/customize-sdks)
- [ ] ♻️ Refine your SDK quickly by iterating locally with the [Speakeasy CLI](https://github.com/speakeasy-api/speakeasy)
- [ ] 🎁 Publish your SDK to package managers by [configuring automatic publishing](https://www.speakeasyapi.dev/docs/advanced-setup/publish-sdks)
- [ ] ✨ When ready to productionize, delete this section from the README

<!-- Start SDK Installation [installation] -->
## SDK Installation

```bash
go get github.com/hathora/ci/internal/sdk
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	"github.com/hathora/ci/internal/sdk"
	"github.com/hathora/ci/internal/sdk/pkg/models/shared"
	"log"
)

func main() {
	s := sdk.New(
		sdk.WithSecurity(shared.Security{
			HathoraDevToken: sdk.String("<YOUR_BEARER_TOKEN_HERE>"),
		}),
		sdk.WithAppID(sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")),
	)

	ctx := context.Background()
	res, err := s.AppV1.GetApps(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.ApplicationWithLatestDeploymentAndBuilds != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

### [AppV1](docs/sdks/appv1/README.md)

* [GetApps](docs/sdks/appv1/README.md#getapps) - Returns an unsorted list of your organization’s [applications](https://hathora.dev/docs/concepts/hathora-entities#application). An application is uniquely identified by an `appId`.
* [CreateApp](docs/sdks/appv1/README.md#createapp) - Create a new [application](https://hathora.dev/docs/concepts/hathora-entities#application).
* [UpdateApp](docs/sdks/appv1/README.md#updateapp) - Update data for an existing [application](https://hathora.dev/docs/concepts/hathora-entities#application) using `appId`.
* [GetAppInfo](docs/sdks/appv1/README.md#getappinfo) - Get details for an [application](https://hathora.dev/docs/concepts/hathora-entities#application) using `appId`.
* [DeleteApp](docs/sdks/appv1/README.md#deleteapp) - Delete an [application](https://hathora.dev/docs/concepts/hathora-entities#application) using `appId`. Your organization will lose access to this application.

### [AuthV1](docs/sdks/authv1/README.md)

* [LoginAnonymous](docs/sdks/authv1/README.md#loginanonymous) - Returns a unique player token for an anonymous user.
* [LoginNickname](docs/sdks/authv1/README.md#loginnickname) - Returns a unique player token with a specified nickname for a user.
* [LoginGoogle](docs/sdks/authv1/README.md#logingoogle) - Returns a unique player token using a Google-signed OIDC `idToken`.

### [BillingV1](docs/sdks/billingv1/README.md)

* [GetBalance](docs/sdks/billingv1/README.md#getbalance)
* [GetPaymentMethod](docs/sdks/billingv1/README.md#getpaymentmethod)
* [InitStripeCustomerPortalURL](docs/sdks/billingv1/README.md#initstripecustomerportalurl)
* [GetInvoices](docs/sdks/billingv1/README.md#getinvoices)

### [BuildV1](docs/sdks/buildv1/README.md)

* [~~GetBuildsDeprecated~~](docs/sdks/buildv1/README.md#getbuildsdeprecated) - Returns an array of [builds](https://hathora.dev/docs/concepts/hathora-entities#build) for an [application](https://hathora.dev/docs/concepts/hathora-entities#application). :warning: **Deprecated**
* [~~GetBuildInfoDeprecated~~](docs/sdks/buildv1/README.md#getbuildinfodeprecated) - Get details for a [build](https://hathora.dev/docs/concepts/hathora-entities#build). :warning: **Deprecated**
* [~~CreateBuildDeprecated~~](docs/sdks/buildv1/README.md#createbuilddeprecated) - Creates a new [build](https://hathora.dev/docs/concepts/hathora-entities#build). Responds with a `buildId` that you must pass to [`RunBuild()`](https://hathora.dev/api#tag/BuildV1/operation/RunBuild) to build the game server artifact. You can optionally pass in a `buildTag` to associate an external version with a build. :warning: **Deprecated**
* [~~RunBuildDeprecated~~](docs/sdks/buildv1/README.md#runbuilddeprecated) - Builds a game server artifact from a tarball you provide. Pass in the `buildId` generated from [`CreateBuild()`](https://hathora.dev/api#tag/BuildV1/operation/CreateBuild). :warning: **Deprecated**
* [~~DeleteBuildDeprecated~~](docs/sdks/buildv1/README.md#deletebuilddeprecated) - Delete a [build](https://hathora.dev/docs/concepts/hathora-entities#build). All associated metadata is deleted. :warning: **Deprecated**

### [BuildV2](docs/sdks/buildv2/README.md)

* [GetBuilds](docs/sdks/buildv2/README.md#getbuilds) - Returns an array of [builds](https://hathora.dev/docs/concepts/hathora-entities#build) for an [application](https://hathora.dev/docs/concepts/hathora-entities#application).
* [GetBuildInfo](docs/sdks/buildv2/README.md#getbuildinfo) - Get details for a [build](https://hathora.dev/docs/concepts/hathora-entities#build).
* [CreateBuild](docs/sdks/buildv2/README.md#createbuild) - Creates a new [build](https://hathora.dev/docs/concepts/hathora-entities#build). Responds with a `buildId` that you must pass to [`RunBuild()`](https://hathora.dev/api#tag/BuildV1/operation/RunBuild) to build the game server artifact. You can optionally pass in a `buildTag` to associate an external version with a build.
* [RunBuild](docs/sdks/buildv2/README.md#runbuild) - Builds a game server artifact from a tarball you provide. Pass in the `buildId` generated from [`CreateBuild()`](https://hathora.dev/api#tag/BuildV1/operation/CreateBuild).
* [DeleteBuild](docs/sdks/buildv2/README.md#deletebuild) - Delete a [build](https://hathora.dev/docs/concepts/hathora-entities#build). All associated metadata is deleted.

### [DeploymentV1](docs/sdks/deploymentv1/README.md)

* [~~GetDeploymentsDeprecated~~](docs/sdks/deploymentv1/README.md#getdeploymentsdeprecated) - Returns an array of [deployments](https://hathora.dev/docs/concepts/hathora-entities#deployment) for an [application](https://hathora.dev/docs/concepts/hathora-entities#application). :warning: **Deprecated**
* [~~GetLatestDeploymentDeprecated~~](docs/sdks/deploymentv1/README.md#getlatestdeploymentdeprecated) - Get the latest [deployment](https://hathora.dev/docs/concepts/hathora-entities#deployment) for an [application](https://hathora.dev/docs/concepts/hathora-entities#application). :warning: **Deprecated**
* [~~GetDeploymentInfoDeprecated~~](docs/sdks/deploymentv1/README.md#getdeploymentinfodeprecated) - Get details for a [deployment](https://hathora.dev/docs/concepts/hathora-entities#deployment). :warning: **Deprecated**
* [~~CreateDeploymentDeprecated~~](docs/sdks/deploymentv1/README.md#createdeploymentdeprecated) - Create a new [deployment](https://hathora.dev/docs/concepts/hathora-entities#deployment). Creating a new deployment means all new rooms created will use the latest deployment configuration, but existing games in progress will not be affected. :warning: **Deprecated**

### [DeploymentV2](docs/sdks/deploymentv2/README.md)

* [GetDeployments](docs/sdks/deploymentv2/README.md#getdeployments) - Returns an array of [deployments](https://hathora.dev/docs/concepts/hathora-entities#deployment) for an [application](https://hathora.dev/docs/concepts/hathora-entities#application).
* [GetLatestDeployment](docs/sdks/deploymentv2/README.md#getlatestdeployment) - Get the latest [deployment](https://hathora.dev/docs/concepts/hathora-entities#deployment) for an [application](https://hathora.dev/docs/concepts/hathora-entities#application).
* [GetDeploymentInfo](docs/sdks/deploymentv2/README.md#getdeploymentinfo) - Get details for a [deployment](https://hathora.dev/docs/concepts/hathora-entities#deployment).
* [CreateDeployment](docs/sdks/deploymentv2/README.md#createdeployment) - Create a new [deployment](https://hathora.dev/docs/concepts/hathora-entities#deployment). Creating a new deployment means all new rooms created will use the latest deployment configuration, but existing games in progress will not be affected.

### [DiscoveryV1](docs/sdks/discoveryv1/README.md)

* [GetPingServiceEndpoints](docs/sdks/discoveryv1/README.md#getpingserviceendpoints) - Returns an array of all regions with a host and port that a client can directly ping. Open a websocket connection to `wss://<host>:<port>/ws` and send a packet. To calculate ping, measure the time it takes to get an echo packet back.

### [LobbyV1](docs/sdks/lobbyv1/README.md)

* [~~CreatePrivateLobbyDeprecated~~](docs/sdks/lobbyv1/README.md#createprivatelobbydeprecated) - :warning: **Deprecated**
* [~~CreatePublicLobbyDeprecated~~](docs/sdks/lobbyv1/README.md#createpubliclobbydeprecated) - :warning: **Deprecated**
* [~~ListActivePublicLobbiesDeprecatedV1~~](docs/sdks/lobbyv1/README.md#listactivepubliclobbiesdeprecatedv1) - :warning: **Deprecated**

### [LobbyV2](docs/sdks/lobbyv2/README.md)

* [~~CreatePrivateLobby~~](docs/sdks/lobbyv2/README.md#createprivatelobby) - :warning: **Deprecated**
* [~~CreatePublicLobby~~](docs/sdks/lobbyv2/README.md#createpubliclobby) - :warning: **Deprecated**
* [~~CreateLocalLobby~~](docs/sdks/lobbyv2/README.md#createlocallobby) - :warning: **Deprecated**
* [~~CreateLobbyDeprecated~~](docs/sdks/lobbyv2/README.md#createlobbydeprecated) - Create a new lobby for an [application](https://hathora.dev/docs/concepts/hathora-entities#application). A lobby object is a wrapper around a [room](https://hathora.dev/docs/concepts/hathora-entities#room) object. With a lobby, you get additional functionality like configuring the visibility of the room, managing the state of a match, and retrieving a list of public lobbies to display to players. :warning: **Deprecated**
* [~~ListActivePublicLobbiesDeprecatedV2~~](docs/sdks/lobbyv2/README.md#listactivepubliclobbiesdeprecatedv2) - Get all active lobbies for a an [application](https://hathora.dev/docs/concepts/hathora-entities#application). Filter by optionally passing in a `region`. Use this endpoint to display all public lobbies that a player can join in the game client. :warning: **Deprecated**
* [~~GetLobbyInfo~~](docs/sdks/lobbyv2/README.md#getlobbyinfo) - Get details for a lobby. :warning: **Deprecated**
* [~~SetLobbyState~~](docs/sdks/lobbyv2/README.md#setlobbystate) - Set the state of a lobby. State is intended to be set by the server and must be smaller than 1MB. Use this endpoint to store match data like live player count to enforce max number of clients or persist end-game data (i.e. winner or final scores). :warning: **Deprecated**

### [LobbyV3](docs/sdks/lobbyv3/README.md)

* [CreateLobby](docs/sdks/lobbyv3/README.md#createlobby) - Create a new lobby for an [application](https://hathora.dev/docs/concepts/hathora-entities#application). A lobby object is a wrapper around a [room](https://hathora.dev/docs/concepts/hathora-entities#room) object. With a lobby, you get additional functionality like configuring the visibility of the room, managing the state of a match, and retrieving a list of public lobbies to display to players.
* [ListActivePublicLobbies](docs/sdks/lobbyv3/README.md#listactivepubliclobbies) - Get all active lobbies for a given [application](https://hathora.dev/docs/concepts/hathora-entities#application). Filter the array by optionally passing in a `region`. Use this endpoint to display all public lobbies that a player can join in the game client.
* [GetLobbyInfoByRoomID](docs/sdks/lobbyv3/README.md#getlobbyinfobyroomid) - Get details for a lobby.
* [GetLobbyInfoByShortCode](docs/sdks/lobbyv3/README.md#getlobbyinfobyshortcode) - Get details for a lobby. If 2 or more lobbies have the same `shortCode`, then the most recently created lobby will be returned.

### [LogV1](docs/sdks/logv1/README.md)

* [~~GetLogsForApp~~](docs/sdks/logv1/README.md#getlogsforapp) - Returns a stream of logs for an [application](https://hathora.dev/docs/concepts/hathora-entities#application) using `appId`. :warning: **Deprecated**
* [GetLogsForProcess](docs/sdks/logv1/README.md#getlogsforprocess) - Returns a stream of logs for a [process](https://hathora.dev/docs/concepts/hathora-entities#process) using `appId` and `processId`.
* [DownloadLogForProcess](docs/sdks/logv1/README.md#downloadlogforprocess) - Download entire log file for a stopped process.
* [~~GetLogsForDeployment~~](docs/sdks/logv1/README.md#getlogsfordeployment) - Returns a stream of logs for a [deployment](https://hathora.dev/docs/concepts/hathora-entities#deployment) using `appId` and `deploymentId`. :warning: **Deprecated**

### [ManagementV1](docs/sdks/managementv1/README.md)

* [SendVerificationEmail](docs/sdks/managementv1/README.md#sendverificationemail)

### [MetricsV1](docs/sdks/metricsv1/README.md)

* [GetMetrics](docs/sdks/metricsv1/README.md#getmetrics) - Get metrics for a [process](https://hathora.dev/docs/concepts/hathora-entities#process) using `appId` and `processId`.

### [OrganizationsV1](docs/sdks/organizationsv1/README.md)

* [InviteUser](docs/sdks/organizationsv1/README.md#inviteuser)
* [RescindInvite](docs/sdks/organizationsv1/README.md#rescindinvite)
* [GetOrgPendingInvites](docs/sdks/organizationsv1/README.md#getorgpendinginvites)
* [GetUserPendingInvites](docs/sdks/organizationsv1/README.md#getuserpendinginvites)
* [AcceptInvite](docs/sdks/organizationsv1/README.md#acceptinvite)
* [RejectInvite](docs/sdks/organizationsv1/README.md#rejectinvite)

### [ProcessesV1](docs/sdks/processesv1/README.md)

* [~~GetRunningProcesses~~](docs/sdks/processesv1/README.md#getrunningprocesses) - Retrieve 10 most recently started [process](https://hathora.dev/docs/concepts/hathora-entities#process) objects for an [application](https://hathora.dev/docs/concepts/hathora-entities#application). Filter the array by optionally passing in a `region`. :warning: **Deprecated**
* [~~GetStoppedProcesses~~](docs/sdks/processesv1/README.md#getstoppedprocesses) - Retrieve 10 most recently stopped [process](https://hathora.dev/docs/concepts/hathora-entities#process) objects for an [application](https://hathora.dev/docs/concepts/hathora-entities#application). Filter the array by optionally passing in a `region`. :warning: **Deprecated**
* [~~GetProcessInfoDeprecated~~](docs/sdks/processesv1/README.md#getprocessinfodeprecated) - Get details for a [process](https://hathora.dev/docs/concepts/hathora-entities#process). :warning: **Deprecated**

### [ProcessesV2](docs/sdks/processesv2/README.md)

* [GetProcessInfo](docs/sdks/processesv2/README.md#getprocessinfo) - Get details for a [process](https://hathora.dev/docs/concepts/hathora-entities#process).
* [GetLatestProcesses](docs/sdks/processesv2/README.md#getlatestprocesses) - Retrieve the 10 most recent [processes](https://hathora.dev/docs/concepts/hathora-entities#process) objects for an [application](https://hathora.dev/docs/concepts/hathora-entities#application). Filter the array by optionally passing in a `status` or `region`.
* [StopProcess](docs/sdks/processesv2/README.md#stopprocess) - Stops a [process](https://hathora.dev/docs/concepts/hathora-entities#process) immediately.
* [CreateProcess](docs/sdks/processesv2/README.md#createprocess) - Creates a [process](https://hathora.dev/docs/concepts/hathora-entities#process) without a room. Use this to pre-allocate processes ahead of time so that subsequent room assignment via [CreateRoom()](https://hathora.dev/api#tag/RoomV2/operation/CreateRoom) can be instant.

### [RoomV1](docs/sdks/roomv1/README.md)

* [~~CreateRoomDeprecated~~](docs/sdks/roomv1/README.md#createroomdeprecated) - :warning: **Deprecated**
* [~~GetRoomInfoDeprecated~~](docs/sdks/roomv1/README.md#getroominfodeprecated) - :warning: **Deprecated**
* [~~GetActiveRoomsForProcessDeprecated~~](docs/sdks/roomv1/README.md#getactiveroomsforprocessdeprecated) - :warning: **Deprecated**
* [~~GetInactiveRoomsForProcessDeprecated~~](docs/sdks/roomv1/README.md#getinactiveroomsforprocessdeprecated) - :warning: **Deprecated**
* [~~DestroyRoomDeprecated~~](docs/sdks/roomv1/README.md#destroyroomdeprecated) - :warning: **Deprecated**
* [~~SuspendRoomDeprecated~~](docs/sdks/roomv1/README.md#suspendroomdeprecated) - :warning: **Deprecated**
* [~~GetConnectionInfoDeprecated~~](docs/sdks/roomv1/README.md#getconnectioninfodeprecated) - :warning: **Deprecated**

### [RoomV2](docs/sdks/roomv2/README.md)

* [CreateRoom](docs/sdks/roomv2/README.md#createroom) - Create a new [room](https://hathora.dev/docs/concepts/hathora-entities#room) for an existing [application](https://hathora.dev/docs/concepts/hathora-entities#application). Poll the [`GetConnectionInfo()`](https://hathora.dev/api#tag/RoomV2/operation/GetConnectionInfo) endpoint to get connection details for an active room.
* [GetRoomInfo](docs/sdks/roomv2/README.md#getroominfo) - Retreive current and historical allocation data for a [room](https://hathora.dev/docs/concepts/hathora-entities#room).
* [GetActiveRoomsForProcess](docs/sdks/roomv2/README.md#getactiveroomsforprocess) - Get all active [rooms](https://hathora.dev/docs/concepts/hathora-entities#room) for a given [process](https://hathora.dev/docs/concepts/hathora-entities#process).
* [GetInactiveRoomsForProcess](docs/sdks/roomv2/README.md#getinactiveroomsforprocess) - Get all inactive [rooms](https://hathora.dev/docs/concepts/hathora-entities#room) for a given [process](https://hathora.dev/docs/concepts/hathora-entities#process).
* [DestroyRoom](docs/sdks/roomv2/README.md#destroyroom) - Destroy a [room](https://hathora.dev/docs/concepts/hathora-entities#room). All associated metadata is deleted.
* [SuspendRoom](docs/sdks/roomv2/README.md#suspendroom) - Suspend a [room](https://hathora.dev/docs/concepts/hathora-entities#room). The room is unallocated from the process but can be rescheduled later using the same `roomId`.
* [GetConnectionInfo](docs/sdks/roomv2/README.md#getconnectioninfo) - Poll this endpoint to get connection details to a [room](https://hathora.dev/docs/concepts/hathora-entities#room). Clients can call this endpoint without authentication.
* [UpdateRoomConfig](docs/sdks/roomv2/README.md#updateroomconfig)

### [OrgTokensV1](docs/sdks/orgtokensv1/README.md)

* [GetOrgTokens](docs/sdks/orgtokensv1/README.md#getorgtokens) - List all organization tokens for a given org.
* [CreateOrgToken](docs/sdks/orgtokensv1/README.md#createorgtoken) - Create a new organization token.
* [RevokeOrgToken](docs/sdks/orgtokensv1/README.md#revokeorgtoken) - Revoke an organization token.
<!-- End Available Resources and Operations [operations] -->

<!-- Start Global Parameters [global-parameters] -->
## Global Parameters

A parameter is configured globally. This parameter may be set on the SDK client instance itself during initialization. When configured as an option during SDK initialization, This global value will be used as the default on the operations that use it. When such operations are called, there is a place in each to override the global value, if needed.

For example, you can set `appId` to `"app-af469a92-5b45-4565-b3c4-b79878de67d2"` at SDK initialization and then you do not have to pass the same value on calls to operations like `UpdateApp`. But if you want to do so you may, which will locally override the global setting. See the example code below for a demonstration.


### Available Globals

The following global parameter is available.

| Name | Type | Required | Description |
| ---- | ---- |:--------:| ----------- |
| AppID | string |  | The AppID parameter. |


### Example

```go
package main

import (
	"context"
	"github.com/hathora/ci/internal/sdk"
	"github.com/hathora/ci/internal/sdk/pkg/models/shared"
	"log"
)

func main() {
	s := sdk.New(
		sdk.WithSecurity(shared.Security{
			HathoraDevToken: sdk.String("<YOUR_BEARER_TOKEN_HERE>"),
		}),
		sdk.WithAppID(sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")),
	)

	appConfig := shared.AppConfig{
		AuthConfiguration: shared.AuthConfiguration{},
		AppName:           "minecraft",
	}

	var appID *string = sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")

	ctx := context.Background()
	res, err := s.AppV1.UpdateApp(ctx, appConfig, appID)
	if err != nil {
		log.Fatal(err)
	}
	if res.Application != nil {
		// handle response
	}
}

```
<!-- End Global Parameters [global-parameters] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,422,429,500    | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

### Example

```go
package main

import (
	"context"
	"errors"
	"github.com/hathora/ci/internal/sdk"
	"github.com/hathora/ci/internal/sdk/pkg/models/sdkerrors"
	"github.com/hathora/ci/internal/sdk/pkg/models/shared"
	"log"
)

func main() {
	s := sdk.New(
		sdk.WithSecurity(shared.Security{
			HathoraDevToken: sdk.String("<YOUR_BEARER_TOKEN_HERE>"),
		}),
		sdk.WithAppID(sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")),
	)

	ctx := context.Background()
	res, err := s.AppV1.CreateApp(ctx, shared.AppConfig{
		AuthConfiguration: shared.AuthConfiguration{},
		AppName:           "minecraft",
	})
	if err != nil {

		var e *sdkerrors.APIError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://api.hathora.dev` | None |
| 1 | `https:///` | None |

#### Example

```go
package main

import (
	"context"
	"github.com/hathora/ci/internal/sdk"
	"github.com/hathora/ci/internal/sdk/pkg/models/shared"
	"log"
)

func main() {
	s := sdk.New(
		sdk.WithServerIndex(1),
		sdk.WithSecurity(shared.Security{
			HathoraDevToken: sdk.String("<YOUR_BEARER_TOKEN_HERE>"),
		}),
		sdk.WithAppID(sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")),
	)

	ctx := context.Background()
	res, err := s.AppV1.GetApps(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.ApplicationWithLatestDeploymentAndBuilds != nil {
		// handle response
	}
}

```


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	"github.com/hathora/ci/internal/sdk"
	"github.com/hathora/ci/internal/sdk/pkg/models/shared"
	"log"
)

func main() {
	s := sdk.New(
		sdk.WithServerURL("https://api.hathora.dev"),
		sdk.WithSecurity(shared.Security{
			HathoraDevToken: sdk.String("<YOUR_BEARER_TOKEN_HERE>"),
		}),
		sdk.WithAppID(sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")),
	)

	ctx := context.Background()
	res, err := s.AppV1.GetApps(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.ApplicationWithLatestDeploymentAndBuilds != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name              | Type              | Scheme            |
| ----------------- | ----------------- | ----------------- |
| `HathoraDevToken` | http              | HTTP Bearer       |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	"github.com/hathora/ci/internal/sdk"
	"github.com/hathora/ci/internal/sdk/pkg/models/shared"
	"log"
)

func main() {
	s := sdk.New(
		sdk.WithSecurity(shared.Security{
			HathoraDevToken: sdk.String("<YOUR_BEARER_TOKEN_HERE>"),
		}),
		sdk.WithAppID(sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")),
	)

	ctx := context.Background()
	res, err := s.AppV1.GetApps(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.ApplicationWithLatestDeploymentAndBuilds != nil {
		// handle response
	}
}

```

### Per-Operation Security Schemes

Some operations in this SDK require the security scheme to be specified at the request level. For example:
```go
package main

import (
	"context"
	"github.com/hathora/ci/internal/sdk"
	"github.com/hathora/ci/internal/sdk/pkg/models/operations"
	"github.com/hathora/ci/internal/sdk/pkg/models/shared"
	"log"
)

func main() {
	s := sdk.New(
		sdk.WithAppID(sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")),
	)

	var appID *string = sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")

	var region *shared.Region = shared.RegionLondon.ToPointer()

	var local *bool = sdk.Bool(false)

	operationSecurity := operations.CreatePrivateLobbyDeprecatedSecurity{
		PlayerAuth: "<YOUR_BEARER_TOKEN_HERE>",
	}

	ctx := context.Background()
	res, err := s.LobbyV1.CreatePrivateLobbyDeprecated(ctx, operationSecurity, appID, region, local)
	if err != nil {
		log.Fatal(err)
	}
	if res.RoomID != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
