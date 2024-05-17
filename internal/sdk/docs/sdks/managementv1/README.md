# ManagementV1
(*ManagementV1*)

## Overview

 

### Available Operations

* [SendVerificationEmail](#sendverificationemail)

## SendVerificationEmail

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
        sdk.WithAppID(sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")),
    )

    request := shared.VerificationEmailRequest{
        UserID: "<value>",
    }
    
    ctx := context.Background()
    res, err := s.ManagementV1.SendVerificationEmail(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.VerificationEmailSuccess != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [shared.VerificationEmailRequest](../../models/shared/verificationemailrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.SendVerificationEmailResponse](../../models/operations/sendverificationemailresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,429,500        | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
