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
	"cloudapi"
	"context"
	"cloudapi/pkg/models/shared"
	"log"
)

func main() {
    s := cloudapi.New(
        cloudapi.WithAppID(cloudapi.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")),
    )

    ctx := context.Background()
    res, err := s.ManagementV1.SendVerificationEmail(ctx, shared.VerificationEmailRequest{
        UserID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VerificationEmailSuccess != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [shared.VerificationEmailRequest](../../pkg/models/shared/verificationemailrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.SendVerificationEmailResponse](../../pkg/models/operations/sendverificationemailresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,429,500        | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
