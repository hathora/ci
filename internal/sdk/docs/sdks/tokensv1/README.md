# TokensV1
(*TokensV1*)

## Overview

 

### Available Operations

* [GetOrgTokens](#getorgtokens) - List all organization tokens for a given org.
* [CreateOrgToken](#createorgtoken) - Create a new organization token.
* [RevokeOrgToken](#revokeorgtoken) - Revoke an organization token.

## GetOrgTokens

List all organization tokens for a given org.

### Example Usage

```go
package main

import(
	"github.com/hathora/ci/internal/sdk/models/shared"
	"os"
	"github.com/hathora/ci/internal/sdk"
	"context"
	"log"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            HathoraDevToken: sdk.String(os.Getenv("HATHORA_DEV_TOKEN")),
        }),
        sdk.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
    )
    var orgID string = "org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"
    ctx := context.Background()
    res, err := s.TokensV1.GetOrgTokens(ctx, orgID)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListOrgTokens != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `orgID`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | org-6f706e83-0ec1-437a-9a46-7d4281eb2f39                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |


### Response

**[*operations.GetOrgTokensResponse](../../models/operations/getorgtokensresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,404,429        | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateOrgToken

Create a new organization token.

### Example Usage

```go
package main

import(
	"github.com/hathora/ci/internal/sdk/models/shared"
	"os"
	"github.com/hathora/ci/internal/sdk"
	"context"
	"log"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            HathoraDevToken: sdk.String(os.Getenv("HATHORA_DEV_TOKEN")),
        }),
        sdk.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
    )
    var orgID string = "org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"

    createOrgToken := shared.CreateOrgToken{
        Name: "ci-token",
    }
    ctx := context.Background()
    res, err := s.TokensV1.CreateOrgToken(ctx, orgID, createOrgToken)
    if err != nil {
        log.Fatal(err)
    }
    if res.CreatedOrgToken != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    | Example                                                        |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |                                                                |
| `orgID`                                                        | *string*                                                       | :heavy_check_mark:                                             | N/A                                                            | org-6f706e83-0ec1-437a-9a46-7d4281eb2f39                       |
| `createOrgToken`                                               | [shared.CreateOrgToken](../../models/shared/createorgtoken.md) | :heavy_check_mark:                                             | N/A                                                            |                                                                |
| `opts`                                                         | [][operations.Option](../../models/operations/option.md)       | :heavy_minus_sign:                                             | The options for this request.                                  |                                                                |


### Response

**[*operations.CreateOrgTokenResponse](../../models/operations/createorgtokenresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,404,422,429    | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## RevokeOrgToken

Revoke an organization token.

### Example Usage

```go
package main

import(
	"github.com/hathora/ci/internal/sdk/models/shared"
	"os"
	"github.com/hathora/ci/internal/sdk"
	"context"
	"log"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            HathoraDevToken: sdk.String(os.Getenv("HATHORA_DEV_TOKEN")),
        }),
        sdk.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
    )
    var orgID string = "org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"

    var orgTokenID string = "org-token-af469a92-5b45-4565-b3c4-b79878de67d2"
    ctx := context.Background()
    res, err := s.TokensV1.RevokeOrgToken(ctx, orgID, orgTokenID)
    if err != nil {
        log.Fatal(err)
    }
    if res.Boolean != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `orgID`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | org-6f706e83-0ec1-437a-9a46-7d4281eb2f39                 |
| `orgTokenID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | org-token-af469a92-5b45-4565-b3c4-b79878de67d2           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |


### Response

**[*operations.RevokeOrgTokenResponse](../../models/operations/revokeorgtokenresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,404,429        | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
