# ProcessesV2
(*ProcessesV2*)

## Overview

Operations to get data on active and stopped [processes](https://hathora.dev/docs/concepts/hathora-entities#process).

### Available Operations

* [GetProcessInfo](#getprocessinfo) - Get details for a [process](https://hathora.dev/docs/concepts/hathora-entities#process).
* [GetLatestProcesses](#getlatestprocesses) - Retrieve the 10 most recent [processes](https://hathora.dev/docs/concepts/hathora-entities#process) objects for an [application](https://hathora.dev/docs/concepts/hathora-entities#application). Filter the array by optionally passing in a `status` or `region`.
* [StopProcess](#stopprocess) - Stops a [process](https://hathora.dev/docs/concepts/hathora-entities#process) immediately.
* [CreateProcess](#createprocess) - Creates a [process](https://hathora.dev/docs/concepts/hathora-entities#process) without a room. Use this to pre-allocate processes ahead of time so that subsequent room assignment via [CreateRoom()](https://hathora.dev/api#tag/RoomV2/operation/CreateRoom) can be instant.

## GetProcessInfo

Get details for a [process](https://hathora.dev/docs/concepts/hathora-entities#process).

### Example Usage

```go
package main

import(
	"github.com/hathora/ci/internal/sdk/models/shared"
	"os"
	"github.com/hathora/ci/internal/sdk"
	"context"
	"log"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            HathoraDevToken: sdk.String(os.Getenv("HATHORA_DEV_TOKEN")),
        }),
        sdk.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
    )
    var processID string = "cbfcddd2-0006-43ae-996c-995fff7bed2e"

    var appID *string = sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")
    ctx := context.Background()
    res, err := s.ProcessesV2.GetProcessInfo(ctx, processID, appID)
    if err != nil {
        log.Fatal(err)
    }
    if res.ProcessV2 != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `processID`                                              | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | cbfcddd2-0006-43ae-996c-995fff7bed2e                     |
| `appID`                                                  | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | app-af469a92-5b45-4565-b3c4-b79878de67d2                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |


### Response

**[*operations.GetProcessInfoResponse](../../models/operations/getprocessinforesponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,404,429        | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetLatestProcesses

Retrieve the 10 most recent [processes](https://hathora.dev/docs/concepts/hathora-entities#process) objects for an [application](https://hathora.dev/docs/concepts/hathora-entities#application). Filter the array by optionally passing in a `status` or `region`.

### Example Usage

```go
package main

import(
	"github.com/hathora/ci/internal/sdk/models/shared"
	"os"
	"github.com/hathora/ci/internal/sdk"
	"context"
	"log"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            HathoraDevToken: sdk.String(os.Getenv("HATHORA_DEV_TOKEN")),
        }),
        sdk.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
    )
    var appID *string = sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")

    var status []shared.ProcessStatus = []shared.ProcessStatus{
        shared.ProcessStatusStopped,
    }

    var region []shared.Region = []shared.Region{
        shared.RegionFrankfurt,
    }
    ctx := context.Background()
    res, err := s.ProcessesV2.GetLatestProcesses(ctx, appID, status, region)
    if err != nil {
        log.Fatal(err)
    }
    if res.ProcessV2s != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    | Example                                                        |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |                                                                |
| `appID`                                                        | **string*                                                      | :heavy_minus_sign:                                             | N/A                                                            | app-af469a92-5b45-4565-b3c4-b79878de67d2                       |
| `status`                                                       | [][shared.ProcessStatus](../../models/shared/processstatus.md) | :heavy_minus_sign:                                             | N/A                                                            |                                                                |
| `region`                                                       | [][shared.Region](../../models/shared/region.md)               | :heavy_minus_sign:                                             | N/A                                                            |                                                                |
| `opts`                                                         | [][operations.Option](../../models/operations/option.md)       | :heavy_minus_sign:                                             | The options for this request.                                  |                                                                |


### Response

**[*operations.GetLatestProcessesResponse](../../models/operations/getlatestprocessesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,404,429        | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## StopProcess

Stops a [process](https://hathora.dev/docs/concepts/hathora-entities#process) immediately.

### Example Usage

```go
package main

import(
	"github.com/hathora/ci/internal/sdk/models/shared"
	"os"
	"github.com/hathora/ci/internal/sdk"
	"context"
	"log"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            HathoraDevToken: sdk.String(os.Getenv("HATHORA_DEV_TOKEN")),
        }),
        sdk.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
    )
    var processID string = "cbfcddd2-0006-43ae-996c-995fff7bed2e"

    var appID *string = sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")
    ctx := context.Background()
    res, err := s.ProcessesV2.StopProcess(ctx, processID, appID)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `processID`                                              | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | cbfcddd2-0006-43ae-996c-995fff7bed2e                     |
| `appID`                                                  | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | app-af469a92-5b45-4565-b3c4-b79878de67d2                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |


### Response

**[*operations.StopProcessResponse](../../models/operations/stopprocessresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,404,429,500    | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateProcess

Creates a [process](https://hathora.dev/docs/concepts/hathora-entities#process) without a room. Use this to pre-allocate processes ahead of time so that subsequent room assignment via [CreateRoom()](https://hathora.dev/api#tag/RoomV2/operation/CreateRoom) can be instant.

### Example Usage

```go
package main

import(
	"github.com/hathora/ci/internal/sdk/models/shared"
	"os"
	"github.com/hathora/ci/internal/sdk"
	"context"
	"log"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            HathoraDevToken: sdk.String(os.Getenv("HATHORA_DEV_TOKEN")),
        }),
        sdk.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
    )
    var region shared.Region = shared.RegionTokyo

    var appID *string = sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")
    ctx := context.Background()
    res, err := s.ProcessesV2.CreateProcess(ctx, region, appID)
    if err != nil {
        log.Fatal(err)
    }
    if res.ProcessV2 != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `region`                                                 | [shared.Region](../../models/shared/region.md)           | :heavy_check_mark:                                       | N/A                                                      |                                                          |
| `appID`                                                  | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | app-af469a92-5b45-4565-b3c4-b79878de67d2                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |


### Response

**[*operations.CreateProcessResponse](../../models/operations/createprocessresponse.md), error**
| Error Object        | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| sdkerrors.APIError  | 401,402,404,429,500 | application/json    |
| sdkerrors.SDKError  | 4xx-5xx             | */*                 |
