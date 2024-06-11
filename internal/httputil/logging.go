package httputil

import (
	"net/http"
	"net/http/httputil"

	"go.uber.org/zap"
)

type loggingRoundTripper struct {
	verbosity  int
	logger     *zap.Logger
	underlying http.RoundTripper
}

func (l *loggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	l.beforeRequest(req)
	res, err := l.underlying.RoundTrip(req)
	if err != nil {
		l.afterError(res, err)
	} else {
		l.afterSuccess(res)
	}

	return res, err
}

var _ http.RoundTripper = (*loggingRoundTripper)(nil)

func LoggingRoundTripper(underlying http.RoundTripper, verbosity int) http.RoundTripper {
	return &loggingRoundTripper{
		verbosity:  verbosity,
		logger:     zap.L().Named("http.client"),
		underlying: underlying,
	}
}

func (h *loggingRoundTripper) beforeRequest(req *http.Request) {
	// if the verbosity is all the way up, dump outgoing request bodies to logs
	if h.verbosity > 2 {
		reqDump, err := httputil.DumpRequestOut(req, true)
		if err == nil {
			h.logger.Debug(
				"request",
				zap.String("http.method", req.Method),
				zap.Stringer("http.url", req.URL),
				zap.String("http.request.dump", string(reqDump)),
			)
			return
		}
	}

	h.logger.Debug(
		"request",
		zap.String("method", req.Method),
		zap.Stringer("url", req.URL),
	)
}

func (h *loggingRoundTripper) afterSuccess(res *http.Response) {
	// if the verbosity is all the way up, dump outgoing response bodies to logs
	if h.verbosity > 2 {
		resDump, err := httputil.DumpResponse(res, true)
		if err == nil {
			h.logger.Debug(
				"response",
				zap.String("http.method", res.Request.Method),
				zap.Stringer("http.url", res.Request.URL),
				zap.Int("http.status", res.StatusCode),
				zap.String("http.response.dump", string(resDump)),
			)
			return
		}
	}

	h.logger.Debug(
		"response",
		zap.String("http.method", res.Request.Method),
		zap.Stringer("http.url", res.Request.URL),
		zap.Int("http.status", res.StatusCode),
	)
}

func (h *loggingRoundTripper) afterError(res *http.Response, err error) {
	method := "unknown"
	url := "unknown"
	status := 0

	if res != nil && res.Request != nil {
		method = res.Request.Method
		url = res.Request.URL.String()
		status = res.StatusCode
	}

	h.logger.Debug(
		"response",
		zap.String("http.method", method),
		zap.String("http.url", url),
		zap.Int("http.status", status),
		zap.Error(err),
	)
}
