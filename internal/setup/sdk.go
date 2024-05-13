package setup

import (
	"github.com/hathora/ci/internal/sdk"
	"github.com/hathora/ci/internal/sdk/pkg/models/shared"
)

func SDK(token string, url string) *sdk.SDK {
	sdkInstance := sdk.New(
		sdk.WithSecurity(shared.Security{ HathoraDevToken: sdk.String(token)}),
		sdk.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
		sdk.WithServerURL(url))
	
	return sdkInstance
}