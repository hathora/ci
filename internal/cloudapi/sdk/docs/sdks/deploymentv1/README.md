# DeploymentV1
(*DeploymentV1*)

## Overview

Deprecated. Use [DeploymentV2](https://hathora.dev/api#tag/DeploymentV2).

### Available Operations

* [~~GetDeploymentsDeprecated~~](#getdeploymentsdeprecated) - Returns an array of [deployments](https://hathora.dev/docs/concepts/hathora-entities#deployment) for an [application](https://hathora.dev/docs/concepts/hathora-entities#application). :warning: **Deprecated**
* [~~GetLatestDeploymentDeprecated~~](#getlatestdeploymentdeprecated) - Get the latest [deployment](https://hathora.dev/docs/concepts/hathora-entities#deployment) for an [application](https://hathora.dev/docs/concepts/hathora-entities#application). :warning: **Deprecated**
* [~~GetDeploymentInfoDeprecated~~](#getdeploymentinfodeprecated) - Get details for a [deployment](https://hathora.dev/docs/concepts/hathora-entities#deployment). :warning: **Deprecated**
* [~~CreateDeploymentDeprecated~~](#createdeploymentdeprecated) - Create a new [deployment](https://hathora.dev/docs/concepts/hathora-entities#deployment). Creating a new deployment means all new rooms created will use the latest deployment configuration, but existing games in progress will not be affected. :warning: **Deprecated**

## ~~GetDeploymentsDeprecated~~

Returns an array of [deployments](https://hathora.dev/docs/concepts/hathora-entities#deployment) for an [application](https://hathora.dev/docs/concepts/hathora-entities#application).

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"cloudapi/pkg/models/shared"
	"cloudapi"
	"context"
	"log"
)

func main() {
    s := cloudapi.New(
        cloudapi.WithSecurity(shared.Security{
            HathoraDevToken: cloudapi.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
        cloudapi.WithAppID(cloudapi.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")),
    )


    var appID *string = cloudapi.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")

    ctx := context.Background()
    res, err := s.DeploymentV1.GetDeploymentsDeprecated(ctx, appID)
    if err != nil {
        log.Fatal(err)
    }
    if res.Deployments != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `appID`                                               | **string*                                             | :heavy_minus_sign:                                    | N/A                                                   | app-af469a92-5b45-4565-b3c4-b79878de67d2              |


### Response

**[*operations.GetDeploymentsDeprecatedResponse](../../pkg/models/operations/getdeploymentsdeprecatedresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,404            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ~~GetLatestDeploymentDeprecated~~

Get the latest [deployment](https://hathora.dev/docs/concepts/hathora-entities#deployment) for an [application](https://hathora.dev/docs/concepts/hathora-entities#application).

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"cloudapi/pkg/models/shared"
	"cloudapi"
	"context"
	"log"
)

func main() {
    s := cloudapi.New(
        cloudapi.WithSecurity(shared.Security{
            HathoraDevToken: cloudapi.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
        cloudapi.WithAppID(cloudapi.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")),
    )


    var appID *string = cloudapi.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")

    ctx := context.Background()
    res, err := s.DeploymentV1.GetLatestDeploymentDeprecated(ctx, appID)
    if err != nil {
        log.Fatal(err)
    }
    if res.Deployment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `appID`                                               | **string*                                             | :heavy_minus_sign:                                    | N/A                                                   | app-af469a92-5b45-4565-b3c4-b79878de67d2              |


### Response

**[*operations.GetLatestDeploymentDeprecatedResponse](../../pkg/models/operations/getlatestdeploymentdeprecatedresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,404            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ~~GetDeploymentInfoDeprecated~~

Get details for a [deployment](https://hathora.dev/docs/concepts/hathora-entities#deployment).

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"cloudapi/pkg/models/shared"
	"cloudapi"
	"context"
	"log"
)

func main() {
    s := cloudapi.New(
        cloudapi.WithSecurity(shared.Security{
            HathoraDevToken: cloudapi.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
        cloudapi.WithAppID(cloudapi.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")),
    )


    var deploymentID int = 1

    var appID *string = cloudapi.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")

    ctx := context.Background()
    res, err := s.DeploymentV1.GetDeploymentInfoDeprecated(ctx, deploymentID, appID)
    if err != nil {
        log.Fatal(err)
    }
    if res.Deployment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `deploymentID`                                        | *int*                                                 | :heavy_check_mark:                                    | N/A                                                   | 1                                                     |
| `appID`                                               | **string*                                             | :heavy_minus_sign:                                    | N/A                                                   | app-af469a92-5b45-4565-b3c4-b79878de67d2              |


### Response

**[*operations.GetDeploymentInfoDeprecatedResponse](../../pkg/models/operations/getdeploymentinfodeprecatedresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,404            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ~~CreateDeploymentDeprecated~~

Create a new [deployment](https://hathora.dev/docs/concepts/hathora-entities#deployment). Creating a new deployment means all new rooms created will use the latest deployment configuration, but existing games in progress will not be affected.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"cloudapi/pkg/models/shared"
	"cloudapi"
	"context"
	"log"
)

func main() {
    s := cloudapi.New(
        cloudapi.WithSecurity(shared.Security{
            HathoraDevToken: cloudapi.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
        cloudapi.WithAppID(cloudapi.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")),
    )


    var buildID int = 1

    deploymentConfig := shared.DeploymentConfig{
        Env: []shared.DeploymentConfigEnv{
            shared.DeploymentConfigEnv{
                Value: "TRUE",
                Name: "EULA",
            },
        },
        RoomsPerProcess: 3,
        PlanName: shared.PlanNameTiny,
        AdditionalContainerPorts: []shared.ContainerPort{
            shared.ContainerPort{
                TransportType: shared.TransportTypeUDP,
                Port: 8000,
                Name: "default",
            },
        },
        TransportType: shared.TransportTypeTLS,
        ContainerPort: 4000,
    }

    var appID *string = cloudapi.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")

    ctx := context.Background()
    res, err := s.DeploymentV1.CreateDeploymentDeprecated(ctx, buildID, deploymentConfig, appID)
    if err != nil {
        log.Fatal(err)
    }
    if res.Deployment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            | Example                                                                |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |                                                                        |
| `buildID`                                                              | *int*                                                                  | :heavy_check_mark:                                                     | N/A                                                                    | 1                                                                      |
| `deploymentConfig`                                                     | [shared.DeploymentConfig](../../pkg/models/shared/deploymentconfig.md) | :heavy_check_mark:                                                     | N/A                                                                    |                                                                        |
| `appID`                                                                | **string*                                                              | :heavy_minus_sign:                                                     | N/A                                                                    | app-af469a92-5b45-4565-b3c4-b79878de67d2                               |


### Response

**[*operations.CreateDeploymentDeprecatedResponse](../../pkg/models/operations/createdeploymentdeprecatedresponse.md), error**
| Error Object        | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| sdkerrors.APIError  | 400,401,404,429,500 | application/json    |
| sdkerrors.SDKError  | 4xx-5xx             | */*                 |
