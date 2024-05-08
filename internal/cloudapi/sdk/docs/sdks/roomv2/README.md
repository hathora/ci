# RoomV2
(*RoomV2*)

## Overview

Operations to create, manage, and connect to [rooms](https://hathora.dev/docs/concepts/hathora-entities#room).

### Available Operations

* [CreateRoom](#createroom) - Create a new [room](https://hathora.dev/docs/concepts/hathora-entities#room) for an existing [application](https://hathora.dev/docs/concepts/hathora-entities#application). Poll the [`GetConnectionInfo()`](https://hathora.dev/api#tag/RoomV2/operation/GetConnectionInfo) endpoint to get connection details for an active room.
* [GetRoomInfo](#getroominfo) - Retreive current and historical allocation data for a [room](https://hathora.dev/docs/concepts/hathora-entities#room).
* [GetActiveRoomsForProcess](#getactiveroomsforprocess) - Get all active [rooms](https://hathora.dev/docs/concepts/hathora-entities#room) for a given [process](https://hathora.dev/docs/concepts/hathora-entities#process).
* [GetInactiveRoomsForProcess](#getinactiveroomsforprocess) - Get all inactive [rooms](https://hathora.dev/docs/concepts/hathora-entities#room) for a given [process](https://hathora.dev/docs/concepts/hathora-entities#process).
* [DestroyRoom](#destroyroom) - Destroy a [room](https://hathora.dev/docs/concepts/hathora-entities#room). All associated metadata is deleted.
* [SuspendRoom](#suspendroom) - Suspend a [room](https://hathora.dev/docs/concepts/hathora-entities#room). The room is unallocated from the process but can be rescheduled later using the same `roomId`.
* [GetConnectionInfo](#getconnectioninfo) - Poll this endpoint to get connection details to a [room](https://hathora.dev/docs/concepts/hathora-entities#room). Clients can call this endpoint without authentication.
* [UpdateRoomConfig](#updateroomconfig)

## CreateRoom

Create a new [room](https://hathora.dev/docs/concepts/hathora-entities#room) for an existing [application](https://hathora.dev/docs/concepts/hathora-entities#application). Poll the [`GetConnectionInfo()`](https://hathora.dev/api#tag/RoomV2/operation/GetConnectionInfo) endpoint to get connection details for an active room.

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


    createRoomParams := shared.CreateRoomParams{
        RoomConfig: cloudapi.String("{\"name\":\"my-room\"}"),
        Region: shared.RegionSydney,
    }

    var appID *string = cloudapi.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")

    var roomID *string = cloudapi.String("2swovpy1fnunu")

    ctx := context.Background()
    res, err := s.RoomV2.CreateRoom(ctx, createRoomParams, appID, roomID)
    if err != nil {
        log.Fatal(err)
    }
    if res.RoomConnectionData != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            | Example                                                                |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |                                                                        |
| `createRoomParams`                                                     | [shared.CreateRoomParams](../../pkg/models/shared/createroomparams.md) | :heavy_check_mark:                                                     | N/A                                                                    |                                                                        |
| `appID`                                                                | **string*                                                              | :heavy_minus_sign:                                                     | N/A                                                                    | app-af469a92-5b45-4565-b3c4-b79878de67d2                               |
| `roomID`                                                               | **string*                                                              | :heavy_minus_sign:                                                     | N/A                                                                    | 2swovpy1fnunu                                                          |


### Response

**[*operations.CreateRoomResponse](../../pkg/models/operations/createroomresponse.md), error**
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.APIError          | 400,401,402,403,404,429,500 | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |

## GetRoomInfo

Retreive current and historical allocation data for a [room](https://hathora.dev/docs/concepts/hathora-entities#room).

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


    var roomID string = "2swovpy1fnunu"

    var appID *string = cloudapi.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")

    ctx := context.Background()
    res, err := s.RoomV2.GetRoomInfo(ctx, roomID, appID)
    if err != nil {
        log.Fatal(err)
    }
    if res.Room != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `roomID`                                              | *string*                                              | :heavy_check_mark:                                    | N/A                                                   | 2swovpy1fnunu                                         |
| `appID`                                               | **string*                                             | :heavy_minus_sign:                                    | N/A                                                   | app-af469a92-5b45-4565-b3c4-b79878de67d2              |


### Response

**[*operations.GetRoomInfoResponse](../../pkg/models/operations/getroominforesponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,404            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetActiveRoomsForProcess

Get all active [rooms](https://hathora.dev/docs/concepts/hathora-entities#room) for a given [process](https://hathora.dev/docs/concepts/hathora-entities#process).

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
    res, err := s.RoomV2.GetActiveRoomsForProcess(ctx, processID, appID)
    if err != nil {
        log.Fatal(err)
    }
    if res.RoomWithoutAllocations != nil {
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

**[*operations.GetActiveRoomsForProcessResponse](../../pkg/models/operations/getactiveroomsforprocessresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,404            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetInactiveRoomsForProcess

Get all inactive [rooms](https://hathora.dev/docs/concepts/hathora-entities#room) for a given [process](https://hathora.dev/docs/concepts/hathora-entities#process).

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
    res, err := s.RoomV2.GetInactiveRoomsForProcess(ctx, processID, appID)
    if err != nil {
        log.Fatal(err)
    }
    if res.RoomWithoutAllocations != nil {
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

**[*operations.GetInactiveRoomsForProcessResponse](../../pkg/models/operations/getinactiveroomsforprocessresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,404            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DestroyRoom

Destroy a [room](https://hathora.dev/docs/concepts/hathora-entities#room). All associated metadata is deleted.

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


    var roomID string = "2swovpy1fnunu"

    var appID *string = cloudapi.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")

    ctx := context.Background()
    res, err := s.RoomV2.DestroyRoom(ctx, roomID, appID)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `roomID`                                              | *string*                                              | :heavy_check_mark:                                    | N/A                                                   | 2swovpy1fnunu                                         |
| `appID`                                               | **string*                                             | :heavy_minus_sign:                                    | N/A                                                   | app-af469a92-5b45-4565-b3c4-b79878de67d2              |


### Response

**[*operations.DestroyRoomResponse](../../pkg/models/operations/destroyroomresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,404,429,500    | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## SuspendRoom

Suspend a [room](https://hathora.dev/docs/concepts/hathora-entities#room). The room is unallocated from the process but can be rescheduled later using the same `roomId`.

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


    var roomID string = "2swovpy1fnunu"

    var appID *string = cloudapi.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")

    ctx := context.Background()
    res, err := s.RoomV2.SuspendRoom(ctx, roomID, appID)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `roomID`                                              | *string*                                              | :heavy_check_mark:                                    | N/A                                                   | 2swovpy1fnunu                                         |
| `appID`                                               | **string*                                             | :heavy_minus_sign:                                    | N/A                                                   | app-af469a92-5b45-4565-b3c4-b79878de67d2              |


### Response

**[*operations.SuspendRoomResponse](../../pkg/models/operations/suspendroomresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,404,429,500    | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetConnectionInfo

Poll this endpoint to get connection details to a [room](https://hathora.dev/docs/concepts/hathora-entities#room). Clients can call this endpoint without authentication.

### Example Usage

```go
package main

import(
	"cloudapi"
	"context"
	"log"
)

func main() {
    s := cloudapi.New(
        cloudapi.WithAppID(cloudapi.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")),
    )


    var roomID string = "2swovpy1fnunu"

    var appID *string = cloudapi.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")

    ctx := context.Background()
    res, err := s.RoomV2.GetConnectionInfo(ctx, roomID, appID)
    if err != nil {
        log.Fatal(err)
    }
    if res.ConnectionInfoV2 != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `roomID`                                              | *string*                                              | :heavy_check_mark:                                    | N/A                                                   | 2swovpy1fnunu                                         |
| `appID`                                               | **string*                                             | :heavy_minus_sign:                                    | N/A                                                   | app-af469a92-5b45-4565-b3c4-b79878de67d2              |


### Response

**[*operations.GetConnectionInfoResponse](../../pkg/models/operations/getconnectioninforesponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 400,402,404,500    | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateRoomConfig

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


    var roomID string = "2swovpy1fnunu"

    updateRoomConfigParams := shared.UpdateRoomConfigParams{
        RoomConfig: "{\"name\":\"my-room\"}",
    }

    var appID *string = cloudapi.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")

    ctx := context.Background()
    res, err := s.RoomV2.UpdateRoomConfig(ctx, roomID, updateRoomConfigParams, appID)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `roomID`                                                                           | *string*                                                                           | :heavy_check_mark:                                                                 | N/A                                                                                | 2swovpy1fnunu                                                                      |
| `updateRoomConfigParams`                                                           | [shared.UpdateRoomConfigParams](../../pkg/models/shared/updateroomconfigparams.md) | :heavy_check_mark:                                                                 | N/A                                                                                |                                                                                    |
| `appID`                                                                            | **string*                                                                          | :heavy_minus_sign:                                                                 | N/A                                                                                | app-af469a92-5b45-4565-b3c4-b79878de67d2                                           |


### Response

**[*operations.UpdateRoomConfigResponse](../../pkg/models/operations/updateroomconfigresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,404,429,500    | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
