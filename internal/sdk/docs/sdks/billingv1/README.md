# BillingV1
(*BillingV1*)

## Overview

 

### Available Operations

* [GetBalance](#getbalance)
* [GetUpcomingInvoiceItems](#getupcominginvoiceitems)
* [GetUpcomingInvoiceTotal](#getupcominginvoicetotal)
* [GetPaymentMethod](#getpaymentmethod)
* [InitStripeCustomerPortalURL](#initstripecustomerportalurl)
* [GetInvoices](#getinvoices)

## GetBalance

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

    ctx := context.Background()
    res, err := s.BillingV1.GetBalance(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Number != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |


### Response

**[*operations.GetBalanceResponse](../../models/operations/getbalanceresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,404,429        | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetUpcomingInvoiceItems

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

    ctx := context.Background()
    res, err := s.BillingV1.GetUpcomingInvoiceItems(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.InvoiceItemPage != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |


### Response

**[*operations.GetUpcomingInvoiceItemsResponse](../../models/operations/getupcominginvoiceitemsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,404,429        | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetUpcomingInvoiceTotal

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

    ctx := context.Background()
    res, err := s.BillingV1.GetUpcomingInvoiceTotal(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |


### Response

**[*operations.GetUpcomingInvoiceTotalResponse](../../models/operations/getupcominginvoicetotalresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,404,429        | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetPaymentMethod

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

    ctx := context.Background()
    res, err := s.BillingV1.GetPaymentMethod(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentMethod != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |


### Response

**[*operations.GetPaymentMethodResponse](../../models/operations/getpaymentmethodresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,404,429,500    | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## InitStripeCustomerPortalURL

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
    request := shared.CustomerPortalURL{
        ReturnURL: "<value>",
    }
    ctx := context.Background()
    res, err := s.BillingV1.InitStripeCustomerPortalURL(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.String != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [shared.CustomerPortalURL](../../models/shared/customerportalurl.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |


### Response

**[*operations.InitStripeCustomerPortalURLResponse](../../models/operations/initstripecustomerportalurlresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,404,429        | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetInvoices

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

    ctx := context.Background()
    res, err := s.BillingV1.GetInvoices(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Invoices != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |


### Response

**[*operations.GetInvoicesResponse](../../models/operations/getinvoicesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,404,429        | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
