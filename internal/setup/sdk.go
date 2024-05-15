package setup

import (
	"github.com/hathora/ci/internal/sdk"
	"github.com/hathora/ci/internal/sdk/pkg/models/shared"
)

func SDK(token string, url string) *sdk.SDK {
	sdkInstance := sdk.New(
		sdk.WithSecurity(shared.Security{ HathoraDevToken: sdk.String(token)}),
		sdk.WithServerURL(url))
	return sdkInstance
}