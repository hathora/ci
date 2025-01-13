package mock

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hathora/ci/internal/httputil"
	"github.com/hathora/ci/internal/sdk"
	"github.com/stretchr/testify/require"
)

type HathoraIntegration struct {
	Endpoint string
	SDK      *sdk.HathoraCloud

	mock *mockHathora
}

func (h *HathoraIntegration) ReceivedRequest() (*http.Request, *json.RawMessage) {
	if h == nil || h.mock == nil {
		return nil, nil
	}
	requestIndex := len(h.mock.capturedRequests) - 1
	if requestIndex < 0 {
		return nil, nil
	}
	return h.mock.capturedRequests[requestIndex], h.mock.capturedRequestBodies[requestIndex]
}

func (h *HathoraIntegration) ReceivedRequests() ([]*http.Request, []*json.RawMessage) {
	if h == nil || h.mock == nil {
		return nil, nil
	}

	return h.mock.capturedRequests, h.mock.capturedRequestBodies
}

func (h *HathoraIntegration) ReceivedRequestCount() int {
	if h == nil || h.mock == nil {
		return 0
	}
	return len(h.mock.capturedRequests)
}

func (h *HathoraIntegration) ReceivedNthRequest(index int) (*http.Request, *json.RawMessage) {
	if h == nil || h.mock == nil {
		return nil, nil
	}

	if index < 0 || index >= len(h.mock.capturedRequests) {
		return nil, nil
	}

	return h.mock.capturedRequests[index], h.mock.capturedRequestBodies[index]
}

func Hathora(t *testing.T, opts ...HathoraOption) *HathoraIntegration {
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
	t                     *testing.T
	capturedRequests      []*http.Request
	capturedRequestBodies []*json.RawMessage
	nextResponseIndex     int
	cannedResponses       [][]byte
	cannedStatuses        []int
}

type HathoraOption func(*mockHathora)

func RespondsWithJSON(response []byte) HathoraOption {
	return func(m *mockHathora) {
		m.t.Helper()
		m.cannedResponses = append(m.cannedResponses, response)
	}
}

func RespondsWithStatus(status int) HathoraOption {
	return func(m *mockHathora) {
		m.t.Helper()
		m.cannedStatuses = append(m.cannedStatuses, status)
	}
}

func RespondsWithJSONObject(response any) HathoraOption {
	return func(m *mockHathora) {
		m.t.Helper()
		body, err := json.Marshal(response)
		if err != nil {
			require.NoError(m.t, err, "harness failed to marshal response")
		}
		m.cannedResponses = append(m.cannedResponses, body)
	}
}

func (m *mockHathora) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.t.Helper()
	m.capturedRequests = append(m.capturedRequests, r)
	body, err := io.ReadAll(r.Body)
	require.NoError(m.t, err, "failed to read request body")
	jsonBody := json.RawMessage(body)
	m.capturedRequestBodies = append(m.capturedRequestBodies, &jsonBody)

	w.Header().Set(httputil.NameContentType, httputil.ValueApplicationJSON)
	w.WriteHeader(m.cannedStatuses[m.nextResponseIndex])

	_, err = w.Write(m.cannedResponses[m.nextResponseIndex])
	require.NoError(m.t, err, "failed to write mock response")
	// don't increment if we're at the end of the responses list
	if m.nextResponseIndex < len(m.cannedResponses)-1 && m.nextResponseIndex < len(m.cannedStatuses)-1 {
		m.nextResponseIndex++
		return
	}
}
