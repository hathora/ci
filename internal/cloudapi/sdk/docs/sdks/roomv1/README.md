# RoomV1
(*RoomV1*)

## Overview

Deprecated. Use [RoomV2](https://hathora.dev/api#tag/RoomV2).

### Available Operations

* [~~CreateRoomDeprecated~~](#createroomdeprecated) - :warning: **Deprecated**
* [~~GetRoomInfoDeprecated~~](#getroominfodeprecated) - :warning: **Deprecated**
* [~~GetActiveRoomsForProcessDeprecated~~](#getactiveroomsforprocessdeprecated) - :warning: **Deprecated**
* [~~GetInactiveRoomsForProcessDeprecated~~](#getinactiveroomsforprocessdeprecated) - :warning: **Deprecated**
* [~~DestroyRoomDeprecated~~](#destroyroomdeprecated) - :warning: **Deprecated**
* [~~SuspendRoomDeprecated~~](#suspendroomdeprecated) - :warning: **Deprecated**
* [~~GetConnectionInfoDeprecated~~](#getconnectioninfodeprecated) - :warning: **Deprecated**

## ~~CreateRoomDeprecated~~

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


    createRoomParams := shared.CreateRoomParams{
        RoomConfig: cloudapi.String("{\"name\":\"my-room\"}"),
        Region: shared.RegionChicago,
    }

    var appID *string = cloudapi.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")

    var roomID *string = cloudapi.String("2swovpy1fnunu")

    ctx := context.Background()
    res, err := s.RoomV1.CreateRoomDeprecated(ctx, createRoomParams, appID, roomID)
    if err != nil {
        log.Fatal(err)
    }
    if res.RoomID != nil {
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

**[*operations.CreateRoomDeprecatedResponse](../../pkg/models/operations/createroomdeprecatedresponse.md), error**
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.APIError          | 400,401,402,403,404,429,500 | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |

## ~~GetRoomInfoDeprecated~~

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


    var roomID string = "2swovpy1fnunu"

    var appID *string = cloudapi.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")

    ctx := context.Background()
    res, err := s.RoomV1.GetRoomInfoDeprecated(ctx, roomID, appID)
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

**[*operations.GetRoomInfoDeprecatedResponse](../../pkg/models/operations/getroominfodeprecatedresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,404            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ~~GetActiveRoomsForProcessDeprecated~~

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


    var processID string = "cbfcddd2-0006-43ae-996c-995fff7bed2e"

    var appID *string = cloudapi.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")

    ctx := context.Background()
    res, err := s.RoomV1.GetActiveRoomsForProcessDeprecated(ctx, processID, appID)
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

**[*operations.GetActiveRoomsForProcessDeprecatedResponse](../../pkg/models/operations/getactiveroomsforprocessdeprecatedresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,404            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ~~GetInactiveRoomsForProcessDeprecated~~

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


    var processID string = "cbfcddd2-0006-43ae-996c-995fff7bed2e"

    var appID *string = cloudapi.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")

    ctx := context.Background()
    res, err := s.RoomV1.GetInactiveRoomsForProcessDeprecated(ctx, processID, appID)
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

**[*operations.GetInactiveRoomsForProcessDeprecatedResponse](../../pkg/models/operations/getinactiveroomsforprocessdeprecatedresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,404            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ~~DestroyRoomDeprecated~~

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


    var roomID string = "2swovpy1fnunu"

    var appID *string = cloudapi.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")

    ctx := context.Background()
    res, err := s.RoomV1.DestroyRoomDeprecated(ctx, roomID, appID)
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

**[*operations.DestroyRoomDeprecatedResponse](../../pkg/models/operations/destroyroomdeprecatedresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,404,429,500    | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ~~SuspendRoomDeprecated~~

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


    var roomID string = "2swovpy1fnunu"

    var appID *string = cloudapi.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")

    ctx := context.Background()
    res, err := s.RoomV1.SuspendRoomDeprecated(ctx, roomID, appID)
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

**[*operations.SuspendRoomDeprecatedResponse](../../pkg/models/operations/suspendroomdeprecatedresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,404,429,500    | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ~~GetConnectionInfoDeprecated~~

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

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
    res, err := s.RoomV1.GetConnectionInfoDeprecated(ctx, roomID, appID)
    if err != nil {
        log.Fatal(err)
    }
    if res.ConnectionInfo != nil {
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

**[*operations.GetConnectionInfoDeprecatedResponse](../../pkg/models/operations/getconnectioninfodeprecatedresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 400,402,404,500    | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
