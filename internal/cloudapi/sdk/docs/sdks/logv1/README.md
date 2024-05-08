# LogV1
(*LogV1*)

## Overview

Operations to get logs by [applications](https://hathora.dev/docs/concepts/hathora-entities#application), [processes](https://hathora.dev/docs/concepts/hathora-entities#process), and [deployments](https://hathora.dev/docs/concepts/hathora-entities#deployment). We store 20GB of logs data.

### Available Operations

* [~~GetLogsForApp~~](#getlogsforapp) - Returns a stream of logs for an [application](https://hathora.dev/docs/concepts/hathora-entities#application) using `appId`. :warning: **Deprecated**
* [GetLogsForProcess](#getlogsforprocess) - Returns a stream of logs for a [process](https://hathora.dev/docs/concepts/hathora-entities#process) using `appId` and `processId`.
* [DownloadLogForProcess](#downloadlogforprocess) - Download entire log file for a stopped process.
* [~~GetLogsForDeployment~~](#getlogsfordeployment) - Returns a stream of logs for a [deployment](https://hathora.dev/docs/concepts/hathora-entities#deployment) using `appId` and `deploymentId`. :warning: **Deprecated**

## ~~GetLogsForApp~~

Returns a stream of logs for an [application](https://hathora.dev/docs/concepts/hathora-entities#application) using `appId`.

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

    var follow *bool = cloudapi.Bool(true)

    var tailLines *int = cloudapi.Int(100)

    var region *shared.Region = shared.RegionSingapore.ToPointer()

    ctx := context.Background()
    res, err := s.LogV1.GetLogsForApp(ctx, appID, follow, tailLines, region)
    if err != nil {
        log.Fatal(err)
    }
    if res.Stream != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `appID`                                               | **string*                                             | :heavy_minus_sign:                                    | N/A                                                   | app-af469a92-5b45-4565-b3c4-b79878de67d2              |
| `follow`                                              | **bool*                                               | :heavy_minus_sign:                                    | N/A                                                   |                                                       |
| `tailLines`                                           | **int*                                                | :heavy_minus_sign:                                    | N/A                                                   | 100                                                   |
| `region`                                              | [*shared.Region](../../pkg/models/shared/region.md)   | :heavy_minus_sign:                                    | N/A                                                   |                                                       |


### Response

**[*operations.GetLogsForAppResponse](../../pkg/models/operations/getlogsforappresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,404            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetLogsForProcess

Returns a stream of logs for a [process](https://hathora.dev/docs/concepts/hathora-entities#process) using `appId` and `processId`.

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


    var processID string = "cbfcddd2-0006-43ae-996c-995fff7bed2e"

    var appID *string = cloudapi.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")

    var follow *bool = cloudapi.Bool(true)

    var tailLines *int = cloudapi.Int(100)

    ctx := context.Background()
    res, err := s.LogV1.GetLogsForProcess(ctx, processID, appID, follow, tailLines)
    if err != nil {
        log.Fatal(err)
    }
    if res.Stream != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `processID`                                           | *string*                                              | :heavy_check_mark:                                    | N/A                                                   | cbfcddd2-0006-43ae-996c-995fff7bed2e                  |
| `appID`                                               | **string*                                             | :heavy_minus_sign:                                    | N/A                                                   | app-af469a92-5b45-4565-b3c4-b79878de67d2              |
| `follow`                                              | **bool*                                               | :heavy_minus_sign:                                    | N/A                                                   |                                                       |
| `tailLines`                                           | **int*                                                | :heavy_minus_sign:                                    | N/A                                                   | 100                                                   |


### Response

**[*operations.GetLogsForProcessResponse](../../pkg/models/operations/getlogsforprocessresponse.md), error**
| Error Object        | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| sdkerrors.APIError  | 400,401,404,410,500 | application/json    |
| sdkerrors.SDKError  | 4xx-5xx             | */*                 |

## DownloadLogForProcess

Download entire log file for a stopped process.

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


    var processID string = "cbfcddd2-0006-43ae-996c-995fff7bed2e"

    var appID *string = cloudapi.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")

    ctx := context.Background()
    res, err := s.LogV1.DownloadLogForProcess(ctx, processID, appID)
    if err != nil {
        log.Fatal(err)
    }
    if res.Stream != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `processID`                                           | *string*                                              | :heavy_check_mark:                                    | N/A                                                   | cbfcddd2-0006-43ae-996c-995fff7bed2e                  |
| `appID`                                               | **string*                                             | :heavy_minus_sign:                                    | N/A                                                   | app-af469a92-5b45-4565-b3c4-b79878de67d2              |


### Response

**[*operations.DownloadLogForProcessResponse](../../pkg/models/operations/downloadlogforprocessresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 400,401,404,410    | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ~~GetLogsForDeployment~~

Returns a stream of logs for a [deployment](https://hathora.dev/docs/concepts/hathora-entities#deployment) using `appId` and `deploymentId`.

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

    var follow *bool = cloudapi.Bool(true)

    var tailLines *int = cloudapi.Int(100)

    ctx := context.Background()
    res, err := s.LogV1.GetLogsForDeployment(ctx, deploymentID, appID, follow, tailLines)
    if err != nil {
        log.Fatal(err)
    }
    if res.Stream != nil {
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
| `follow`                                              | **bool*                                               | :heavy_minus_sign:                                    | N/A                                                   |                                                       |
| `tailLines`                                           | **int*                                                | :heavy_minus_sign:                                    | N/A                                                   | 100                                                   |


### Response

**[*operations.GetLogsForDeploymentResponse](../../pkg/models/operations/getlogsfordeploymentresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,404            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
