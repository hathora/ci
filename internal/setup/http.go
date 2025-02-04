package setup

import (
	"github.com/hashicorp/go-cleanhttp"

	sdk "github.com/hathora/cloud-sdk-go/hathoracloud"

	"github.com/hathora/ci/internal/httputil"
)

func HTTPClient(loggingVerbosity int) sdk.HTTPClient {
	client := cleanhttp.DefaultClient()
	client.Transport = httputil.ContentTypeRoundTripper(client.Transport)
	if loggingVerbosity < 2 {
		return client
	}

	client.Transport = httputil.LoggingRoundTripper(client.Transport, loggingVerbosity)
	return client
}
