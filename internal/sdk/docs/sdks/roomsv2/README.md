# RoomsV2
(*RoomsV2*)

## Overview

### Available Operations

* [CreateRoom](#createroom) - CreateRoom
* [GetRoomInfo](#getroominfo) - GetRoomInfo
* [GetActiveRoomsForProcess](#getactiveroomsforprocess) - GetActiveRoomsForProcess
* [GetInactiveRoomsForProcess](#getinactiveroomsforprocess) - GetInactiveRoomsForProcess
* [DestroyRoom](#destroyroom) - DestroyRoom
* [~~SuspendRoomV2Deprecated~~](#suspendroomv2deprecated) - SuspendRoomV2Deprecated :warning: **Deprecated**
* [GetConnectionInfo](#getconnectioninfo) - GetConnectionInfo
* [UpdateRoomConfig](#updateroomconfig) - UpdateRoomConfig

## CreateRoom

Create a new [room](https://hathora.dev/docs/concepts/hathora-entities#room) for an existing [application](https://hathora.dev/docs/concepts/hathora-entities#application). Poll the [`GetConnectionInfo()`](https://hathora.dev/api#tag/RoomV2/operation/GetConnectionInfo) endpoint to get connection details for an active room.

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

    res, err := s.RoomsV2.CreateRoom(ctx, components.CreateRoomParams{
        DeploymentID: hathoracloud.String("dep-6d4c6a71-2d75-4b42-94e1-f312f57f33c5"),
        ClientIPs: []string{
            "123.123.123.123",
        },
        RoomConfig: hathoracloud.String("{\"name\":\"my-room\"}"),
        Region: components.RegionDallas,
    }, hathoracloud.String("app-af469a92-5b45-4565-b3c4-b79878de67d2"), hathoracloud.String("2swovpy1fnunu"))
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                | Example                                                                    |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |                                                                            |
| `createRoomParams`                                                         | [components.CreateRoomParams](../../models/components/createroomparams.md) | :heavy_check_mark:                                                         | N/A                                                                        |                                                                            |
| `appID`                                                                    | **string*                                                                  | :heavy_minus_sign:                                                         | N/A                                                                        | app-af469a92-5b45-4565-b3c4-b79878de67d2                                   |
| `roomID`                                                                   | **string*                                                                  | :heavy_minus_sign:                                                         | N/A                                                                        | 2swovpy1fnunu                                                              |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |                                                                            |

### Response

**[*components.RoomConnectionData](../../models/components/roomconnectiondata.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| errors.APIError                   | 400, 401, 402, 404, 422, 429, 500 | application/json                  |
| errors.SDKError                   | 4XX, 5XX                          | \*/\*                             |

## GetRoomInfo

Retreive current and historical allocation data for a [room](https://hathora.dev/docs/concepts/hathora-entities#room).

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

    res, err := s.RoomsV2.GetRoomInfo(ctx, "2swovpy1fnunu", hathoracloud.String("app-af469a92-5b45-4565-b3c4-b79878de67d2"))
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
| `roomID`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | 2swovpy1fnunu                                            |
| `appID`                                                  | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | app-af469a92-5b45-4565-b3c4-b79878de67d2                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*components.Room](../../models/components/room.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| errors.APIError    | 401, 404, 422, 429 | application/json   |
| errors.SDKError    | 4XX, 5XX           | \*/\*              |

## GetActiveRoomsForProcess

Get all active [rooms](https://hathora.dev/docs/concepts/hathora-entities#room) for a given [process](https://hathora.dev/docs/concepts/hathora-entities#process).

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

    res, err := s.RoomsV2.GetActiveRoomsForProcess(ctx, "cbfcddd2-0006-43ae-996c-995fff7bed2e", hathoracloud.String("app-af469a92-5b45-4565-b3c4-b79878de67d2"))
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

**[[]components.RoomWithoutAllocations](../../.md), error**

### Errors

| Error Type       | Status Code      | Content Type     |
| ---------------- | ---------------- | ---------------- |
| errors.APIError  | 401, 404, 429    | application/json |
| errors.SDKError  | 4XX, 5XX         | \*/\*            |

## GetInactiveRoomsForProcess

Get all inactive [rooms](https://hathora.dev/docs/concepts/hathora-entities#room) for a given [process](https://hathora.dev/docs/concepts/hathora-entities#process).

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

    res, err := s.RoomsV2.GetInactiveRoomsForProcess(ctx, "cbfcddd2-0006-43ae-996c-995fff7bed2e", hathoracloud.String("app-af469a92-5b45-4565-b3c4-b79878de67d2"))
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

**[[]components.RoomWithoutAllocations](../../.md), error**

### Errors

| Error Type       | Status Code      | Content Type     |
| ---------------- | ---------------- | ---------------- |
| errors.APIError  | 401, 404, 429    | application/json |
| errors.SDKError  | 4XX, 5XX         | \*/\*            |

## DestroyRoom

Destroy a [room](https://hathora.dev/docs/concepts/hathora-entities#room). All associated metadata is deleted.

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

    err := s.RoomsV2.DestroyRoom(ctx, "2swovpy1fnunu", hathoracloud.String("app-af469a92-5b45-4565-b3c4-b79878de67d2"))
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `roomID`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | 2swovpy1fnunu                                            |
| `appID`                                                  | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | app-af469a92-5b45-4565-b3c4-b79878de67d2                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| errors.APIError    | 401, 404, 429, 500 | application/json   |
| errors.SDKError    | 4XX, 5XX           | \*/\*              |

## ~~SuspendRoomV2Deprecated~~

Suspend a [room](https://hathora.dev/docs/concepts/hathora-entities#room). The room is unallocated from the process but can be rescheduled later using the same `roomId`.

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

    err := s.RoomsV2.SuspendRoomV2Deprecated(ctx, "2swovpy1fnunu", hathoracloud.String("app-af469a92-5b45-4565-b3c4-b79878de67d2"))
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `roomID`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | 2swovpy1fnunu                                            |
| `appID`                                                  | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | app-af469a92-5b45-4565-b3c4-b79878de67d2                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| errors.APIError    | 401, 404, 429, 500 | application/json   |
| errors.SDKError    | 4XX, 5XX           | \*/\*              |

## GetConnectionInfo

Poll this endpoint to get connection details to a [room](https://hathora.dev/docs/concepts/hathora-entities#room). Clients can call this endpoint without authentication.

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
        hathoracloud.WithOrgID("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"),
        hathoracloud.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
    )

    res, err := s.RoomsV2.GetConnectionInfo(ctx, "2swovpy1fnunu", hathoracloud.String("app-af469a92-5b45-4565-b3c4-b79878de67d2"))
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
| `roomID`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | 2swovpy1fnunu                                            |
| `appID`                                                  | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | app-af469a92-5b45-4565-b3c4-b79878de67d2                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*components.ConnectionInfoV2](../../models/components/connectioninfov2.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.APIError              | 400, 402, 404, 422, 429, 500 | application/json             |
| errors.SDKError              | 4XX, 5XX                     | \*/\*                        |

## UpdateRoomConfig

UpdateRoomConfig

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

    err := s.RoomsV2.UpdateRoomConfig(ctx, "2swovpy1fnunu", components.UpdateRoomConfigParams{
        RoomConfig: "{\"name\":\"my-room\"}",
    }, hathoracloud.String("app-af469a92-5b45-4565-b3c4-b79878de67d2"))
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            | Example                                                                                |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |                                                                                        |
| `roomID`                                                                               | *string*                                                                               | :heavy_check_mark:                                                                     | N/A                                                                                    | 2swovpy1fnunu                                                                          |
| `updateRoomConfigParams`                                                               | [components.UpdateRoomConfigParams](../../models/components/updateroomconfigparams.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |                                                                                        |
| `appID`                                                                                | **string*                                                                              | :heavy_minus_sign:                                                                     | N/A                                                                                    | app-af469a92-5b45-4565-b3c4-b79878de67d2                                               |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |                                                                                        |

### Response

**error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| errors.APIError         | 401, 404, 422, 429, 500 | application/json        |
| errors.SDKError         | 4XX, 5XX                | \*/\*                   |