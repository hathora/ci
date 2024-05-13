# MetricsV1
(*MetricsV1*)

## Overview

Operations to get metrics by [process](https://hathora.dev/docs/concepts/hathora-entities#process). We store 72 hours of metrics data.

### Available Operations

* [GetMetrics](#getmetrics) - Get metrics for a [process](https://hathora.dev/docs/concepts/hathora-entities#process) using `appId` and `processId`.

## GetMetrics

Get metrics for a [process](https://hathora.dev/docs/concepts/hathora-entities#process) using `appId` and `processId`.

### Example Usage

```go
package main

import(
	"github.com/hathora/ci/internal/sdk/pkg/models/shared"
	"github.com/hathora/ci/internal/sdk"
	"context"
	"github.com/hathora/ci/internal/sdk/pkg/models/operations"
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
    res, err := s.MetricsV1.GetMetrics(ctx, operations.GetMetricsRequest{
        ProcessID: "cbfcddd2-0006-43ae-996c-995fff7bed2e",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MetricsData != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetMetricsRequest](../../pkg/models/operations/getmetricsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.GetMetricsResponse](../../pkg/models/operations/getmetricsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,404,422,500    | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
