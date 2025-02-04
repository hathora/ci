package setup

import (
	sdk "github.com/hathora/cloud-sdk-go/hathoracloud"
)

func SDK(token, baseURL string, loggingVerbosity int) *sdk.HathoraCloud {
	return sdk.New(
		sdk.WithSecurity(token),
		sdk.WithServerURL(baseURL),
		sdk.WithClient(HTTPClient(loggingVerbosity)),
	)
}
