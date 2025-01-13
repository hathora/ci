# ProcessesV2
(*ProcessesV2*)

## Overview

Deprecated. Use [ProcessesV3](https://hathora.dev/api#tag/ProcessesV3).

### Available Operations

* [~~GetProcessInfoV2Deprecated~~](#getprocessinfov2deprecated) - GetProcessInfoV2Deprecated :warning: **Deprecated**
* [~~GetLatestProcessesV2Deprecated~~](#getlatestprocessesv2deprecated) - GetLatestProcessesV2Deprecated :warning: **Deprecated**
* [~~GetProcessesCountExperimentalV2Deprecated~~](#getprocessescountexperimentalv2deprecated) - GetProcessesCountExperimentalV2Deprecated :warning: **Deprecated**
* [~~StopProcessV2Deprecated~~](#stopprocessv2deprecated) - StopProcessV2Deprecated :warning: **Deprecated**
* [~~CreateProcessV2Deprecated~~](#createprocessv2deprecated) - CreateProcessV2Deprecated :warning: **Deprecated**

## ~~GetProcessInfoV2Deprecated~~

Get details for a [process](https://hathora.dev/docs/concepts/hathora-entities#process).

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
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

    res, err := s.ProcessesV2.GetProcessInfoV2Deprecated(ctx, "cbfcddd2-0006-43ae-996c-995fff7bed2e", hathoracloud.String("app-af469a92-5b45-4565-b3c4-b79878de67d2"))
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

**[*components.ProcessV2](../../models/components/processv2.md), error**

### Errors

| Error Type       | Status Code      | Content Type     |
| ---------------- | ---------------- | ---------------- |
| errors.APIError  | 401, 404, 429    | application/json |
| errors.SDKError  | 4XX, 5XX         | \*/\*            |

## ~~GetLatestProcessesV2Deprecated~~

Retrieve the 10 most recent [processes](https://hathora.dev/docs/concepts/hathora-entities#process) objects for an [application](https://hathora.dev/docs/concepts/hathora-entities#application). Filter the array by optionally passing in a `status` or `region`.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
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

    res, err := s.ProcessesV2.GetLatestProcessesV2Deprecated(ctx, hathoracloud.String("app-af469a92-5b45-4565-b3c4-b79878de67d2"), nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            | Example                                                                |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |                                                                        |
| `appID`                                                                | **string*                                                              | :heavy_minus_sign:                                                     | N/A                                                                    | app-af469a92-5b45-4565-b3c4-b79878de67d2                               |
| `status`                                                               | [][components.ProcessStatus](../../models/components/processstatus.md) | :heavy_minus_sign:                                                     | N/A                                                                    |                                                                        |
| `region`                                                               | [][components.Region](../../models/components/region.md)               | :heavy_minus_sign:                                                     | N/A                                                                    |                                                                        |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |                                                                        |

### Response

**[[]components.ProcessV2](../../.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| errors.APIError    | 401, 404, 422, 429 | application/json   |
| errors.SDKError    | 4XX, 5XX           | \*/\*              |

## ~~GetProcessesCountExperimentalV2Deprecated~~

Count the number of [processes](https://hathora.dev/docs/concepts/hathora-entities#process) objects for an [application](https://hathora.dev/docs/concepts/hathora-entities#application). Filter by optionally passing in a `status` or `region`.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
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

    res, err := s.ProcessesV2.GetProcessesCountExperimentalV2Deprecated(ctx, hathoracloud.String("app-af469a92-5b45-4565-b3c4-b79878de67d2"), nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            | Example                                                                |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |                                                                        |
| `appID`                                                                | **string*                                                              | :heavy_minus_sign:                                                     | N/A                                                                    | app-af469a92-5b45-4565-b3c4-b79878de67d2                               |
| `status`                                                               | [][components.ProcessStatus](../../models/components/processstatus.md) | :heavy_minus_sign:                                                     | N/A                                                                    |                                                                        |
| `region`                                                               | [][components.Region](../../models/components/region.md)               | :heavy_minus_sign:                                                     | N/A                                                                    |                                                                        |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |                                                                        |

### Response

**[*operations.GetProcessesCountExperimentalV2DeprecatedResponseBody](../../models/operations/getprocessescountexperimentalv2deprecatedresponsebody.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| errors.APIError    | 401, 404, 422, 429 | application/json   |
| errors.SDKError    | 4XX, 5XX           | \*/\*              |

## ~~StopProcessV2Deprecated~~

Stops a [process](https://hathora.dev/docs/concepts/hathora-entities#process) immediately.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
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

    err := s.ProcessesV2.StopProcessV2Deprecated(ctx, "cbfcddd2-0006-43ae-996c-995fff7bed2e", hathoracloud.String("app-af469a92-5b45-4565-b3c4-b79878de67d2"))
    if err != nil {
        log.Fatal(err)
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

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| errors.APIError    | 401, 404, 429, 500 | application/json   |
| errors.SDKError    | 4XX, 5XX           | \*/\*              |

## ~~CreateProcessV2Deprecated~~

Creates a [process](https://hathora.dev/docs/concepts/hathora-entities#process) without a room. Use this to pre-allocate processes ahead of time so that subsequent room assignment via [CreateRoom()](https://hathora.dev/api#tag/RoomV2/operation/CreateRoom) can be instant.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"context"
	"hathoracloud"
	"hathoracloud/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := hathoracloud.New(
        hathoracloud.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
        hathoracloud.WithOrgID("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"),
        hathoracloud.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
    )

    res, err := s.ProcessesV2.CreateProcessV2Deprecated(ctx, components.RegionTokyo, hathoracloud.String("app-af469a92-5b45-4565-b3c4-b79878de67d2"))
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
| `region`                                                 | [components.Region](../../models/components/region.md)   | :heavy_check_mark:                                       | N/A                                                      |                                                          |
| `appID`                                                  | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | app-af469a92-5b45-4565-b3c4-b79878de67d2                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*components.ProcessV2](../../models/components/processv2.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.APIError              | 401, 402, 404, 422, 429, 500 | application/json             |
| errors.SDKError              | 4XX, 5XX                     | \*/\*                        |