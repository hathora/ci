<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"github.com/hathora/ci/internal/sdk"
	"github.com/hathora/ci/internal/sdk/models/shared"
	"log"
	"os"
)

func main() {
	s := sdk.New(
		sdk.WithSecurity(shared.Security{
			HathoraDevToken: sdk.String(os.Getenv("HATHORA_DEV_TOKEN")),
		}),
		sdk.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
	)
	var orgID string = "org-6f706e83-0ec1-437a-9a46-7d4281eb2f39"
	ctx := context.Background()
	res, err := s.TokensV1.GetOrgTokens(ctx, orgID)
	if err != nil {
		log.Fatal(err)
	}
	if res.ListOrgTokens != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->