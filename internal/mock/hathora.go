package mock

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hathora/ci/internal/sdk"
	"github.com/stretchr/testify/require"
)

type HathoraIntegration struct {
	Endpoint string
	SDK      *sdk.SDK

	mock *mockHathora
}

func (h *HathoraIntegration) ReceivedRequest() (*http.Request, *json.RawMessage) {
	if h == nil || h.mock == nil {
		return nil, nil
	}
	return h.mock.capturedRequest, h.mock.capturedRequestBody
}

func Hathora(t *testing.T, opts ...mockHathoraOption) *HathoraIntegration {
	t.Helper()
	m := &mockHathora{
		t: t,
	}
	for _, opt := range opts {
		opt(m)
	}
	mockServer := httptest.NewServer(m)
	t.Cleanup(mockServer.Close)

	endpoint := mockServer.URL
	sdk := sdk.New(sdk.WithServerURL(endpoint))
	return &HathoraIntegration{
		Endpoint: endpoint,
		SDK:      sdk,
		mock:     m,
	}
}

type mockHathora struct {
	t                   *testing.T
	capturedRequest     *http.Request
	capturedRequestBody *json.RawMessage
	cannedResponse      []byte
	cannedStatus        int
}

type mockHathoraOption func(*mockHathora)

func RespondsWithJSON(response []byte) mockHathoraOption {
	return func(m *mockHathora) {
		m.t.Helper()
		m.cannedResponse = response
	}
}

func RespondsWithStatus(status int) mockHathoraOption {
	return func(m *mockHathora) {
		m.t.Helper()
		m.cannedStatus = status
	}
}

func RespondsWithJSONObject(response any) mockHathoraOption {
	return func(m *mockHathora) {
		m.t.Helper()
		body, err := json.Marshal(response)
		if err != nil {
			require.NoError(m.t, err, "harness failed to marshal response")
		}
		m.cannedResponse = body
	}
}

func (m *mockHathora) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.t.Helper()
	m.capturedRequest = r
	m.capturedRequestBody = nil
	body, err := io.ReadAll(r.Body)
	require.NoError(m.t, err, "failed to read request body")
	jsonBody := json.RawMessage(body)
	m.capturedRequestBody = &jsonBody

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(m.cannedStatus)

	_, err = w.Write(m.cannedResponse)
	require.NoError(m.t, err, "failed to write mock response")
}
