# LobbyV1
(*LobbyV1*)

## Overview

Deprecated. Use [LobbyV3](https://hathora.dev/api#tag/LobbyV3).

### Available Operations

* [~~CreatePrivateLobbyDeprecated~~](#createprivatelobbydeprecated) - :warning: **Deprecated**
* [~~CreatePublicLobbyDeprecated~~](#createpubliclobbydeprecated) - :warning: **Deprecated**
* [~~ListActivePublicLobbiesDeprecatedV1~~](#listactivepubliclobbiesdeprecatedv1) - :warning: **Deprecated**

## ~~CreatePrivateLobbyDeprecated~~

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


    var appID *string = cloudapi.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")

    var region *shared.Region = shared.RegionLondon.ToPointer()

    var local *bool = cloudapi.Bool(false)

    operationSecurity := operations.CreatePrivateLobbyDeprecatedSecurity{
            PlayerAuth: "<YOUR_BEARER_TOKEN_HERE>",
        }

    ctx := context.Background()
    res, err := s.LobbyV1.CreatePrivateLobbyDeprecated(ctx, operationSecurity, appID, region, local)
    if err != nil {
        log.Fatal(err)
    }
    if res.RoomID != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            | Example                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |                                                                                                                        |
| `security`                                                                                                             | [operations.CreatePrivateLobbyDeprecatedSecurity](../../pkg/models/operations/createprivatelobbydeprecatedsecurity.md) | :heavy_check_mark:                                                                                                     | The security requirements to use for the request.                                                                      |                                                                                                                        |
| `appID`                                                                                                                | **string*                                                                                                              | :heavy_minus_sign:                                                                                                     | N/A                                                                                                                    | app-af469a92-5b45-4565-b3c4-b79878de67d2                                                                               |
| `region`                                                                                                               | [*shared.Region](../../pkg/models/shared/region.md)                                                                    | :heavy_minus_sign:                                                                                                     | N/A                                                                                                                    |                                                                                                                        |
| `local`                                                                                                                | **bool*                                                                                                                | :heavy_minus_sign:                                                                                                     | N/A                                                                                                                    |                                                                                                                        |


### Response

**[*operations.CreatePrivateLobbyDeprecatedResponse](../../pkg/models/operations/createprivatelobbydeprecatedresponse.md), error**
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.APIError          | 400,401,402,404,422,429,500 | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |

## ~~CreatePublicLobbyDeprecated~~

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


    var appID *string = cloudapi.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")

    var region *shared.Region = shared.RegionLondon.ToPointer()

    var local *bool = cloudapi.Bool(false)

    operationSecurity := operations.CreatePublicLobbyDeprecatedSecurity{
            PlayerAuth: "<YOUR_BEARER_TOKEN_HERE>",
        }

    ctx := context.Background()
    res, err := s.LobbyV1.CreatePublicLobbyDeprecated(ctx, operationSecurity, appID, region, local)
    if err != nil {
        log.Fatal(err)
    }
    if res.RoomID != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          | Example                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |                                                                                                                      |
| `security`                                                                                                           | [operations.CreatePublicLobbyDeprecatedSecurity](../../pkg/models/operations/createpubliclobbydeprecatedsecurity.md) | :heavy_check_mark:                                                                                                   | The security requirements to use for the request.                                                                    |                                                                                                                      |
| `appID`                                                                                                              | **string*                                                                                                            | :heavy_minus_sign:                                                                                                   | N/A                                                                                                                  | app-af469a92-5b45-4565-b3c4-b79878de67d2                                                                             |
| `region`                                                                                                             | [*shared.Region](../../pkg/models/shared/region.md)                                                                  | :heavy_minus_sign:                                                                                                   | N/A                                                                                                                  |                                                                                                                      |
| `local`                                                                                                              | **bool*                                                                                                              | :heavy_minus_sign:                                                                                                   | N/A                                                                                                                  |                                                                                                                      |


### Response

**[*operations.CreatePublicLobbyDeprecatedResponse](../../pkg/models/operations/createpubliclobbydeprecatedresponse.md), error**
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.APIError          | 400,401,402,404,422,429,500 | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |

## ~~ListActivePublicLobbiesDeprecatedV1~~

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

    var local *bool = cloudapi.Bool(false)

    var region *shared.Region = shared.RegionTokyo.ToPointer()

    ctx := context.Background()
    res, err := s.LobbyV1.ListActivePublicLobbiesDeprecatedV1(ctx, appID, local, region)
    if err != nil {
        log.Fatal(err)
    }
    if res.Lobbies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `appID`                                               | **string*                                             | :heavy_minus_sign:                                    | N/A                                                   | app-af469a92-5b45-4565-b3c4-b79878de67d2              |
| `local`                                               | **bool*                                               | :heavy_minus_sign:                                    | N/A                                                   |                                                       |
| `region`                                              | [*shared.Region](../../pkg/models/shared/region.md)   | :heavy_minus_sign:                                    | N/A                                                   |                                                       |


### Response

**[*operations.ListActivePublicLobbiesDeprecatedV1Response](../../pkg/models/operations/listactivepubliclobbiesdeprecatedv1response.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 404                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
