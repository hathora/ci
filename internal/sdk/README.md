# github.com/hathora/cloud-sdk-go

Developer-friendly & type-safe Go SDK specifically catered to leverage *github.com/hathora/cloud-sdk-go* API.

<div align="left">
    <a href="https://www.speakeasy.com/?utm_source=github-com/hathora/cloud-sdk-go&utm_campaign=go"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>


<br /><br />
> [!IMPORTANT]
> This SDK is not yet ready for production use. To complete setup please follow the steps outlined in your [workspace](https://app.speakeasy.com/org/hathora/hathora). Delete this section before > publishing to a package manager.

<!-- Start Summary [summary] -->
## Summary

Hathora Cloud API: Welcome to the Hathora Cloud API documentation! Learn how to use the Hathora Cloud APIs to build and scale your game servers globally.
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [github.com/hathora/cloud-sdk-go](#githubcomhathoracloud-sdk-go)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Authentication](#authentication)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Global Parameters](#global-parameters)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get hathoracloud
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	"hathoracloud"
	"log"
)

func main() {
	ctx := context.Background()

	s := hathoracloud.New(
		hathoracloud.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
		hathoracloud.WithOrgID("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"),
		hathoracloud.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
	)

	res, err := s.TokensV1.GetOrgTokens(ctx, "org-6f706e83-0ec1-437a-9a46-7d4281eb2f39")
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name              | Type | Scheme      |
| ----------------- | ---- | ----------- |
| `HathoraDevToken` | http | HTTP Bearer |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	"hathoracloud"
	"log"
)

func main() {
	ctx := context.Background()

	s := hathoracloud.New(
		hathoracloud.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
		hathoracloud.WithOrgID("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"),
		hathoracloud.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
	)

	res, err := s.TokensV1.GetOrgTokens(ctx, "org-6f706e83-0ec1-437a-9a46-7d4281eb2f39")
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
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
	"hathoracloud"
	"hathoracloud/models/components"
	"hathoracloud/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := hathoracloud.New(
		hathoracloud.WithOrgID("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"),
		hathoracloud.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
	)

	res, err := s.LobbiesV3.CreateLobby(ctx, operations.CreateLobbySecurity{
		PlayerAuth: "<YOUR_BEARER_TOKEN_HERE>",
	}, components.CreateLobbyV3Params{
		Visibility: components.LobbyVisibilityPrivate,
		RoomConfig: hathoracloud.String("{\"name\":\"my-room\"}"),
		Region:     components.RegionSeattle,
	}, hathoracloud.String("app-af469a92-5b45-4565-b3c4-b79878de67d2"), hathoracloud.String("LFG4"), hathoracloud.String("2swovpy1fnunu"))
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [~~AppsV1~~](docs/sdks/appsv1/README.md)

* [~~GetAppsV1Deprecated~~](docs/sdks/appsv1/README.md#getappsv1deprecated) - GetAppsV1Deprecated :warning: **Deprecated**
* [~~CreateAppV1Deprecated~~](docs/sdks/appsv1/README.md#createappv1deprecated) - CreateAppV1Deprecated :warning: **Deprecated**
* [~~UpdateAppV1Deprecated~~](docs/sdks/appsv1/README.md#updateappv1deprecated) - UpdateAppV1Deprecated :warning: **Deprecated**
* [~~GetAppInfoV1Deprecated~~](docs/sdks/appsv1/README.md#getappinfov1deprecated) - GetAppInfoV1Deprecated :warning: **Deprecated**
* [~~DeleteAppV1Deprecated~~](docs/sdks/appsv1/README.md#deleteappv1deprecated) - DeleteAppV1Deprecated :warning: **Deprecated**

### [AppsV2](docs/sdks/appsv2/README.md)

* [GetApps](docs/sdks/appsv2/README.md#getapps) - GetApps
* [CreateApp](docs/sdks/appsv2/README.md#createapp) - CreateApp
* [GetApp](docs/sdks/appsv2/README.md#getapp) - GetApp
* [UpdateApp](docs/sdks/appsv2/README.md#updateapp) - UpdateApp
* [DeleteApp](docs/sdks/appsv2/README.md#deleteapp) - DeleteApp

### [AuthV1](docs/sdks/authv1/README.md)

* [LoginAnonymous](docs/sdks/authv1/README.md#loginanonymous) - LoginAnonymous
* [LoginNickname](docs/sdks/authv1/README.md#loginnickname) - LoginNickname
* [LoginGoogle](docs/sdks/authv1/README.md#logingoogle) - LoginGoogle

### [BillingV1](docs/sdks/billingv1/README.md)

* [GetBalance](docs/sdks/billingv1/README.md#getbalance) - GetBalance
* [GetUpcomingInvoiceItems](docs/sdks/billingv1/README.md#getupcominginvoiceitems) - GetUpcomingInvoiceItems
* [GetUpcomingInvoiceTotal](docs/sdks/billingv1/README.md#getupcominginvoicetotal) - GetUpcomingInvoiceTotal
* [GetPaymentMethod](docs/sdks/billingv1/README.md#getpaymentmethod) - GetPaymentMethod
* [InitStripeCustomerPortalURL](docs/sdks/billingv1/README.md#initstripecustomerportalurl) - InitStripeCustomerPortalUrl
* [GetInvoices](docs/sdks/billingv1/README.md#getinvoices) - GetInvoices

### [~~BuildsV1~~](docs/sdks/buildsv1/README.md)

* [~~GetBuildsDeprecated~~](docs/sdks/buildsv1/README.md#getbuildsdeprecated) - GetBuildsDeprecated :warning: **Deprecated**
* [~~GetBuildInfoDeprecated~~](docs/sdks/buildsv1/README.md#getbuildinfodeprecated) - GetBuildInfoDeprecated :warning: **Deprecated**
* [~~CreateBuildDeprecated~~](docs/sdks/buildsv1/README.md#createbuilddeprecated) - CreateBuildDeprecated :warning: **Deprecated**
* [~~DeleteBuildDeprecated~~](docs/sdks/buildsv1/README.md#deletebuilddeprecated) - DeleteBuildDeprecated :warning: **Deprecated**
* [~~RunBuildDeprecated~~](docs/sdks/buildsv1/README.md#runbuilddeprecated) - RunBuildDeprecated :warning: **Deprecated**

### [~~BuildsV2~~](docs/sdks/buildsv2/README.md)

* [~~GetBuildsV2Deprecated~~](docs/sdks/buildsv2/README.md#getbuildsv2deprecated) - GetBuildsV2Deprecated :warning: **Deprecated**
* [~~GetBuildInfoV2Deprecated~~](docs/sdks/buildsv2/README.md#getbuildinfov2deprecated) - GetBuildInfoV2Deprecated :warning: **Deprecated**
* [~~CreateBuildV2Deprecated~~](docs/sdks/buildsv2/README.md#createbuildv2deprecated) - CreateBuildV2Deprecated :warning: **Deprecated**
* [~~CreateBuildWithUploadURLV2Deprecated~~](docs/sdks/buildsv2/README.md#createbuildwithuploadurlv2deprecated) - CreateBuildWithUploadUrlV2Deprecated :warning: **Deprecated**
* [~~CreateWithMultipartUploadsV2Deprecated~~](docs/sdks/buildsv2/README.md#createwithmultipartuploadsv2deprecated) - CreateWithMultipartUploadsV2Deprecated :warning: **Deprecated**
* [~~DeleteBuildV2Deprecated~~](docs/sdks/buildsv2/README.md#deletebuildv2deprecated) - DeleteBuildV2Deprecated :warning: **Deprecated**
* [~~RunBuildV2Deprecated~~](docs/sdks/buildsv2/README.md#runbuildv2deprecated) - RunBuildV2Deprecated :warning: **Deprecated**

### [BuildsV3](docs/sdks/buildsv3/README.md)

* [GetBuilds](docs/sdks/buildsv3/README.md#getbuilds) - GetBuilds
* [CreateBuild](docs/sdks/buildsv3/README.md#createbuild) - CreateBuild
* [GetBuild](docs/sdks/buildsv3/README.md#getbuild) - GetBuild
* [DeleteBuild](docs/sdks/buildsv3/README.md#deletebuild) - DeleteBuild
* [RunBuild](docs/sdks/buildsv3/README.md#runbuild) - RunBuild

### [~~DeploymentsV1~~](docs/sdks/deploymentsv1/README.md)

* [~~GetDeploymentsV1Deprecated~~](docs/sdks/deploymentsv1/README.md#getdeploymentsv1deprecated) - GetDeploymentsV1Deprecated :warning: **Deprecated**
* [~~GetLatestDeploymentV1Deprecated~~](docs/sdks/deploymentsv1/README.md#getlatestdeploymentv1deprecated) - GetLatestDeploymentV1Deprecated :warning: **Deprecated**
* [~~GetDeploymentInfoV1Deprecated~~](docs/sdks/deploymentsv1/README.md#getdeploymentinfov1deprecated) - GetDeploymentInfoV1Deprecated :warning: **Deprecated**
* [~~CreateDeploymentV1Deprecated~~](docs/sdks/deploymentsv1/README.md#createdeploymentv1deprecated) - CreateDeploymentV1Deprecated :warning: **Deprecated**

### [~~DeploymentsV2~~](docs/sdks/deploymentsv2/README.md)

* [~~GetDeploymentsV2Deprecated~~](docs/sdks/deploymentsv2/README.md#getdeploymentsv2deprecated) - GetDeploymentsV2Deprecated :warning: **Deprecated**
* [~~GetLatestDeploymentV2Deprecated~~](docs/sdks/deploymentsv2/README.md#getlatestdeploymentv2deprecated) - GetLatestDeploymentV2Deprecated :warning: **Deprecated**
* [~~GetDeploymentInfoV2Deprecated~~](docs/sdks/deploymentsv2/README.md#getdeploymentinfov2deprecated) - GetDeploymentInfoV2Deprecated :warning: **Deprecated**
* [~~CreateDeploymentV2Deprecated~~](docs/sdks/deploymentsv2/README.md#createdeploymentv2deprecated) - CreateDeploymentV2Deprecated :warning: **Deprecated**

### [DeploymentsV3](docs/sdks/deploymentsv3/README.md)

* [GetDeployments](docs/sdks/deploymentsv3/README.md#getdeployments) - GetDeployments
* [CreateDeployment](docs/sdks/deploymentsv3/README.md#createdeployment) - CreateDeployment
* [GetLatestDeployment](docs/sdks/deploymentsv3/README.md#getlatestdeployment) - GetLatestDeployment
* [GetDeployment](docs/sdks/deploymentsv3/README.md#getdeployment) - GetDeployment

### [~~DiscoveryV1~~](docs/sdks/discoveryv1/README.md)

* [~~GetPingServiceEndpointsDeprecated~~](docs/sdks/discoveryv1/README.md#getpingserviceendpointsdeprecated) - GetPingServiceEndpointsDeprecated :warning: **Deprecated**

### [DiscoveryV2](docs/sdks/discoveryv2/README.md)

* [GetPingServiceEndpoints](docs/sdks/discoveryv2/README.md#getpingserviceendpoints) - GetPingServiceEndpoints

### [FleetsV1](docs/sdks/fleetsv1/README.md)

* [GetFleets](docs/sdks/fleetsv1/README.md#getfleets) - GetFleets
* [GetFleetRegion](docs/sdks/fleetsv1/README.md#getfleetregion) - GetFleetRegion
* [UpdateFleetRegion](docs/sdks/fleetsv1/README.md#updatefleetregion) - UpdateFleetRegion
* [GetFleetMetrics](docs/sdks/fleetsv1/README.md#getfleetmetrics) - GetFleetMetrics


### [~~LobbiesV1~~](docs/sdks/lobbiesv1/README.md)

* [~~CreatePrivateLobbyDeprecated~~](docs/sdks/lobbiesv1/README.md#createprivatelobbydeprecated) - CreatePrivateLobbyDeprecated :warning: **Deprecated**
* [~~CreatePublicLobbyDeprecated~~](docs/sdks/lobbiesv1/README.md#createpubliclobbydeprecated) - CreatePublicLobbyDeprecated :warning: **Deprecated**
* [~~ListActivePublicLobbiesDeprecatedV1~~](docs/sdks/lobbiesv1/README.md#listactivepubliclobbiesdeprecatedv1) - ListActivePublicLobbiesDeprecatedV1 :warning: **Deprecated**

### [~~LobbiesV2~~](docs/sdks/lobbiesv2/README.md)

* [~~CreatePrivateLobby~~](docs/sdks/lobbiesv2/README.md#createprivatelobby) - CreatePrivateLobby :warning: **Deprecated**
* [~~CreatePublicLobby~~](docs/sdks/lobbiesv2/README.md#createpubliclobby) - CreatePublicLobby :warning: **Deprecated**
* [~~CreateLocalLobby~~](docs/sdks/lobbiesv2/README.md#createlocallobby) - CreateLocalLobby :warning: **Deprecated**
* [~~CreateLobbyDeprecated~~](docs/sdks/lobbiesv2/README.md#createlobbydeprecated) - CreateLobbyDeprecated :warning: **Deprecated**
* [~~ListActivePublicLobbiesDeprecatedV2~~](docs/sdks/lobbiesv2/README.md#listactivepubliclobbiesdeprecatedv2) - ListActivePublicLobbiesDeprecatedV2 :warning: **Deprecated**
* [~~GetLobbyInfo~~](docs/sdks/lobbiesv2/README.md#getlobbyinfo) - GetLobbyInfo :warning: **Deprecated**
* [~~SetLobbyState~~](docs/sdks/lobbiesv2/README.md#setlobbystate) - SetLobbyState :warning: **Deprecated**

### [LobbiesV3](docs/sdks/lobbiesv3/README.md)

* [CreateLobby](docs/sdks/lobbiesv3/README.md#createlobby) - CreateLobby
* [ListActivePublicLobbies](docs/sdks/lobbiesv3/README.md#listactivepubliclobbies) - ListActivePublicLobbies
* [GetLobbyInfoByRoomID](docs/sdks/lobbiesv3/README.md#getlobbyinfobyroomid) - GetLobbyInfoByRoomId
* [GetLobbyInfoByShortCode](docs/sdks/lobbiesv3/README.md#getlobbyinfobyshortcode) - GetLobbyInfoByShortCode

### [LogsV1](docs/sdks/logsv1/README.md)

* [GetLogsForProcess](docs/sdks/logsv1/README.md#getlogsforprocess) - GetLogsForProcess
* [DownloadLogForProcess](docs/sdks/logsv1/README.md#downloadlogforprocess) - DownloadLogForProcess

### [ManagementV1](docs/sdks/managementv1/README.md)

* [SendVerificationEmail](docs/sdks/managementv1/README.md#sendverificationemail) - SendVerificationEmail

### [~~MetricsV1~~](docs/sdks/metricsv1/README.md)

* [~~GetMetricsDeprecated~~](docs/sdks/metricsv1/README.md#getmetricsdeprecated) - GetMetricsDeprecated :warning: **Deprecated**

### [OrganizationsV1](docs/sdks/organizationsv1/README.md)

* [GetOrgs](docs/sdks/organizationsv1/README.md#getorgs) - GetOrgs
* [GetUserPendingInvites](docs/sdks/organizationsv1/README.md#getuserpendinginvites) - GetUserPendingInvites
* [GetOrgMembers](docs/sdks/organizationsv1/README.md#getorgmembers) - GetOrgMembers
* [InviteUser](docs/sdks/organizationsv1/README.md#inviteuser) - InviteUser
* [UpdateUserInvite](docs/sdks/organizationsv1/README.md#updateuserinvite) - UpdateUserInvite
* [RescindInvite](docs/sdks/organizationsv1/README.md#rescindinvite) - RescindInvite
* [GetOrgPendingInvites](docs/sdks/organizationsv1/README.md#getorgpendinginvites) - GetOrgPendingInvites
* [AcceptInvite](docs/sdks/organizationsv1/README.md#acceptinvite) - AcceptInvite
* [RejectInvite](docs/sdks/organizationsv1/README.md#rejectinvite) - RejectInvite
* [GetUsageLimits](docs/sdks/organizationsv1/README.md#getusagelimits) - GetUsageLimits

### [~~ProcessesV1~~](docs/sdks/processesv1/README.md)

* [~~GetRunningProcesses~~](docs/sdks/processesv1/README.md#getrunningprocesses) - GetRunningProcesses :warning: **Deprecated**
* [~~GetStoppedProcesses~~](docs/sdks/processesv1/README.md#getstoppedprocesses) - GetStoppedProcesses :warning: **Deprecated**
* [~~GetProcessInfoDeprecated~~](docs/sdks/processesv1/README.md#getprocessinfodeprecated) - GetProcessInfoDeprecated :warning: **Deprecated**

### [~~ProcessesV2~~](docs/sdks/processesv2/README.md)

* [~~GetProcessInfoV2Deprecated~~](docs/sdks/processesv2/README.md#getprocessinfov2deprecated) - GetProcessInfoV2Deprecated :warning: **Deprecated**
* [~~GetLatestProcessesV2Deprecated~~](docs/sdks/processesv2/README.md#getlatestprocessesv2deprecated) - GetLatestProcessesV2Deprecated :warning: **Deprecated**
* [~~GetProcessesCountExperimentalV2Deprecated~~](docs/sdks/processesv2/README.md#getprocessescountexperimentalv2deprecated) - GetProcessesCountExperimentalV2Deprecated :warning: **Deprecated**
* [~~StopProcessV2Deprecated~~](docs/sdks/processesv2/README.md#stopprocessv2deprecated) - StopProcessV2Deprecated :warning: **Deprecated**
* [~~CreateProcessV2Deprecated~~](docs/sdks/processesv2/README.md#createprocessv2deprecated) - CreateProcessV2Deprecated :warning: **Deprecated**

### [ProcessesV3](docs/sdks/processesv3/README.md)

* [GetLatestProcesses](docs/sdks/processesv3/README.md#getlatestprocesses) - GetLatestProcesses
* [GetProcessesCountExperimental](docs/sdks/processesv3/README.md#getprocessescountexperimental) - GetProcessesCountExperimental
* [CreateProcess](docs/sdks/processesv3/README.md#createprocess) - CreateProcess
* [GetProcess](docs/sdks/processesv3/README.md#getprocess) - GetProcess
* [StopProcess](docs/sdks/processesv3/README.md#stopprocess) - StopProcess
* [GetProcessMetrics](docs/sdks/processesv3/README.md#getprocessmetrics) - GetProcessMetrics

### [~~RoomsV1~~](docs/sdks/roomsv1/README.md)

* [~~CreateRoomDeprecated~~](docs/sdks/roomsv1/README.md#createroomdeprecated) - CreateRoomDeprecated :warning: **Deprecated**
* [~~GetRoomInfoDeprecated~~](docs/sdks/roomsv1/README.md#getroominfodeprecated) - GetRoomInfoDeprecated :warning: **Deprecated**
* [~~GetActiveRoomsForProcessDeprecated~~](docs/sdks/roomsv1/README.md#getactiveroomsforprocessdeprecated) - GetActiveRoomsForProcessDeprecated :warning: **Deprecated**
* [~~GetInactiveRoomsForProcessDeprecated~~](docs/sdks/roomsv1/README.md#getinactiveroomsforprocessdeprecated) - GetInactiveRoomsForProcessDeprecated :warning: **Deprecated**
* [~~DestroyRoomDeprecated~~](docs/sdks/roomsv1/README.md#destroyroomdeprecated) - DestroyRoomDeprecated :warning: **Deprecated**
* [~~SuspendRoomDeprecated~~](docs/sdks/roomsv1/README.md#suspendroomdeprecated) - SuspendRoomDeprecated :warning: **Deprecated**
* [~~GetConnectionInfoDeprecated~~](docs/sdks/roomsv1/README.md#getconnectioninfodeprecated) - GetConnectionInfoDeprecated :warning: **Deprecated**

### [RoomsV2](docs/sdks/roomsv2/README.md)

* [CreateRoom](docs/sdks/roomsv2/README.md#createroom) - CreateRoom
* [GetRoomInfo](docs/sdks/roomsv2/README.md#getroominfo) - GetRoomInfo
* [GetActiveRoomsForProcess](docs/sdks/roomsv2/README.md#getactiveroomsforprocess) - GetActiveRoomsForProcess
* [GetInactiveRoomsForProcess](docs/sdks/roomsv2/README.md#getinactiveroomsforprocess) - GetInactiveRoomsForProcess
* [DestroyRoom](docs/sdks/roomsv2/README.md#destroyroom) - DestroyRoom
* [~~SuspendRoomV2Deprecated~~](docs/sdks/roomsv2/README.md#suspendroomv2deprecated) - SuspendRoomV2Deprecated :warning: **Deprecated**
* [GetConnectionInfo](docs/sdks/roomsv2/README.md#getconnectioninfo) - GetConnectionInfo
* [UpdateRoomConfig](docs/sdks/roomsv2/README.md#updateroomconfig) - UpdateRoomConfig

### [TokensV1](docs/sdks/tokensv1/README.md)

* [GetOrgTokens](docs/sdks/tokensv1/README.md#getorgtokens) - GetOrgTokens
* [CreateOrgToken](docs/sdks/tokensv1/README.md#createorgtoken) - CreateOrgToken
* [RevokeOrgToken](docs/sdks/tokensv1/README.md#revokeorgtoken) - RevokeOrgToken

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Global Parameters [global-parameters] -->
## Global Parameters

Certain parameters are configured globally. These parameters may be set on the SDK client instance itself during initialization. When configured as an option during SDK initialization, These global values will be used as defaults on the operations that use them. When such operations are called, there is a place in each to override the global value, if needed.

For example, you can set `orgId` to `"org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"` at SDK initialization and then you do not have to pass the same value on calls to operations like `GetOrgTokens`. But if you want to do so you may, which will locally override the global setting. See the example code below for a demonstration.


### Available Globals

The following global parameters are available.

| Name  | Type   | Description          |
| ----- | ------ | -------------------- |
| OrgID | string | The OrgID parameter. |
| AppID | string | The AppID parameter. |

### Example

```go
package main

import (
	"context"
	"hathoracloud"
	"log"
)

func main() {
	ctx := context.Background()

	s := hathoracloud.New(
		hathoracloud.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
		hathoracloud.WithOrgID("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"),
		hathoracloud.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
	)

	res, err := s.TokensV1.GetOrgTokens(ctx, "org-6f706e83-0ec1-437a-9a46-7d4281eb2f39")
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End Global Parameters [global-parameters] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	"hathoracloud"
	"hathoracloud/retry"
	"log"
	"models/operations"
)

func main() {
	ctx := context.Background()

	s := hathoracloud.New(
		hathoracloud.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
		hathoracloud.WithOrgID("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"),
		hathoracloud.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
	)

	res, err := s.TokensV1.GetOrgTokens(ctx, "org-6f706e83-0ec1-437a-9a46-7d4281eb2f39", operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	"hathoracloud"
	"hathoracloud/retry"
	"log"
)

func main() {
	ctx := context.Background()

	s := hathoracloud.New(
		hathoracloud.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		hathoracloud.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
		hathoracloud.WithOrgID("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"),
		hathoracloud.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
	)

	res, err := s.TokensV1.GetOrgTokens(ctx, "org-6f706e83-0ec1-437a-9a46-7d4281eb2f39")
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `errors.SDKError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `GetOrgTokens` function may return the following errors:

| Error Type      | Status Code   | Content Type     |
| --------------- | ------------- | ---------------- |
| errors.APIError | 401, 404, 429 | application/json |
| errors.SDKError | 4XX, 5XX      | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	"hathoracloud"
	"hathoracloud/models/errors"
	"log"
)

func main() {
	ctx := context.Background()

	s := hathoracloud.New(
		hathoracloud.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
		hathoracloud.WithOrgID("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"),
		hathoracloud.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
	)

	res, err := s.TokensV1.GetOrgTokens(ctx, "org-6f706e83-0ec1-437a-9a46-7d4281eb2f39")
	if err != nil {

		var e *errors.APIError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *errors.SDKError
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

You can override the default server globally using the `WithServerIndex(serverIndex int)` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| #   | Server                    |
| --- | ------------------------- |
| 0   | `https://api.hathora.dev` |
| 1   | `https:///`               |

#### Example

```go
package main

import (
	"context"
	"hathoracloud"
	"log"
)

func main() {
	ctx := context.Background()

	s := hathoracloud.New(
		hathoracloud.WithServerIndex(1),
		hathoracloud.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
		hathoracloud.WithOrgID("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"),
		hathoracloud.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
	)

	res, err := s.TokensV1.GetOrgTokens(ctx, "org-6f706e83-0ec1-437a-9a46-7d4281eb2f39")
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	"hathoracloud"
	"log"
)

func main() {
	ctx := context.Background()

	s := hathoracloud.New(
		hathoracloud.WithServerURL("https://api.hathora.dev"),
		hathoracloud.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
		hathoracloud.WithOrgID("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"),
		hathoracloud.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
	)

	res, err := s.TokensV1.GetOrgTokens(ctx, "org-6f706e83-0ec1-437a-9a46-7d4281eb2f39")
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
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

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=github-com/hathora/cloud-sdk-go&utm_campaign=go)
