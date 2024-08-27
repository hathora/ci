# BuildsV2
(*BuildsV2*)

### Available Operations

* [GetBuildsV2Deprecated](#getbuildsv2deprecated) - Returns an array of [builds](https://hathora.dev/docs/concepts/hathora-entities#build) for an [application](https://hathora.dev/docs/concepts/hathora-entities#application).
* [GetBuildInfoV2Deprecated](#getbuildinfov2deprecated) - Get details for a [build](https://hathora.dev/docs/concepts/hathora-entities#build).
* [CreateBuildV2Deprecated](#createbuildv2deprecated) - Creates a new [build](https://hathora.dev/docs/concepts/hathora-entities#build). Responds with a `buildId` that you must pass to [`RunBuild()`](https://hathora.dev/api#tag/BuildV1/operation/RunBuild) to build the game server artifact. You can optionally pass in a `buildTag` to associate an external version with a build.
* [CreateBuildWithUploadURLV2Deprecated](#createbuildwithuploadurlv2deprecated) - Creates a new [build](https://hathora.dev/docs/concepts/hathora-entities#build) with `uploadUrl` that can be used to upload the build to before calling `runBuild`. Responds with a `buildId` that you must pass to [`RunBuild()`](https://hathora.dev/api#tag/BuildV1/operation/RunBuild) to build the game server artifact. You can optionally pass in a `buildTag` to associate an external version with a build.
* [CreateWithMultipartUploadsV2Deprecated](#createwithmultipartuploadsv2deprecated) - Creates a new [build](https://hathora.dev/docs/concepts/hathora-entities#build) with optional `multipartUploadUrls` that can be used to upload larger builds in parts before calling `runBuild`. Responds with a `buildId` that you must pass to [`RunBuild()`](https://hathora.dev/api#tag/BuildV1/operation/RunBuild) to build the game server artifact. You can optionally pass in a `buildTag` to associate an external version with a build.
* [DeleteBuildV2Deprecated](#deletebuildv2deprecated) - Delete a [build](https://hathora.dev/docs/concepts/hathora-entities#build). All associated metadata is deleted.
* [RunBuildV2Deprecated](#runbuildv2deprecated) - Builds a game server artifact from a tarball you provide. Pass in the `buildId` generated from [`CreateBuild()`](https://hathora.dev/api#tag/BuildV1/operation/CreateBuild).

## GetBuildsV2Deprecated

Returns an array of [builds](https://hathora.dev/docs/concepts/hathora-entities#build) for an [application](https://hathora.dev/docs/concepts/hathora-entities#application).

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
    var appID *string = sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")
    ctx := context.Background()
    res, err := s.BuildsV2.GetBuildsV2Deprecated(ctx, appID)
    if err != nil {
        log.Fatal(err)
    }
    if res.Builds != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `appID`                                                  | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | app-af469a92-5b45-4565-b3c4-b79878de67d2                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |


### Response

**[*operations.GetBuildsV2DeprecatedResponse](../../models/operations/getbuildsv2deprecatedresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,404,429        | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetBuildInfoV2Deprecated

Get details for a [build](https://hathora.dev/docs/concepts/hathora-entities#build).

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
    var buildID int = 1

    var appID *string = sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")
    ctx := context.Background()
    res, err := s.BuildsV2.GetBuildInfoV2Deprecated(ctx, buildID, appID)
    if err != nil {
        log.Fatal(err)
    }
    if res.Build != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `buildID`                                                | *int*                                                    | :heavy_check_mark:                                       | N/A                                                      | 1                                                        |
| `appID`                                                  | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | app-af469a92-5b45-4565-b3c4-b79878de67d2                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |


### Response

**[*operations.GetBuildInfoV2DeprecatedResponse](../../models/operations/getbuildinfov2deprecatedresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,404,429        | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateBuildV2Deprecated

Creates a new [build](https://hathora.dev/docs/concepts/hathora-entities#build). Responds with a `buildId` that you must pass to [`RunBuild()`](https://hathora.dev/api#tag/BuildV1/operation/RunBuild) to build the game server artifact. You can optionally pass in a `buildTag` to associate an external version with a build.

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
    createBuildParams := shared.CreateBuildParams{
        BuildTag: sdk.String("0.1.14-14c793"),
    }

    var appID *string = sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")
    ctx := context.Background()
    res, err := s.BuildsV2.CreateBuildV2Deprecated(ctx, createBuildParams, appID)
    if err != nil {
        log.Fatal(err)
    }
    if res.Build != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          | Example                                                              |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |                                                                      |
| `createBuildParams`                                                  | [shared.CreateBuildParams](../../models/shared/createbuildparams.md) | :heavy_check_mark:                                                   | N/A                                                                  |                                                                      |
| `appID`                                                              | **string*                                                            | :heavy_minus_sign:                                                   | N/A                                                                  | app-af469a92-5b45-4565-b3c4-b79878de67d2                             |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |                                                                      |


### Response

**[*operations.CreateBuildV2DeprecatedResponse](../../models/operations/createbuildv2deprecatedresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,404,429,500    | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateBuildWithUploadURLV2Deprecated

Creates a new [build](https://hathora.dev/docs/concepts/hathora-entities#build) with `uploadUrl` that can be used to upload the build to before calling `runBuild`. Responds with a `buildId` that you must pass to [`RunBuild()`](https://hathora.dev/api#tag/BuildV1/operation/RunBuild) to build the game server artifact. You can optionally pass in a `buildTag` to associate an external version with a build.

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
    createBuildParams := shared.CreateBuildParams{
        BuildTag: sdk.String("0.1.14-14c793"),
    }

    var appID *string = sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")
    ctx := context.Background()
    res, err := s.BuildsV2.CreateBuildWithUploadURLV2Deprecated(ctx, createBuildParams, appID)
    if err != nil {
        log.Fatal(err)
    }
    if res.BuildWithUploadURL != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          | Example                                                              |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |                                                                      |
| `createBuildParams`                                                  | [shared.CreateBuildParams](../../models/shared/createbuildparams.md) | :heavy_check_mark:                                                   | N/A                                                                  |                                                                      |
| `appID`                                                              | **string*                                                            | :heavy_minus_sign:                                                   | N/A                                                                  | app-af469a92-5b45-4565-b3c4-b79878de67d2                             |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |                                                                      |


### Response

**[*operations.CreateBuildWithUploadURLV2DeprecatedResponse](../../models/operations/createbuildwithuploadurlv2deprecatedresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,404,429,500    | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateWithMultipartUploadsV2Deprecated

Creates a new [build](https://hathora.dev/docs/concepts/hathora-entities#build) with optional `multipartUploadUrls` that can be used to upload larger builds in parts before calling `runBuild`. Responds with a `buildId` that you must pass to [`RunBuild()`](https://hathora.dev/api#tag/BuildV1/operation/RunBuild) to build the game server artifact. You can optionally pass in a `buildTag` to associate an external version with a build.

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
    createMultipartBuildParams := shared.CreateMultipartBuildParams{
        BuildTag: sdk.String("0.1.14-14c793"),
        BuildSizeInBytes: 3146.66,
    }

    var appID *string = sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")
    ctx := context.Background()
    res, err := s.BuildsV2.CreateWithMultipartUploadsV2Deprecated(ctx, createMultipartBuildParams, appID)
    if err != nil {
        log.Fatal(err)
    }
    if res.BuildWithMultipartUrls != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            | Example                                                                                |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |                                                                                        |
| `createMultipartBuildParams`                                                           | [shared.CreateMultipartBuildParams](../../models/shared/createmultipartbuildparams.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |                                                                                        |
| `appID`                                                                                | **string*                                                                              | :heavy_minus_sign:                                                                     | N/A                                                                                    | app-af469a92-5b45-4565-b3c4-b79878de67d2                                               |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |                                                                                        |


### Response

**[*operations.CreateWithMultipartUploadsV2DeprecatedResponse](../../models/operations/createwithmultipartuploadsv2deprecatedresponse.md), error**
| Error Object        | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| sdkerrors.APIError  | 400,401,404,429,500 | application/json    |
| sdkerrors.SDKError  | 4xx-5xx             | */*                 |

## DeleteBuildV2Deprecated

Delete a [build](https://hathora.dev/docs/concepts/hathora-entities#build). All associated metadata is deleted.

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
    var buildID int = 1

    var appID *string = sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")
    ctx := context.Background()
    res, err := s.BuildsV2.DeleteBuildV2Deprecated(ctx, buildID, appID)
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
| `buildID`                                                | *int*                                                    | :heavy_check_mark:                                       | N/A                                                      | 1                                                        |
| `appID`                                                  | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | app-af469a92-5b45-4565-b3c4-b79878de67d2                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |


### Response

**[*operations.DeleteBuildV2DeprecatedResponse](../../models/operations/deletebuildv2deprecatedresponse.md), error**
| Error Object        | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| sdkerrors.APIError  | 401,404,422,429,500 | application/json    |
| sdkerrors.SDKError  | 4xx-5xx             | */*                 |

## RunBuildV2Deprecated

Builds a game server artifact from a tarball you provide. Pass in the `buildId` generated from [`CreateBuild()`](https://hathora.dev/api#tag/BuildV1/operation/CreateBuild).

### Example Usage

```go
package main

import(
	"github.com/hathora/ci/internal/sdk/models/shared"
	"os"
	"github.com/hathora/ci/internal/sdk"
	"github.com/hathora/ci/internal/sdk/models/operations"
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
    var buildID int = 1

    requestBody := operations.RunBuildV2DeprecatedRequestBody{}

    var appID *string = sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")
    ctx := context.Background()
    res, err := s.BuildsV2.RunBuildV2Deprecated(ctx, buildID, requestBody, appID)
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseStream != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              | Example                                                                                                  |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |                                                                                                          |
| `buildID`                                                                                                | *int*                                                                                                    | :heavy_check_mark:                                                                                       | N/A                                                                                                      | 1                                                                                                        |
| `requestBody`                                                                                            | [operations.RunBuildV2DeprecatedRequestBody](../../models/operations/runbuildv2deprecatedrequestbody.md) | :heavy_check_mark:                                                                                       | N/A                                                                                                      |                                                                                                          |
| `appID`                                                                                                  | **string*                                                                                                | :heavy_minus_sign:                                                                                       | N/A                                                                                                      | app-af469a92-5b45-4565-b3c4-b79878de67d2                                                                 |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |                                                                                                          |


### Response

**[*operations.RunBuildV2DeprecatedResponse](../../models/operations/runbuildv2deprecatedresponse.md), error**
| Error Object        | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| sdkerrors.APIError  | 400,401,404,429,500 | application/json    |
| sdkerrors.SDKError  | 4xx-5xx             | */*                 |
