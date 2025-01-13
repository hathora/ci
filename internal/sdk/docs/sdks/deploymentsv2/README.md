# DeploymentsV2
(*DeploymentsV2*)

## Overview

### Available Operations

* [~~GetDeploymentsV2Deprecated~~](#getdeploymentsv2deprecated) - GetDeploymentsV2Deprecated :warning: **Deprecated**
* [~~GetLatestDeploymentV2Deprecated~~](#getlatestdeploymentv2deprecated) - GetLatestDeploymentV2Deprecated :warning: **Deprecated**
* [~~GetDeploymentInfoV2Deprecated~~](#getdeploymentinfov2deprecated) - GetDeploymentInfoV2Deprecated :warning: **Deprecated**
* [~~CreateDeploymentV2Deprecated~~](#createdeploymentv2deprecated) - CreateDeploymentV2Deprecated :warning: **Deprecated**

## ~~GetDeploymentsV2Deprecated~~

Returns an array of [deployments](https://hathora.dev/docs/concepts/hathora-entities#deployment) for an [application](https://hathora.dev/docs/concepts/hathora-entities#application).

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
        hathoracloud.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
        hathoracloud.WithOrgID("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"),
        hathoracloud.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
    )

    res, err := s.DeploymentsV2.GetDeploymentsV2Deprecated(ctx, hathoracloud.String("app-af469a92-5b45-4565-b3c4-b79878de67d2"))
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
| `appID`                                                  | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | app-af469a92-5b45-4565-b3c4-b79878de67d2                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[[]components.DeploymentV2](../../.md), error**

### Errors

| Error Type       | Status Code      | Content Type     |
| ---------------- | ---------------- | ---------------- |
| errors.APIError  | 401, 404, 429    | application/json |
| errors.SDKError  | 4XX, 5XX         | \*/\*            |

## ~~GetLatestDeploymentV2Deprecated~~

Get the latest [deployment](https://hathora.dev/docs/concepts/hathora-entities#deployment) for an [application](https://hathora.dev/docs/concepts/hathora-entities#application).

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
        hathoracloud.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
        hathoracloud.WithOrgID("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"),
        hathoracloud.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
    )

    res, err := s.DeploymentsV2.GetLatestDeploymentV2Deprecated(ctx, hathoracloud.String("app-af469a92-5b45-4565-b3c4-b79878de67d2"))
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
| `appID`                                                  | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | app-af469a92-5b45-4565-b3c4-b79878de67d2                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*components.DeploymentV2](../../models/components/deploymentv2.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| errors.APIError    | 401, 404, 422, 429 | application/json   |
| errors.SDKError    | 4XX, 5XX           | \*/\*              |

## ~~GetDeploymentInfoV2Deprecated~~

Get details for a [deployment](https://hathora.dev/docs/concepts/hathora-entities#deployment).

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
        hathoracloud.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
        hathoracloud.WithOrgID("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"),
        hathoracloud.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
    )

    res, err := s.DeploymentsV2.GetDeploymentInfoV2Deprecated(ctx, 1, hathoracloud.String("app-af469a92-5b45-4565-b3c4-b79878de67d2"))
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
| `deploymentID`                                           | *int*                                                    | :heavy_check_mark:                                       | N/A                                                      | 1                                                        |
| `appID`                                                  | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | app-af469a92-5b45-4565-b3c4-b79878de67d2                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*components.DeploymentV2](../../models/components/deploymentv2.md), error**

### Errors

| Error Type       | Status Code      | Content Type     |
| ---------------- | ---------------- | ---------------- |
| errors.APIError  | 401, 404, 429    | application/json |
| errors.SDKError  | 4XX, 5XX         | \*/\*            |

## ~~CreateDeploymentV2Deprecated~~

Create a new [deployment](https://hathora.dev/docs/concepts/hathora-entities#deployment). Creating a new deployment means all new rooms created will use the latest deployment configuration, but existing games in progress will not be affected.

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

    res, err := s.DeploymentsV2.CreateDeploymentV2Deprecated(ctx, 1, components.DeploymentConfigV2{
        IdleTimeoutEnabled: false,
        Env: []components.DeploymentConfigV2Env{
            components.DeploymentConfigV2Env{
                Value: "TRUE",
                Name: "EULA",
            },
            components.DeploymentConfigV2Env{
                Value: "TRUE",
                Name: "EULA",
            },
        },
        RoomsPerProcess: 3,
        AdditionalContainerPorts: []components.ContainerPort{
            components.ContainerPort{
                TransportType: components.TransportTypeTLS,
                Port: 8000,
                Name: "default",
            },
        },
        TransportType: components.TransportTypeTLS,
        ContainerPort: 4000,
        RequestedMemoryMB: 1024,
        RequestedCPU: 0.5,
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    | Example                                                                        |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |                                                                                |
| `buildID`                                                                      | *int*                                                                          | :heavy_check_mark:                                                             | N/A                                                                            | 1                                                                              |
| `deploymentConfigV2`                                                           | [components.DeploymentConfigV2](../../models/components/deploymentconfigv2.md) | :heavy_check_mark:                                                             | N/A                                                                            |                                                                                |
| `appID`                                                                        | **string*                                                                      | :heavy_minus_sign:                                                             | N/A                                                                            | app-af469a92-5b45-4565-b3c4-b79878de67d2                                       |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |                                                                                |

### Response

**[*components.DeploymentV2](../../models/components/deploymentv2.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.APIError              | 400, 401, 404, 422, 429, 500 | application/json             |
| errors.SDKError              | 4XX, 5XX                     | \*/\*                        |