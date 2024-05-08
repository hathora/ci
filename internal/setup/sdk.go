package setup

import (
	"github.com/hathora/ci/internal/cloudapi/sdk"
)

func SDK(token string, url string) *cloudapi.SDK {
	sdk := cloudapi.New(
		cloudapi.WithSecurity(token),
		// TODO replace with env var
		cloudapi.WithAppID("app-af469a92-5b45-4565-b3c4-b79878de67d2"),
		cloudapi.WithServerURL(url))
	
	return sdk
}