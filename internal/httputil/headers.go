package httputil

import (
	"net/http"
)

const (
	NameContentType             = "Content-Type"
	ValueApplicationJSON        = "application/json"
	ValueApplicationOctetStream = "application/octet-stream"
	ValueApplicationXML         = "application/xml"
)

type contentTypeRoundTripper struct {
	underlying http.RoundTripper
}

// RoundTrip This is a stop gap until speakeasy respects that empty content type == octet-stream
func (c *contentTypeRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	res, err := c.underlying.RoundTrip(req)
	if err != nil {
		return res, err
	}

	if res.Header == nil {
		return res, err
	}

	if res.Header.Get(NameContentType) == "" {
		res.Header.Set(NameContentType, ValueApplicationOctetStream)
	}

	return res, err
}

var _ http.RoundTripper = (*contentTypeRoundTripper)(nil)

func ContentTypeRoundTripper(underlying http.RoundTripper) http.RoundTripper {
	return &contentTypeRoundTripper{
		underlying: underlying,
	}
}
