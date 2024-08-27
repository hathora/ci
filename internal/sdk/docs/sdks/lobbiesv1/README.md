# LobbiesV1
(*LobbiesV1*)

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
	"github.com/hathora/ci/internal/sdk"
	"github.com/hathora/ci/internal/sdk/models/operations"
	"os"
	"github.com/hathora/ci/internal/sdk/models/shared"
	"context"
	"log"
)

func main() {
    s := sdk.New(
        sdk.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
    )
    security := operations.CreatePrivateLobbyDeprecatedSecurity{
            PlayerAuth: os.Getenv("PLAYER_AUTH"),
        }

    var appID *string = sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")

    var region *shared.Region = shared.RegionLondon.ToPointer()

    var local *bool = sdk.Bool(false)
    ctx := context.Background()
    res, err := s.LobbiesV1.CreatePrivateLobbyDeprecated(ctx, security, appID, region, local)
    if err != nil {
        log.Fatal(err)
    }
    if res.RoomID != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        | Example                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |                                                                                                                    |
| `security`                                                                                                         | [operations.CreatePrivateLobbyDeprecatedSecurity](../../models/operations/createprivatelobbydeprecatedsecurity.md) | :heavy_check_mark:                                                                                                 | The security requirements to use for the request.                                                                  |                                                                                                                    |
| `appID`                                                                                                            | **string*                                                                                                          | :heavy_minus_sign:                                                                                                 | N/A                                                                                                                | app-af469a92-5b45-4565-b3c4-b79878de67d2                                                                           |
| `region`                                                                                                           | [*shared.Region](../../models/shared/region.md)                                                                    | :heavy_minus_sign:                                                                                                 | N/A                                                                                                                |                                                                                                                    |
| `local`                                                                                                            | **bool*                                                                                                            | :heavy_minus_sign:                                                                                                 | N/A                                                                                                                |                                                                                                                    |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |                                                                                                                    |


### Response

**[*operations.CreatePrivateLobbyDeprecatedResponse](../../models/operations/createprivatelobbydeprecatedresponse.md), error**
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
	"github.com/hathora/ci/internal/sdk"
	"github.com/hathora/ci/internal/sdk/models/operations"
	"os"
	"github.com/hathora/ci/internal/sdk/models/shared"
	"context"
	"log"
)

func main() {
    s := sdk.New(
        sdk.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
    )
    security := operations.CreatePublicLobbyDeprecatedSecurity{
            PlayerAuth: os.Getenv("PLAYER_AUTH"),
        }

    var appID *string = sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")

    var region *shared.Region = shared.RegionFrankfurt.ToPointer()

    var local *bool = sdk.Bool(false)
    ctx := context.Background()
    res, err := s.LobbiesV1.CreatePublicLobbyDeprecated(ctx, security, appID, region, local)
    if err != nil {
        log.Fatal(err)
    }
    if res.RoomID != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      | Example                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |                                                                                                                  |
| `security`                                                                                                       | [operations.CreatePublicLobbyDeprecatedSecurity](../../models/operations/createpubliclobbydeprecatedsecurity.md) | :heavy_check_mark:                                                                                               | The security requirements to use for the request.                                                                |                                                                                                                  |
| `appID`                                                                                                          | **string*                                                                                                        | :heavy_minus_sign:                                                                                               | N/A                                                                                                              | app-af469a92-5b45-4565-b3c4-b79878de67d2                                                                         |
| `region`                                                                                                         | [*shared.Region](../../models/shared/region.md)                                                                  | :heavy_minus_sign:                                                                                               | N/A                                                                                                              |                                                                                                                  |
| `local`                                                                                                          | **bool*                                                                                                          | :heavy_minus_sign:                                                                                               | N/A                                                                                                              |                                                                                                                  |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |                                                                                                                  |


### Response

**[*operations.CreatePublicLobbyDeprecatedResponse](../../models/operations/createpubliclobbydeprecatedresponse.md), error**
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
	"github.com/hathora/ci/internal/sdk"
	"github.com/hathora/ci/internal/sdk/models/shared"
	"context"
	"log"
)

func main() {
    s := sdk.New(
        sdk.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
    )
    var appID *string = sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")

    var local *bool = sdk.Bool(false)

    var region *shared.Region = shared.RegionSydney.ToPointer()
    ctx := context.Background()
    res, err := s.LobbiesV1.ListActivePublicLobbiesDeprecatedV1(ctx, appID, local, region)
    if err != nil {
        log.Fatal(err)
    }
    if res.Lobbies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `appID`                                                  | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | app-af469a92-5b45-4565-b3c4-b79878de67d2                 |
| `local`                                                  | **bool*                                                  | :heavy_minus_sign:                                       | N/A                                                      |                                                          |
| `region`                                                 | [*shared.Region](../../models/shared/region.md)          | :heavy_minus_sign:                                       | N/A                                                      |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |


### Response

**[*operations.ListActivePublicLobbiesDeprecatedV1Response](../../models/operations/listactivepubliclobbiesdeprecatedv1response.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 404,429            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
