package setup

import (
	"github.com/hathora/ci/internal/sdk"
	"github.com/hathora/ci/internal/sdk/models/shared"
)

func SDK(token, baseURL string, loggingVerbosity int) *sdk.SDK {
	return sdk.New(
		sdk.WithSecurity(
			shared.Security{
				HathoraDevToken: sdk.String(token),
			},
		),
		sdk.WithServerURL(baseURL),
		sdk.WithClient(HTTPClient(loggingVerbosity)),
	)
}
