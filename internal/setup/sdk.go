package setup

import (
	"github.com/hathora/ci/internal/sdk"
)

func SDK(token, baseURL string, loggingVerbosity int) *sdk.HathoraCloud {
	return sdk.New(
		sdk.WithSecurity(token),
		sdk.WithServerURL(baseURL),
		sdk.WithClient(HTTPClient(loggingVerbosity)),
	)
}
