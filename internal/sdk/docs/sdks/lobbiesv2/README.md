# LobbiesV2
(*LobbiesV2*)

## Overview

### Available Operations

* [~~CreatePrivateLobby~~](#createprivatelobby) - CreatePrivateLobby :warning: **Deprecated**
* [~~CreatePublicLobby~~](#createpubliclobby) - CreatePublicLobby :warning: **Deprecated**
* [~~CreateLocalLobby~~](#createlocallobby) - CreateLocalLobby :warning: **Deprecated**
* [~~CreateLobbyDeprecated~~](#createlobbydeprecated) - CreateLobbyDeprecated :warning: **Deprecated**
* [~~ListActivePublicLobbiesDeprecatedV2~~](#listactivepubliclobbiesdeprecatedv2) - ListActivePublicLobbiesDeprecatedV2 :warning: **Deprecated**
* [~~GetLobbyInfo~~](#getlobbyinfo) - GetLobbyInfo :warning: **Deprecated**
* [~~SetLobbyState~~](#setlobbystate) - SetLobbyState :warning: **Deprecated**

## ~~CreatePrivateLobby~~

CreatePrivateLobby

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

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

    res, err := s.LobbiesV2.CreatePrivateLobby(ctx, operations.CreatePrivateLobbySecurity{
        PlayerAuth: "<YOUR_BEARER_TOKEN_HERE>",
    }, operations.CreatePrivateLobbyRequestBody{
        InitialConfig: "<value>",
        Region: components.RegionLondon,
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          | Example                                                                                              |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |                                                                                                      |
| `security`                                                                                           | [operations.CreatePrivateLobbySecurity](../../models/operations/createprivatelobbysecurity.md)       | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |                                                                                                      |
| `requestBody`                                                                                        | [operations.CreatePrivateLobbyRequestBody](../../models/operations/createprivatelobbyrequestbody.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |                                                                                                      |
| `appID`                                                                                              | **string*                                                                                            | :heavy_minus_sign:                                                                                   | N/A                                                                                                  | app-af469a92-5b45-4565-b3c4-b79878de67d2                                                             |
| `roomID`                                                                                             | **string*                                                                                            | :heavy_minus_sign:                                                                                   | N/A                                                                                                  | 2swovpy1fnunu                                                                                        |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |                                                                                                      |

### Response

**[*components.Lobby](../../models/components/lobby.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| errors.APIError                   | 400, 401, 402, 404, 422, 429, 500 | application/json                  |
| errors.SDKError                   | 4XX, 5XX                          | \*/\*                             |

## ~~CreatePublicLobby~~

CreatePublicLobby

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

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

    res, err := s.LobbiesV2.CreatePublicLobby(ctx, operations.CreatePublicLobbySecurity{
        PlayerAuth: "<YOUR_BEARER_TOKEN_HERE>",
    }, operations.CreatePublicLobbyRequestBody{
        InitialConfig: "<value>",
        Region: components.RegionJohannesburg,
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        | Example                                                                                            |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |                                                                                                    |
| `security`                                                                                         | [operations.CreatePublicLobbySecurity](../../models/operations/createpubliclobbysecurity.md)       | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |                                                                                                    |
| `requestBody`                                                                                      | [operations.CreatePublicLobbyRequestBody](../../models/operations/createpubliclobbyrequestbody.md) | :heavy_check_mark:                                                                                 | N/A                                                                                                |                                                                                                    |
| `appID`                                                                                            | **string*                                                                                          | :heavy_minus_sign:                                                                                 | N/A                                                                                                | app-af469a92-5b45-4565-b3c4-b79878de67d2                                                           |
| `roomID`                                                                                           | **string*                                                                                          | :heavy_minus_sign:                                                                                 | N/A                                                                                                | 2swovpy1fnunu                                                                                      |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |                                                                                                    |

### Response

**[*components.Lobby](../../models/components/lobby.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| errors.APIError                   | 400, 401, 402, 404, 422, 429, 500 | application/json                  |
| errors.SDKError                   | 4XX, 5XX                          | \*/\*                             |

## ~~CreateLocalLobby~~

CreateLocalLobby

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

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

    res, err := s.LobbiesV2.CreateLocalLobby(ctx, operations.CreateLocalLobbySecurity{
        PlayerAuth: "<YOUR_BEARER_TOKEN_HERE>",
    }, operations.CreateLocalLobbyRequestBody{
        InitialConfig: "<value>",
        Region: components.RegionJohannesburg,
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      | Example                                                                                          |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |                                                                                                  |
| `security`                                                                                       | [operations.CreateLocalLobbySecurity](../../models/operations/createlocallobbysecurity.md)       | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |                                                                                                  |
| `requestBody`                                                                                    | [operations.CreateLocalLobbyRequestBody](../../models/operations/createlocallobbyrequestbody.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |                                                                                                  |
| `appID`                                                                                          | **string*                                                                                        | :heavy_minus_sign:                                                                               | N/A                                                                                              | app-af469a92-5b45-4565-b3c4-b79878de67d2                                                         |
| `roomID`                                                                                         | **string*                                                                                        | :heavy_minus_sign:                                                                               | N/A                                                                                              | 2swovpy1fnunu                                                                                    |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |                                                                                                  |

### Response

**[*components.Lobby](../../models/components/lobby.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| errors.APIError                   | 400, 401, 402, 404, 422, 429, 500 | application/json                  |
| errors.SDKError                   | 4XX, 5XX                          | \*/\*                             |

## ~~CreateLobbyDeprecated~~

Create a new lobby for an [application](https://hathora.dev/docs/concepts/hathora-entities#application). A lobby object is a wrapper around a [room](https://hathora.dev/docs/concepts/hathora-entities#room) object. With a lobby, you get additional functionality like configuring the visibility of the room, managing the state of a match, and retrieving a list of public lobbies to display to players.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

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

    res, err := s.LobbiesV2.CreateLobbyDeprecated(ctx, operations.CreateLobbyDeprecatedSecurity{
        PlayerAuth: "<YOUR_BEARER_TOKEN_HERE>",
    }, components.CreateLobbyParams{
        Visibility: components.LobbyVisibilityPrivate,
        InitialConfig: "<value>",
        Region: components.RegionSaoPaulo,
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          | Example                                                                                              |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |                                                                                                      |
| `security`                                                                                           | [operations.CreateLobbyDeprecatedSecurity](../../models/operations/createlobbydeprecatedsecurity.md) | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |                                                                                                      |
| `createLobbyParams`                                                                                  | [components.CreateLobbyParams](../../models/components/createlobbyparams.md)                         | :heavy_check_mark:                                                                                   | N/A                                                                                                  |                                                                                                      |
| `appID`                                                                                              | **string*                                                                                            | :heavy_minus_sign:                                                                                   | N/A                                                                                                  | app-af469a92-5b45-4565-b3c4-b79878de67d2                                                             |
| `roomID`                                                                                             | **string*                                                                                            | :heavy_minus_sign:                                                                                   | N/A                                                                                                  | 2swovpy1fnunu                                                                                        |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |                                                                                                      |

### Response

**[*components.Lobby](../../models/components/lobby.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| errors.APIError                   | 400, 401, 402, 404, 422, 429, 500 | application/json                  |
| errors.SDKError                   | 4XX, 5XX                          | \*/\*                             |

## ~~ListActivePublicLobbiesDeprecatedV2~~

Get all active lobbies for a an [application](https://hathora.dev/docs/concepts/hathora-entities#application). Filter by optionally passing in a `region`. Use this endpoint to display all public lobbies that a player can join in the game client.

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
        hathoracloud.WithOrgID("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"),
        hathoracloud.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
    )

    res, err := s.LobbiesV2.ListActivePublicLobbiesDeprecatedV2(ctx, hathoracloud.String("app-af469a92-5b45-4565-b3c4-b79878de67d2"), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                               | Type                                                                                    | Required                                                                                | Description                                                                             | Example                                                                                 |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `ctx`                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                   | :heavy_check_mark:                                                                      | The context to use for the request.                                                     |                                                                                         |
| `appID`                                                                                 | **string*                                                                               | :heavy_minus_sign:                                                                      | N/A                                                                                     | app-af469a92-5b45-4565-b3c4-b79878de67d2                                                |
| `region`                                                                                | [*components.Region](../../models/components/region.md)                                 | :heavy_minus_sign:                                                                      | Region to filter by. If omitted, active public lobbies in all regions will be returned. |                                                                                         |
| `opts`                                                                                  | [][operations.Option](../../models/operations/option.md)                                | :heavy_minus_sign:                                                                      | The options for this request.                                                           |                                                                                         |

### Response

**[[]components.Lobby](../../.md), error**

### Errors

| Error Type       | Status Code      | Content Type     |
| ---------------- | ---------------- | ---------------- |
| errors.APIError  | 401, 429         | application/json |
| errors.SDKError  | 4XX, 5XX         | \*/\*            |

## ~~GetLobbyInfo~~

Get details for a lobby.

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
        hathoracloud.WithOrgID("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"),
        hathoracloud.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
    )

    res, err := s.LobbiesV2.GetLobbyInfo(ctx, "2swovpy1fnunu", hathoracloud.String("app-af469a92-5b45-4565-b3c4-b79878de67d2"))
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

**[*components.Lobby](../../models/components/lobby.md), error**

### Errors

| Error Type       | Status Code      | Content Type     |
| ---------------- | ---------------- | ---------------- |
| errors.APIError  | 404, 429         | application/json |
| errors.SDKError  | 4XX, 5XX         | \*/\*            |

## ~~SetLobbyState~~

Set the state of a lobby. State is intended to be set by the server and must be smaller than 1MB. Use this endpoint to store match data like live player count to enforce max number of clients or persist end-game data (i.e. winner or final scores).

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

    res, err := s.LobbiesV2.SetLobbyState(ctx, "2swovpy1fnunu", components.SetLobbyStateParams{
        State: "South Dakota",
    }, hathoracloud.String("app-af469a92-5b45-4565-b3c4-b79878de67d2"))
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
| `roomID`                                                                         | *string*                                                                         | :heavy_check_mark:                                                               | N/A                                                                              | 2swovpy1fnunu                                                                    |
| `setLobbyStateParams`                                                            | [components.SetLobbyStateParams](../../models/components/setlobbystateparams.md) | :heavy_check_mark:                                                               | N/A                                                                              |                                                                                  |
| `appID`                                                                          | **string*                                                                        | :heavy_minus_sign:                                                               | N/A                                                                              | app-af469a92-5b45-4565-b3c4-b79878de67d2                                         |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |                                                                                  |

### Response

**[*components.Lobby](../../models/components/lobby.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| errors.APIError    | 401, 404, 422, 429 | application/json   |
| errors.SDKError    | 4XX, 5XX           | \*/\*              |