# LobbiesV3
(*LobbiesV3*)

## Overview

### Available Operations

* [CreateLobby](#createlobby) - CreateLobby
* [ListActivePublicLobbies](#listactivepubliclobbies) - ListActivePublicLobbies
* [GetLobbyInfoByRoomID](#getlobbyinfobyroomid) - GetLobbyInfoByRoomId
* [GetLobbyInfoByShortCode](#getlobbyinfobyshortcode) - GetLobbyInfoByShortCode

## CreateLobby

Create a new lobby for an [application](https://hathora.dev/docs/concepts/hathora-entities#application). A lobby object is a wrapper around a [room](https://hathora.dev/docs/concepts/hathora-entities#room) object. With a lobby, you get additional functionality like configuring the visibility of the room, managing the state of a match, and retrieving a list of public lobbies to display to players.

### Example Usage

```go
package main

import(
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
        Region: components.RegionSeattle,
    }, hathoracloud.String("app-af469a92-5b45-4565-b3c4-b79878de67d2"), hathoracloud.String("LFG4"), hathoracloud.String("2swovpy1fnunu"))
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |                                                                                  |
| `security`                                                                       | [operations.CreateLobbySecurity](../../models/operations/createlobbysecurity.md) | :heavy_check_mark:                                                               | The security requirements to use for the request.                                |                                                                                  |
| `createLobbyV3Params`                                                            | [components.CreateLobbyV3Params](../../models/components/createlobbyv3params.md) | :heavy_check_mark:                                                               | N/A                                                                              |                                                                                  |
| `appID`                                                                          | **string*                                                                        | :heavy_minus_sign:                                                               | N/A                                                                              | app-af469a92-5b45-4565-b3c4-b79878de67d2                                         |
| `shortCode`                                                                      | **string*                                                                        | :heavy_minus_sign:                                                               | N/A                                                                              | LFG4                                                                             |
| `roomID`                                                                         | **string*                                                                        | :heavy_minus_sign:                                                               | N/A                                                                              | 2swovpy1fnunu                                                                    |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |                                                                                  |

### Response

**[*components.LobbyV3](../../models/components/lobbyv3.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| errors.APIError                   | 400, 401, 402, 404, 422, 429, 500 | application/json                  |
| errors.SDKError                   | 4XX, 5XX                          | \*/\*                             |

## ListActivePublicLobbies

Get all active lobbies for a given [application](https://hathora.dev/docs/concepts/hathora-entities#application). Filter the array by optionally passing in a `region`. Use this endpoint to display all public lobbies that a player can join in the game client.

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

    res, err := s.LobbiesV3.ListActivePublicLobbies(ctx, hathoracloud.String("app-af469a92-5b45-4565-b3c4-b79878de67d2"), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        | Example                                                            |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |                                                                    |
| `appID`                                                            | **string*                                                          | :heavy_minus_sign:                                                 | N/A                                                                | app-af469a92-5b45-4565-b3c4-b79878de67d2                           |
| `region`                                                           | [*components.Region](../../models/components/region.md)            | :heavy_minus_sign:                                                 | If omitted, active public lobbies in all regions will be returned. |                                                                    |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |                                                                    |

### Response

**[[]components.LobbyV3](../../.md), error**

### Errors

| Error Type       | Status Code      | Content Type     |
| ---------------- | ---------------- | ---------------- |
| errors.APIError  | 401, 429         | application/json |
| errors.SDKError  | 4XX, 5XX         | \*/\*            |

## GetLobbyInfoByRoomID

Get details for a lobby.

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

    res, err := s.LobbiesV3.GetLobbyInfoByRoomID(ctx, "2swovpy1fnunu", hathoracloud.String("app-af469a92-5b45-4565-b3c4-b79878de67d2"))
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

**[*components.LobbyV3](../../models/components/lobbyv3.md), error**

### Errors

| Error Type       | Status Code      | Content Type     |
| ---------------- | ---------------- | ---------------- |
| errors.APIError  | 404, 422, 429    | application/json |
| errors.SDKError  | 4XX, 5XX         | \*/\*            |

## GetLobbyInfoByShortCode

Get details for a lobby. If 2 or more lobbies have the same `shortCode`, then the most recently created lobby will be returned.

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

    res, err := s.LobbiesV3.GetLobbyInfoByShortCode(ctx, "LFG4", hathoracloud.String("app-af469a92-5b45-4565-b3c4-b79878de67d2"))
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
| `shortCode`                                              | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | LFG4                                                     |
| `appID`                                                  | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | app-af469a92-5b45-4565-b3c4-b79878de67d2                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*components.LobbyV3](../../models/components/lobbyv3.md), error**

### Errors

| Error Type       | Status Code      | Content Type     |
| ---------------- | ---------------- | ---------------- |
| errors.APIError  | 404, 429         | application/json |
| errors.SDKError  | 4XX, 5XX         | \*/\*            |