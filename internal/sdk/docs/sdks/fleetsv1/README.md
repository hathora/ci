# FleetsV1
(*FleetsV1*)

## Overview

Operations to manage and view a [fleet](https://hathora.dev/docs/concepts/hathora-entities#fleet).

### Available Operations

* [GetFleets](#getfleets) - GetFleets
* [GetFleetRegion](#getfleetregion) - GetFleetRegion
* [UpdateFleetRegion](#updatefleetregion) - UpdateFleetRegion
* [GetFleetMetrics](#getfleetmetrics) - GetFleetMetrics

## GetFleets

Returns an array of [fleets](https://hathora.dev/docs/concepts/hathora-entities#fleet).

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

    res, err := s.FleetsV1.GetFleets(ctx, hathoracloud.String("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"))
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
| `orgID`                                                  | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | org-6f706e83-0ec1-437a-9a46-7d4281eb2f39                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*components.FleetsPage](../../models/components/fleetspage.md), error**

### Errors

| Error Type       | Status Code      | Content Type     |
| ---------------- | ---------------- | ---------------- |
| errors.APIError  | 401, 404, 429    | application/json |
| errors.SDKError  | 4XX, 5XX         | \*/\*            |

## GetFleetRegion

Gets the configuration for a given [fleet](https://hathora.dev/docs/concepts/hathora-entities#fleet) in a region.

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

    res, err := s.FleetsV1.GetFleetRegion(ctx, "<id>", components.RegionSaoPaulo, hathoracloud.String("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"))
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
| `fleetID`                                                | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |                                                          |
| `region`                                                 | [components.Region](../../models/components/region.md)   | :heavy_check_mark:                                       | N/A                                                      |                                                          |
| `orgID`                                                  | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | org-6f706e83-0ec1-437a-9a46-7d4281eb2f39                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*components.FleetRegion](../../models/components/fleetregion.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| errors.APIError    | 401, 404, 422, 429 | application/json   |
| errors.SDKError    | 4XX, 5XX           | \*/\*              |

## UpdateFleetRegion

Updates the configuration for a given [fleet](https://hathora.dev/docs/concepts/hathora-entities#fleet) in a region.

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

    err := s.FleetsV1.UpdateFleetRegion(ctx, "<id>", components.RegionSingapore, components.FleetRegionConfig{
        CloudMinVcpus: 511402,
    }, hathoracloud.String("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"))
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  | Example                                                                      |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |                                                                              |
| `fleetID`                                                                    | *string*                                                                     | :heavy_check_mark:                                                           | N/A                                                                          |                                                                              |
| `region`                                                                     | [components.Region](../../models/components/region.md)                       | :heavy_check_mark:                                                           | N/A                                                                          |                                                                              |
| `fleetRegionConfig`                                                          | [components.FleetRegionConfig](../../models/components/fleetregionconfig.md) | :heavy_check_mark:                                                           | N/A                                                                          |                                                                              |
| `orgID`                                                                      | **string*                                                                    | :heavy_minus_sign:                                                           | N/A                                                                          | org-6f706e83-0ec1-437a-9a46-7d4281eb2f39                                     |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |                                                                              |

### Response

**error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| errors.APIError         | 401, 404, 422, 429, 500 | application/json        |
| errors.SDKError         | 4XX, 5XX                | \*/\*                   |

## GetFleetMetrics

Gets metrics for a [fleet](https://hathora.dev/docs/concepts/hathora-entities#fleet) in a region.

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
        hathoracloud.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
        hathoracloud.WithOrgID("org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"),
        hathoracloud.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
    )

    res, err := s.FleetsV1.GetFleetMetrics(ctx, operations.GetFleetMetricsRequest{
        FleetID: "<id>",
        Region: components.RegionLondon,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetFleetMetricsRequest](../../models/operations/getfleetmetricsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*components.FleetMetricsData](../../models/components/fleetmetricsdata.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| errors.APIError         | 401, 404, 422, 429, 500 | application/json        |
| errors.SDKError         | 4XX, 5XX                | \*/\*                   |