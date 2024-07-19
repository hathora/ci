# LobbiesV3
(*LobbiesV3*)

### Available Operations

* [CreateLobby](#createlobby) - Create a new lobby for an [application](https://hathora.dev/docs/concepts/hathora-entities#application). A lobby object is a wrapper around a [room](https://hathora.dev/docs/concepts/hathora-entities#room) object. With a lobby, you get additional functionality like configuring the visibility of the room, managing the state of a match, and retrieving a list of public lobbies to display to players.
* [ListActivePublicLobbies](#listactivepubliclobbies) - Get all active lobbies for a given [application](https://hathora.dev/docs/concepts/hathora-entities#application). Filter the array by optionally passing in a `region`. Use this endpoint to display all public lobbies that a player can join in the game client.
* [GetLobbyInfoByRoomID](#getlobbyinfobyroomid) - Get details for a lobby.
* [GetLobbyInfoByShortCode](#getlobbyinfobyshortcode) - Get details for a lobby. If 2 or more lobbies have the same `shortCode`, then the most recently created lobby will be returned.

## CreateLobby

Create a new lobby for an [application](https://hathora.dev/docs/concepts/hathora-entities#application). A lobby object is a wrapper around a [room](https://hathora.dev/docs/concepts/hathora-entities#room) object. With a lobby, you get additional functionality like configuring the visibility of the room, managing the state of a match, and retrieving a list of public lobbies to display to players.

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

    security := operations.CreateLobbySecurity{
            PlayerAuth: "<YOUR_BEARER_TOKEN_HERE>",
        }

    createLobbyV3Params := shared.CreateLobbyV3Params{
        Visibility: shared.LobbyVisibilityPrivate,
        RoomConfig: sdk.String("{\"name\":\"my-room\"}"),
        Region: shared.RegionSeattle,
    }

    var appID *string = sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")

    var shortCode *string = sdk.String("LFG4")

    var roomID *string = sdk.String("2swovpy1fnunu")
    
    ctx := context.Background()
    res, err := s.LobbiesV3.CreateLobby(ctx, security, createLobbyV3Params, appID, shortCode, roomID)
    if err != nil {
        log.Fatal(err)
    }
    if res.LobbyV3 != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |                                                                                  |
| `security`                                                                       | [operations.CreateLobbySecurity](../../models/operations/createlobbysecurity.md) | :heavy_check_mark:                                                               | The security requirements to use for the request.                                |                                                                                  |
| `createLobbyV3Params`                                                            | [shared.CreateLobbyV3Params](../../models/shared/createlobbyv3params.md)         | :heavy_check_mark:                                                               | N/A                                                                              |                                                                                  |
| `appID`                                                                          | **string*                                                                        | :heavy_minus_sign:                                                               | N/A                                                                              | app-af469a92-5b45-4565-b3c4-b79878de67d2                                         |
| `shortCode`                                                                      | **string*                                                                        | :heavy_minus_sign:                                                               | N/A                                                                              | LFG4                                                                             |
| `roomID`                                                                         | **string*                                                                        | :heavy_minus_sign:                                                               | N/A                                                                              | 2swovpy1fnunu                                                                    |


### Response

**[*operations.CreateLobbyResponse](../../models/operations/createlobbyresponse.md), error**
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.APIError          | 400,401,402,404,422,429,500 | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |

## ListActivePublicLobbies

Get all active lobbies for a given [application](https://hathora.dev/docs/concepts/hathora-entities#application). Filter the array by optionally passing in a `region`. Use this endpoint to display all public lobbies that a player can join in the game client.

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

    var region *shared.Region = shared.RegionSeattle.ToPointer()
    
    ctx := context.Background()
    res, err := s.LobbiesV3.ListActivePublicLobbies(ctx, appID, region)
    if err != nil {
        log.Fatal(err)
    }
    if res.LobbyV3s != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        | Example                                                            |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |                                                                    |
| `appID`                                                            | **string*                                                          | :heavy_minus_sign:                                                 | N/A                                                                | app-af469a92-5b45-4565-b3c4-b79878de67d2                           |
| `region`                                                           | [*shared.Region](../../models/shared/region.md)                    | :heavy_minus_sign:                                                 | If omitted, active public lobbies in all regions will be returned. |                                                                    |


### Response

**[*operations.ListActivePublicLobbiesResponse](../../models/operations/listactivepubliclobbiesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,429            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetLobbyInfoByRoomID

Get details for a lobby.

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
    res, err := s.LobbiesV3.GetLobbyInfoByRoomID(ctx, roomID, appID)
    if err != nil {
        log.Fatal(err)
    }
    if res.LobbyV3 != nil {
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

**[*operations.GetLobbyInfoByRoomIDResponse](../../models/operations/getlobbyinfobyroomidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 404,422,429        | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetLobbyInfoByShortCode

Get details for a lobby. If 2 or more lobbies have the same `shortCode`, then the most recently created lobby will be returned.

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

    var shortCode string = "LFG4"

    var appID *string = sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")
    
    ctx := context.Background()
    res, err := s.LobbiesV3.GetLobbyInfoByShortCode(ctx, shortCode, appID)
    if err != nil {
        log.Fatal(err)
    }
    if res.LobbyV3 != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `shortCode`                                           | *string*                                              | :heavy_check_mark:                                    | N/A                                                   | LFG4                                                  |
| `appID`                                               | **string*                                             | :heavy_minus_sign:                                    | N/A                                                   | app-af469a92-5b45-4565-b3c4-b79878de67d2              |


### Response

**[*operations.GetLobbyInfoByShortCodeResponse](../../models/operations/getlobbyinfobyshortcoderesponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 404,429            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
