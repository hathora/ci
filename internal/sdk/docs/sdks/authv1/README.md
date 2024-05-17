# AuthV1
(*AuthV1*)

## Overview

Operations that allow you to generate a Hathora-signed [JSON web token (JWT)](https://jwt.io/) for [player authentication](https://hathora.dev/docs/lobbies-and-matchmaking/auth-service).

### Available Operations

* [LoginAnonymous](#loginanonymous) - Returns a unique player token for an anonymous user.
* [LoginNickname](#loginnickname) - Returns a unique player token with a specified nickname for a user.
* [LoginGoogle](#logingoogle) - Returns a unique player token using a Google-signed OIDC `idToken`.

## LoginAnonymous

Returns a unique player token for an anonymous user.

### Example Usage

```go
package main

import(
	"github.com/hathora/ci/internal/sdk"
	"context"
	"log"
)

func main() {
    s := sdk.New(
        sdk.WithAppID(sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")),
    )

    var appID *string = sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")
    
    ctx := context.Background()
    res, err := s.AuthV1.LoginAnonymous(ctx, appID)
    if err != nil {
        log.Fatal(err)
    }
    if res.PlayerTokenObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `appID`                                               | **string*                                             | :heavy_minus_sign:                                    | N/A                                                   | app-af469a92-5b45-4565-b3c4-b79878de67d2              |


### Response

**[*operations.LoginAnonymousResponse](../../models/operations/loginanonymousresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 404                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## LoginNickname

Returns a unique player token with a specified nickname for a user.

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

    nicknameObject := shared.NicknameObject{
        Nickname: "squiddytwoshoes",
    }

    var appID *string = sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")
    
    ctx := context.Background()
    res, err := s.AuthV1.LoginNickname(ctx, nicknameObject, appID)
    if err != nil {
        log.Fatal(err)
    }
    if res.PlayerTokenObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    | Example                                                        |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |                                                                |
| `nicknameObject`                                               | [shared.NicknameObject](../../models/shared/nicknameobject.md) | :heavy_check_mark:                                             | N/A                                                            |                                                                |
| `appID`                                                        | **string*                                                      | :heavy_minus_sign:                                             | N/A                                                            | app-af469a92-5b45-4565-b3c4-b79878de67d2                       |


### Response

**[*operations.LoginNicknameResponse](../../models/operations/loginnicknameresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 404                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## LoginGoogle

Returns a unique player token using a Google-signed OIDC `idToken`.

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

    googleIDTokenObject := shared.GoogleIDTokenObject{
        IDToken: "eyJhbGciOiJSUzI1NiIsImtpZCI6ImZkNDhhNzUxMzhkOWQ0OGYwYWE2MzVlZjU2OWM0ZTE5NmY3YWU4ZDYiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJhY2NvdW50cy5nb29nbGUuY29tIiwiYXpwIjoiODQ4NDEyODI2Nzg4LW00bXNyYjZxNDRkbTJ1ZTNrZ3Z1aTBmcTdrZGE1NWxzLmFwcHMuZ29vZ2xldXNlcmNvbnRlbnQuY29tIiwiYXVkIjoiODQ4NDEyODI2Nzg4LW00bXNyYjZxNDRkbTJ1ZTNrZ3Z1aTBmcTdrZGE1NWxzLmFwcHMuZ29vZ2xldXNlcmNvbnRlbnQuY29tIiwic3ViIjoiMTE0NTQyMzMwNzI3MTU2MTMzNzc2IiwiZW1haWwiOiJocGFdkeivmeuzQGdtYWlsLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJhdF9oYXNoIjoidno1NGhhdTNxbnVR",
    }

    var appID *string = sdk.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")
    
    ctx := context.Background()
    res, err := s.AuthV1.LoginGoogle(ctx, googleIDTokenObject, appID)
    if err != nil {
        log.Fatal(err)
    }
    if res.PlayerTokenObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              | Example                                                                  |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |                                                                          |
| `googleIDTokenObject`                                                    | [shared.GoogleIDTokenObject](../../models/shared/googleidtokenobject.md) | :heavy_check_mark:                                                       | N/A                                                                      |                                                                          |
| `appID`                                                                  | **string*                                                                | :heavy_minus_sign:                                                       | N/A                                                                      | app-af469a92-5b45-4565-b3c4-b79878de67d2                                 |


### Response

**[*operations.LoginGoogleResponse](../../models/operations/logingoogleresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 401,404            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
