# LobbyV2
(*LobbyV2*)

## Overview

Deprecated. Use [LobbyV3](https://hathora.dev/api#tag/LobbyV3).

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
	"cloudapi"
	"cloudapi/pkg/models/shared"
	"cloudapi/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := cloudapi.New(
        cloudapi.WithAppID(cloudapi.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")),
    )


    requestBody := operations.CreatePrivateLobbyRequestBody{
        InitialConfig: shared.LobbyInitialConfig{},
        Region: shared.RegionChicago,
    }

    var appID *string = cloudapi.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")

    var roomID *string = cloudapi.String("2swovpy1fnunu")

    operationSecurity := operations.CreatePrivateLobbySecurity{
            PlayerAuth: "<YOUR_BEARER_TOKEN_HERE>",
        }

    ctx := context.Background()
    res, err := s.LobbyV2.CreatePrivateLobby(ctx, operationSecurity, requestBody, appID, roomID)
    if err != nil {
        log.Fatal(err)
    }
    if res.Lobby != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              | Example                                                                                                  |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |                                                                                                          |
| `security`                                                                                               | [operations.CreatePrivateLobbySecurity](../../pkg/models/operations/createprivatelobbysecurity.md)       | :heavy_check_mark:                                                                                       | The security requirements to use for the request.                                                        |                                                                                                          |
| `requestBody`                                                                                            | [operations.CreatePrivateLobbyRequestBody](../../pkg/models/operations/createprivatelobbyrequestbody.md) | :heavy_check_mark:                                                                                       | N/A                                                                                                      |                                                                                                          |
| `appID`                                                                                                  | **string*                                                                                                | :heavy_minus_sign:                                                                                       | N/A                                                                                                      | app-af469a92-5b45-4565-b3c4-b79878de67d2                                                                 |
| `roomID`                                                                                                 | **string*                                                                                                | :heavy_minus_sign:                                                                                       | N/A                                                                                                      | 2swovpy1fnunu                                                                                            |


### Response

**[*operations.CreatePrivateLobbyResponse](../../pkg/models/operations/createprivatelobbyresponse.md), error**
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
	"cloudapi"
	"cloudapi/pkg/models/shared"
	"cloudapi/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := cloudapi.New(
        cloudapi.WithAppID(cloudapi.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")),
    )


    requestBody := operations.CreatePublicLobbyRequestBody{
        InitialConfig: shared.LobbyInitialConfig{},
        Region: shared.RegionSydney,
    }

    var appID *string = cloudapi.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")

    var roomID *string = cloudapi.String("2swovpy1fnunu")

    operationSecurity := operations.CreatePublicLobbySecurity{
            PlayerAuth: "<YOUR_BEARER_TOKEN_HERE>",
        }

    ctx := context.Background()
    res, err := s.LobbyV2.CreatePublicLobby(ctx, operationSecurity, requestBody, appID, roomID)
    if err != nil {
        log.Fatal(err)
    }
    if res.Lobby != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            | Example                                                                                                |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |                                                                                                        |
| `security`                                                                                             | [operations.CreatePublicLobbySecurity](../../pkg/models/operations/createpubliclobbysecurity.md)       | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |                                                                                                        |
| `requestBody`                                                                                          | [operations.CreatePublicLobbyRequestBody](../../pkg/models/operations/createpubliclobbyrequestbody.md) | :heavy_check_mark:                                                                                     | N/A                                                                                                    |                                                                                                        |
| `appID`                                                                                                | **string*                                                                                              | :heavy_minus_sign:                                                                                     | N/A                                                                                                    | app-af469a92-5b45-4565-b3c4-b79878de67d2                                                               |
| `roomID`                                                                                               | **string*                                                                                              | :heavy_minus_sign:                                                                                     | N/A                                                                                                    | 2swovpy1fnunu                                                                                          |


### Response

**[*operations.CreatePublicLobbyResponse](../../pkg/models/operations/createpubliclobbyresponse.md), error**
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
	"cloudapi"
	"cloudapi/pkg/models/shared"
	"cloudapi/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := cloudapi.New(
        cloudapi.WithAppID(cloudapi.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")),
    )


    requestBody := operations.CreateLocalLobbyRequestBody{
        InitialConfig: shared.LobbyInitialConfig{},
        Region: shared.RegionSydney,
    }

    var appID *string = cloudapi.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")

    var roomID *string = cloudapi.String("2swovpy1fnunu")

    operationSecurity := operations.CreateLocalLobbySecurity{
            PlayerAuth: "<YOUR_BEARER_TOKEN_HERE>",
        }

    ctx := context.Background()
    res, err := s.LobbyV2.CreateLocalLobby(ctx, operationSecurity, requestBody, appID, roomID)
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
| `security`                                                                                           | [operations.CreateLocalLobbySecurity](../../pkg/models/operations/createlocallobbysecurity.md)       | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |                                                                                                      |
| `requestBody`                                                                                        | [operations.CreateLocalLobbyRequestBody](../../pkg/models/operations/createlocallobbyrequestbody.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |                                                                                                      |
| `appID`                                                                                              | **string*                                                                                            | :heavy_minus_sign:                                                                                   | N/A                                                                                                  | app-af469a92-5b45-4565-b3c4-b79878de67d2                                                             |
| `roomID`                                                                                             | **string*                                                                                            | :heavy_minus_sign:                                                                                   | N/A                                                                                                  | 2swovpy1fnunu                                                                                        |


### Response

**[*operations.CreateLocalLobbyResponse](../../pkg/models/operations/createlocallobbyresponse.md), error**
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
	"cloudapi"
	"cloudapi/pkg/models/shared"
	"cloudapi/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := cloudapi.New(
        cloudapi.WithAppID(cloudapi.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")),
    )


    createLobbyParams := shared.CreateLobbyParams{
        Visibility: shared.LobbyVisibilityPrivate,
        InitialConfig: shared.LobbyInitialConfig{},
        Region: shared.RegionTokyo,
    }

    var appID *string = cloudapi.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")

    var roomID *string = cloudapi.String("2swovpy1fnunu")

    operationSecurity := operations.CreateLobbyDeprecatedSecurity{
            PlayerAuth: "<YOUR_BEARER_TOKEN_HERE>",
        }

    ctx := context.Background()
    res, err := s.LobbyV2.CreateLobbyDeprecated(ctx, operationSecurity, createLobbyParams, appID, roomID)
    if err != nil {
        log.Fatal(err)
    }
    if res.Lobby != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              | Example                                                                                                  |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |                                                                                                          |
| `security`                                                                                               | [operations.CreateLobbyDeprecatedSecurity](../../pkg/models/operations/createlobbydeprecatedsecurity.md) | :heavy_check_mark:                                                                                       | The security requirements to use for the request.                                                        |                                                                                                          |
| `createLobbyParams`                                                                                      | [shared.CreateLobbyParams](../../pkg/models/shared/createlobbyparams.md)                                 | :heavy_check_mark:                                                                                       | N/A                                                                                                      |                                                                                                          |
| `appID`                                                                                                  | **string*                                                                                                | :heavy_minus_sign:                                                                                       | N/A                                                                                                      | app-af469a92-5b45-4565-b3c4-b79878de67d2                                                                 |
| `roomID`                                                                                                 | **string*                                                                                                | :heavy_minus_sign:                                                                                       | N/A                                                                                                      | 2swovpy1fnunu                                                                                            |


### Response

**[*operations.CreateLobbyDeprecatedResponse](../../pkg/models/operations/createlobbydeprecatedresponse.md), error**
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
	"cloudapi"
	"cloudapi/pkg/models/shared"
	"context"
	"log"
)

func main() {
    s := cloudapi.New(
        cloudapi.WithAppID(cloudapi.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")),
    )


    var appID *string = cloudapi.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")

    var region *shared.Region = shared.RegionFrankfurt.ToPointer()

    ctx := context.Background()
    res, err := s.LobbyV2.ListActivePublicLobbiesDeprecatedV2(ctx, appID, region)
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
| `region`                                                                                | [*shared.Region](../../pkg/models/shared/region.md)                                     | :heavy_minus_sign:                                                                      | Region to filter by. If omitted, active public lobbies in all regions will be returned. |                                                                                         |


### Response

**[*operations.ListActivePublicLobbiesDeprecatedV2Response](../../pkg/models/operations/listactivepubliclobbiesdeprecatedv2response.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ~~GetLobbyInfo~~

Get details for a lobby.

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
    res, err := s.LobbyV2.GetLobbyInfo(ctx, roomID, appID)
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

**[*operations.GetLobbyInfoResponse](../../pkg/models/operations/getlobbyinforesponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 404                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ~~SetLobbyState~~

Set the state of a lobby. State is intended to be set by the server and must be smaller than 1MB. Use this endpoint to store match data like live player count to enforce max number of clients or persist end-game data (i.e. winner or final scores).

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

    setLobbyStateParams := shared.SetLobbyStateParams{
        State: shared.SetLobbyStateParamsState{},
    }

    var appID *string = cloudapi.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")

    ctx := context.Background()
    res, err := s.LobbyV2.SetLobbyState(ctx, roomID, setLobbyStateParams, appID)
    if err != nil {
        log.Fatal(err)
    }
    if res.Lobby != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  | Example                                                                      |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |                                                                              |
| `roomID`                                                                     | *string*                                                                     | :heavy_check_mark:                                                           | N/A                                                                          | 2swovpy1fnunu                                                                |
| `setLobbyStateParams`                                                        | [shared.SetLobbyStateParams](../../pkg/models/shared/setlobbystateparams.md) | :heavy_check_mark:                                                           | N/A                                                                          |                                                                              |
| `appID`                                                                      | **string*                                                                    | :heavy_minus_sign:                                                           | N/A                                                                          | app-af469a92-5b45-4565-b3c4-b79878de67d2                                     |


### Response

**[*operations.SetLobbyStateResponse](../../pkg/models/operations/setlobbystateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,404,422,429    | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
