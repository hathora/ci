<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"cloudapi"
	"cloudapi/pkg/models/shared"
	"context"
	"log"
)

func main() {
	s := cloudapi.New(
		cloudapi.WithSecurity(shared.Security{
			HathoraDevToken: cloudapi.String("<YOUR_BEARER_TOKEN_HERE>"),
		}),
		cloudapi.WithAppID(cloudapi.String("app-af469a92-5b45-4565-b3c4-b79878de67d2")),
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