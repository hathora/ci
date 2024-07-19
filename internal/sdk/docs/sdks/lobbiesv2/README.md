# LobbiesV2
(*LobbiesV2*)

### Available Operations

* [~~CreatePrivateLobby~~](#createprivatelobby) - :warning: **Deprecated**
* [~~CreatePublicLobby~~](#createpubliclobby) - :warning: **Deprecated**
* [~~CreateLocalLobby~~](#createlocallobby) - :warning: **Deprecated**
* [~~CreateLobbyDeprecated~~](#createlobbydeprecated) - Create a new lobby for an [application](https://hathora.dev/docs/concepts/hathora-entities#application). A lobby object is a wrapper around a [room](https://hathora.dev/docs/concepts/hathora-entities#room) object. With a lobby, you get additional functionality like configuring the visibility of the room, managing the state of a match, and retrieving a list of public lobbies to display to players. :warning: **Deprecated**
* [~~ListActivePublicLobbiesDeprecatedV2~~](#listactivepubliclobbiesdeprecatedv2) - Get all active lobbies for a an [application](https://hathora.dev/docs/concepts/hathora-entities#application). Filter by optionally passing in a `region`. Use this endpoint to display all public lobbies that a player can join in the game client. :warning: **Deprecated**
* [~~GetLobbyInfo~~](#getlobbyinfo) - Get details for a lobby. :warning: **Deprecated**
* [~~SetLobbyState~~](#setlobbystate) - Set the state of a lobby. State is intended to be set by the server and must be smaller than 1MB. Use this endpoint to store match data like live player count to enforce max number of clients or persist end-game data (i.e. winner or final scores). :warning: **Deprecated**

## ~~CreatePrivateLobby~~

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"github.com/hathora/ci/internal/sdk"
	"github.com/hathora/ci/internal/sdk/models/operations"
	"github.com/hathora/ci/internal/sdk/models/shared"
	"context"
	"log"
)

func main() {
    s := sdk.New(
        sdk.WithAppID(sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")),
    )

    security := operations.CreatePrivateLobbySecurity{
            PlayerAuth: "<YOUR_BEARER_TOKEN_HERE>",
        }

    requestBody := operations.CreatePrivateLobbyRequestBody{
        InitialConfig: "<value>",
        Region: shared.RegionChicago,
    }

    var appID *string = sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")

    var roomID *string = sdk.String("2swovpy1fnunu")
    
    ctx := context.Background()
    res, err := s.LobbiesV2.CreatePrivateLobby(ctx, security, requestBody, appID, roomID)
    if err != nil {
        log.Fatal(err)
    }
    if res.Lobby != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          | Example                                                                                              |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |                                                                                                      |
| `security`                                                                                           | [operations.CreatePrivateLobbySecurity](../../models/operations/createprivatelobbysecurity.md)       | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |                                                                                                      |
| `requestBody`                                                                                        | [operations.CreatePrivateLobbyRequestBody](../../models/operations/createprivatelobbyrequestbody.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |                                                                                                      |
| `appID`                                                                                              | **string*                                                                                            | :heavy_minus_sign:                                                                                   | N/A                                                                                                  | app-af469a92-5b45-4565-b3c4-b79878de67d2                                                             |
| `roomID`                                                                                             | **string*                                                                                            | :heavy_minus_sign:                                                                                   | N/A                                                                                                  | 2swovpy1fnunu                                                                                        |


### Response

**[*operations.CreatePrivateLobbyResponse](../../models/operations/createprivatelobbyresponse.md), error**
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.APIError          | 400,401,402,404,422,429,500 | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |

## ~~CreatePublicLobby~~

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"github.com/hathora/ci/internal/sdk"
	"github.com/hathora/ci/internal/sdk/models/operations"
	"github.com/hathora/ci/internal/sdk/models/shared"
	"context"
	"log"
)

func main() {
    s := sdk.New(
        sdk.WithAppID(sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")),
    )

    security := operations.CreatePublicLobbySecurity{
            PlayerAuth: "<YOUR_BEARER_TOKEN_HERE>",
        }

    requestBody := operations.CreatePublicLobbyRequestBody{
        InitialConfig: "<value>",
        Region: shared.RegionSaoPaulo,
    }

    var appID *string = sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")

    var roomID *string = sdk.String("2swovpy1fnunu")
    
    ctx := context.Background()
    res, err := s.LobbiesV2.CreatePublicLobby(ctx, security, requestBody, appID, roomID)
    if err != nil {
        log.Fatal(err)
    }
    if res.Lobby != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        | Example                                                                                            |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |                                                                                                    |
| `security`                                                                                         | [operations.CreatePublicLobbySecurity](../../models/operations/createpubliclobbysecurity.md)       | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |                                                                                                    |
| `requestBody`                                                                                      | [operations.CreatePublicLobbyRequestBody](../../models/operations/createpubliclobbyrequestbody.md) | :heavy_check_mark:                                                                                 | N/A                                                                                                |                                                                                                    |
| `appID`                                                                                            | **string*                                                                                          | :heavy_minus_sign:                                                                                 | N/A                                                                                                | app-af469a92-5b45-4565-b3c4-b79878de67d2                                                           |
| `roomID`                                                                                           | **string*                                                                                          | :heavy_minus_sign:                                                                                 | N/A                                                                                                | 2swovpy1fnunu                                                                                      |


### Response

**[*operations.CreatePublicLobbyResponse](../../models/operations/createpubliclobbyresponse.md), error**
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.APIError          | 400,401,402,404,422,429,500 | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |

## ~~CreateLocalLobby~~

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"github.com/hathora/ci/internal/sdk"
	"github.com/hathora/ci/internal/sdk/models/operations"
	"github.com/hathora/ci/internal/sdk/models/shared"
	"context"
	"log"
)

func main() {
    s := sdk.New(
        sdk.WithAppID(sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")),
    )

    security := operations.CreateLocalLobbySecurity{
            PlayerAuth: "<YOUR_BEARER_TOKEN_HERE>",
        }

    requestBody := operations.CreateLocalLobbyRequestBody{
        InitialConfig: "<value>",
        Region: shared.RegionSaoPaulo,
    }

    var appID *string = sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")

    var roomID *string = sdk.String("2swovpy1fnunu")
    
    ctx := context.Background()
    res, err := s.LobbiesV2.CreateLocalLobby(ctx, security, requestBody, appID, roomID)
    if err != nil {
        log.Fatal(err)
    }
    if res.Lobby != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      | Example                                                                                          |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |                                                                                                  |
| `security`                                                                                       | [operations.CreateLocalLobbySecurity](../../models/operations/createlocallobbysecurity.md)       | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |                                                                                                  |
| `requestBody`                                                                                    | [operations.CreateLocalLobbyRequestBody](../../models/operations/createlocallobbyrequestbody.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |                                                                                                  |
| `appID`                                                                                          | **string*                                                                                        | :heavy_minus_sign:                                                                               | N/A                                                                                              | app-af469a92-5b45-4565-b3c4-b79878de67d2                                                         |
| `roomID`                                                                                         | **string*                                                                                        | :heavy_minus_sign:                                                                               | N/A                                                                                              | 2swovpy1fnunu                                                                                    |


### Response

**[*operations.CreateLocalLobbyResponse](../../models/operations/createlocallobbyresponse.md), error**
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.APIError          | 400,401,402,404,422,429,500 | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |

## ~~CreateLobbyDeprecated~~

Create a new lobby for an [application](https://hathora.dev/docs/concepts/hathora-entities#application). A lobby object is a wrapper around a [room](https://hathora.dev/docs/concepts/hathora-entities#room) object. With a lobby, you get additional functionality like configuring the visibility of the room, managing the state of a match, and retrieving a list of public lobbies to display to players.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"github.com/hathora/ci/internal/sdk"
	"github.com/hathora/ci/internal/sdk/models/operations"
	"github.com/hathora/ci/internal/sdk/models/shared"
	"context"
	"log"
)

func main() {
    s := sdk.New(
        sdk.WithAppID(sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")),
    )

    security := operations.CreateLobbyDeprecatedSecurity{
            PlayerAuth: "<YOUR_BEARER_TOKEN_HERE>",
        }

    createLobbyParams := shared.CreateLobbyParams{
        Visibility: shared.LobbyVisibilityPrivate,
        InitialConfig: "<value>",
        Region: shared.RegionTokyo,
    }

    var appID *string = sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")

    var roomID *string = sdk.String("2swovpy1fnunu")
    
    ctx := context.Background()
    res, err := s.LobbiesV2.CreateLobbyDeprecated(ctx, security, createLobbyParams, appID, roomID)
    if err != nil {
        log.Fatal(err)
    }
    if res.Lobby != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          | Example                                                                                              |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |                                                                                                      |
| `security`                                                                                           | [operations.CreateLobbyDeprecatedSecurity](../../models/operations/createlobbydeprecatedsecurity.md) | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |                                                                                                      |
| `createLobbyParams`                                                                                  | [shared.CreateLobbyParams](../../models/shared/createlobbyparams.md)                                 | :heavy_check_mark:                                                                                   | N/A                                                                                                  |                                                                                                      |
| `appID`                                                                                              | **string*                                                                                            | :heavy_minus_sign:                                                                                   | N/A                                                                                                  | app-af469a92-5b45-4565-b3c4-b79878de67d2                                                             |
| `roomID`                                                                                             | **string*                                                                                            | :heavy_minus_sign:                                                                                   | N/A                                                                                                  | 2swovpy1fnunu                                                                                        |


### Response

**[*operations.CreateLobbyDeprecatedResponse](../../models/operations/createlobbydeprecatedresponse.md), error**
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.APIError          | 400,401,402,404,422,429,500 | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |

## ~~ListActivePublicLobbiesDeprecatedV2~~

Get all active lobbies for a an [application](https://hathora.dev/docs/concepts/hathora-entities#application). Filter by optionally passing in a `region`. Use this endpoint to display all public lobbies that a player can join in the game client.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"github.com/hathora/ci/internal/sdk"
	"github.com/hathora/ci/internal/sdk/models/shared"
	"context"
	"log"
)

func main() {
    s := sdk.New(
        sdk.WithAppID(sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")),
    )

    var appID *string = sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")

    var region *shared.Region = shared.RegionFrankfurt.ToPointer()
    
    ctx := context.Background()
    res, err := s.LobbiesV2.ListActivePublicLobbiesDeprecatedV2(ctx, appID, region)
    if err != nil {
        log.Fatal(err)
    }
    if res.Lobbies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                               | Type                                                                                    | Required                                                                                | Description                                                                             | Example                                                                                 |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `ctx`                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                   | :heavy_check_mark:                                                                      | The context to use for the request.                                                     |                                                                                         |
| `appID`                                                                                 | **string*                                                                               | :heavy_minus_sign:                                                                      | N/A                                                                                     | app-af469a92-5b45-4565-b3c4-b79878de67d2                                                |
| `region`                                                                                | [*shared.Region](../../models/shared/region.md)                                         | :heavy_minus_sign:                                                                      | Region to filter by. If omitted, active public lobbies in all regions will be returned. |                                                                                         |


### Response

**[*operations.ListActivePublicLobbiesDeprecatedV2Response](../../models/operations/listactivepubliclobbiesdeprecatedv2response.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,429            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ~~GetLobbyInfo~~

Get details for a lobby.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"github.com/hathora/ci/internal/sdk"
	"context"
	"log"
)

func main() {
    s := sdk.New(
        sdk.WithAppID(sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")),
    )

    var roomID string = "2swovpy1fnunu"

    var appID *string = sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")
    
    ctx := context.Background()
    res, err := s.LobbiesV2.GetLobbyInfo(ctx, roomID, appID)
    if err != nil {
        log.Fatal(err)
    }
    if res.Lobby != nil {
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

**[*operations.GetLobbyInfoResponse](../../models/operations/getlobbyinforesponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 404,429            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ~~SetLobbyState~~

Set the state of a lobby. State is intended to be set by the server and must be smaller than 1MB. Use this endpoint to store match data like live player count to enforce max number of clients or persist end-game data (i.e. winner or final scores).

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"github.com/hathora/ci/internal/sdk/models/shared"
	"github.com/hathora/ci/internal/sdk"
	"context"
	"log"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            HathoraDevToken: sdk.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
        sdk.WithAppID(sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")),
    )

    var roomID string = "2swovpy1fnunu"

    setLobbyStateParams := shared.SetLobbyStateParams{
        State: "<value>",
    }

    var appID *string = sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")
    
    ctx := context.Background()
    res, err := s.LobbiesV2.SetLobbyState(ctx, roomID, setLobbyStateParams, appID)
    if err != nil {
        log.Fatal(err)
    }
    if res.Lobby != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              | Example                                                                  |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |                                                                          |
| `roomID`                                                                 | *string*                                                                 | :heavy_check_mark:                                                       | N/A                                                                      | 2swovpy1fnunu                                                            |
| `setLobbyStateParams`                                                    | [shared.SetLobbyStateParams](../../models/shared/setlobbystateparams.md) | :heavy_check_mark:                                                       | N/A                                                                      |                                                                          |
| `appID`                                                                  | **string*                                                                | :heavy_minus_sign:                                                       | N/A                                                                      | app-af469a92-5b45-4565-b3c4-b79878de67d2                                 |


### Response

**[*operations.SetLobbyStateResponse](../../models/operations/setlobbystateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,404,422,429    | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
