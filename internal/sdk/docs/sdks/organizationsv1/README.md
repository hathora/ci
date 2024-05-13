# OrganizationsV1
(*OrganizationsV1*)

### Available Operations

* [InviteUser](#inviteuser)
* [RescindInvite](#rescindinvite)
* [GetOrgPendingInvites](#getorgpendinginvites)
* [GetUserPendingInvites](#getuserpendinginvites)
* [AcceptInvite](#acceptinvite)
* [RejectInvite](#rejectinvite)

## InviteUser

### Example Usage

```go
package main

import(
	"github.com/hathora/ci/internal/sdk/pkg/models/shared"
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


    var orgID string = "org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"

    createUserInvite := shared.CreateUserInvite{
        UserEmail: "<value>",
    }

    ctx := context.Background()
    res, err := s.OrganizationsV1.InviteUser(ctx, orgID, createUserInvite)
    if err != nil {
        log.Fatal(err)
    }
    if res.OrgPermission != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            | Example                                                                |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |                                                                        |
| `orgID`                                                                | *string*                                                               | :heavy_check_mark:                                                     | N/A                                                                    | org-6f706e83-0ec1-437a-9a46-7d4281eb2f39                               |
| `createUserInvite`                                                     | [shared.CreateUserInvite](../../pkg/models/shared/createuserinvite.md) | :heavy_check_mark:                                                     | N/A                                                                    |                                                                        |


### Response

**[*operations.InviteUserResponse](../../pkg/models/operations/inviteuserresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,422,429        | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## RescindInvite

### Example Usage

```go
package main

import(
	"github.com/hathora/ci/internal/sdk/pkg/models/shared"
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


    var orgID string = "org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"

    rescindUserInvite := shared.RescindUserInvite{
        UserEmail: "<value>",
    }

    ctx := context.Background()
    res, err := s.OrganizationsV1.RescindInvite(ctx, orgID, rescindUserInvite)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              | Example                                                                  |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |                                                                          |
| `orgID`                                                                  | *string*                                                                 | :heavy_check_mark:                                                       | N/A                                                                      | org-6f706e83-0ec1-437a-9a46-7d4281eb2f39                                 |
| `rescindUserInvite`                                                      | [shared.RescindUserInvite](../../pkg/models/shared/rescinduserinvite.md) | :heavy_check_mark:                                                       | N/A                                                                      |                                                                          |


### Response

**[*operations.RescindInviteResponse](../../pkg/models/operations/rescindinviteresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,404,429,500    | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetOrgPendingInvites

### Example Usage

```go
package main

import(
	"github.com/hathora/ci/internal/sdk/pkg/models/shared"
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


    var orgID string = "org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"

    ctx := context.Background()
    res, err := s.OrganizationsV1.GetOrgPendingInvites(ctx, orgID)
    if err != nil {
        log.Fatal(err)
    }
    if res.OrgInvitesPage != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `orgID`                                               | *string*                                              | :heavy_check_mark:                                    | N/A                                                   | org-6f706e83-0ec1-437a-9a46-7d4281eb2f39              |


### Response

**[*operations.GetOrgPendingInvitesResponse](../../pkg/models/operations/getorgpendinginvitesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,429            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetUserPendingInvites

### Example Usage

```go
package main

import(
	"github.com/hathora/ci/internal/sdk/pkg/models/shared"
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

    ctx := context.Background()
    res, err := s.OrganizationsV1.GetUserPendingInvites(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.OrgInvitesPage != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetUserPendingInvitesResponse](../../pkg/models/operations/getuserpendinginvitesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,429            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## AcceptInvite

### Example Usage

```go
package main

import(
	"github.com/hathora/ci/internal/sdk/pkg/models/shared"
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


    var orgID string = "org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"

    ctx := context.Background()
    res, err := s.OrganizationsV1.AcceptInvite(ctx, orgID)
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
| `orgID`                                               | *string*                                              | :heavy_check_mark:                                    | N/A                                                   | org-6f706e83-0ec1-437a-9a46-7d4281eb2f39              |


### Response

**[*operations.AcceptInviteResponse](../../pkg/models/operations/acceptinviteresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,404,429        | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## RejectInvite

### Example Usage

```go
package main

import(
	"github.com/hathora/ci/internal/sdk/pkg/models/shared"
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


    var orgID string = "org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"

    ctx := context.Background()
    res, err := s.OrganizationsV1.RejectInvite(ctx, orgID)
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
| `orgID`                                               | *string*                                              | :heavy_check_mark:                                    | N/A                                                   | org-6f706e83-0ec1-437a-9a46-7d4281eb2f39              |


### Response

**[*operations.RejectInviteResponse](../../pkg/models/operations/rejectinviteresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,404,429        | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
