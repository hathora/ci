package setup

import (
	"github.com/hashicorp/go-cleanhttp"
	"github.com/hathora/ci/internal/httputil"
	"github.com/hathora/ci/internal/sdk"
)

func HTTPClient(loggingVerbosity int) sdk.HTTPClient {
	if loggingVerbosity < 2 {
		return cleanhttp.DefaultClient()
	}
	client := cleanhttp.DefaultClient()
	client.Transport = httputil.LoggingRoundTripper(client.Transport, loggingVerbosity)
	return client
}
