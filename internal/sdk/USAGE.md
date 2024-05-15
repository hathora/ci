<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"github.com/hathora/ci/internal/sdk"
	"github.com/hathora/ci/internal/sdk/models/shared"
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
	res, err := s.AppV1.GetApps(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.ApplicationWithLatestDeploymentAndBuilds != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->