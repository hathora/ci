package setup

import (
	"github.com/hathora/ci/internal/sdk"
	"github.com/hathora/ci/internal/sdk/models/components"
)

func SDK(token, baseURL string, loggingVerbosity int) *sdk.SDK {
	return sdk.New(
		sdk.WithSecurity(
			components.Security{
				HathoraDevToken: sdk.String(token),
			},
		),
		sdk.WithServerURL(baseURL),
		sdk.WithClient(HTTPClient(loggingVerbosity)),
	)
}
