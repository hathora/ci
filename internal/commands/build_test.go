package commands_test

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/hathora/ci/internal/commands"
	"github.com/hathora/ci/internal/mock"
)

func Test_BuildCommands_HelpText(t *testing.T) {
	t.Parallel()

	app := commands.App()
	err := app.Run(context.Background(), []string{"hathora", "build", "--help"})
	assert.Nil(t, err, "command returned an error")
}

func Test_Integration_BuildCommands_Happy(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	t.Parallel()

	tests := []struct {
		name           string
		command        string
		responseStatus int
		responseBody   string
		expectOutput   string
		expectRequest  func(t *testing.T, r *http.Request, requestBody *json.RawMessage)
		skip           bool
	}{
		{
			name:           "get build info",
			command:        "info --build-id bld-1",
			responseStatus: http.StatusOK,
			responseBody: `{
				"buildTag": "0.1.14-14c793",
				"regionalContainerTags": [
					{
						"containerTag": "string",
						"region": "Seattle"
					}
				],
				"imageSize": 0,
				"status": "created",
				"deletedAt": "2019-08-24T14:15:22Z",
				"finishedAt": "2019-08-24T14:15:22Z",
				"startedAt": "2019-08-24T14:15:22Z",
				"createdAt": "2019-08-24T14:15:22Z",
				"createdBy": "google-oauth2|107030234048588177467",
				"buildId": "bld-1",
				"appId": "app-af469a92-5b45-4565-b3c4-b79878de67d2"
			}`,
			expectRequest: func(t *testing.T, r *http.Request, requestBody *json.RawMessage) {
				assert.Equal(t, r.Method, http.MethodGet, "request method should be GET")
				assert.Equal(t, "/builds/v3/builds/bld-1", r.URL.Path, "request path should contain build id")
				assert.Empty(t, requestBody, "request body should be empty")
			},
		},
		{
			name:           "get all builds",
			command:        "list",
			responseStatus: http.StatusOK,
			responseBody: `{
				"builds": [
					{
						"buildTag": "0.1.14-14c793",
						"regionalContainerTags": [
						{
							"containerTag": "string",
							"region": "Seattle"
						}
						],
						"imageSize": 0,
						"status": "created",
						"deletedAt": "2019-08-24T14:15:22Z",
						"finishedAt": "2019-08-24T14:15:22Z",
						"startedAt": "2019-08-24T14:15:22Z",
						"createdAt": "2019-08-24T14:15:22Z",
						"createdBy": "google-oauth2|107030234048588177467",
						"buildId": "bld-1",
						"appId": "app-af469a92-5b45-4565-b3c4-b79878de67d2"
					}
				]
			}`,
			expectRequest: func(t *testing.T, r *http.Request, requestBody *json.RawMessage) {
				assert.Equal(t, r.Method, http.MethodGet, "request method should be GET")
				assert.Equal(t, "/builds/v3/builds", r.URL.Path, "request path should contain app id")
				assert.Empty(t, requestBody, "request body should be empty")
			},
		},
		{
			name:    "create a build",
			command: "create --build-tag test-build-tag --file TODO",
			// TODO setup the file input and mocks for multiple http responses
			skip:           true,
			responseStatus: http.StatusCreated,
			responseBody: `{
				"buildTag": "0.1.14-14c793",
				"regionalContainerTags": [
					{
						"containerTag": "string",
						"region": "Seattle"
					}
				],
				"imageSize": 0,
				"status": "created",
				"deletedAt": "2019-08-24T14:15:22Z",
				"finishedAt": "2019-08-24T14:15:22Z",
				"startedAt": "2019-08-24T14:15:22Z",
				"createdAt": "2019-08-24T14:15:22Z",
				"createdBy": "google-oauth2|107030234048588177467",
				"buildId": "bld-1",
				"appId": "app-af469a92-5b45-4565-b3c4-b79878de67d2"
			}`,
			expectRequest: func(t *testing.T, r *http.Request, requestBody *json.RawMessage) {
				assert.Equal(t, r.Method, http.MethodPost, "request method should be POST")
				assert.Equal(t, "/builds/v3/test-app-id/create", r.URL.Path, "request path should contain app id")
				assert.Equal(t, "application/json", r.Header.Get("Content-Type"), "request should have a JSON content type")
				assert.NotNil(t, requestBody, "request body should not be nil")
				assert.Equal(t, `{"buildTag":"test-build-tag"}`, string(*requestBody), "request body should have supplied build tag")
			},
		},
		{
			name:           "delete a build",
			command:        "delete --build-id bld-1",
			responseStatus: http.StatusNoContent,
			responseBody:   "",
			expectOutput: `{
				"status": "success"
				"message": "Build successfully deleted.",
				"code": 204
			}`,
			expectRequest: func(t *testing.T, r *http.Request, requestBody *json.RawMessage) {
				assert.Equal(t, r.Method, http.MethodDelete, "request method should be DELETE")
				assert.Equal(t, "/builds/v3/builds/bld-1", r.URL.Path, "request path should contain build id")
				assert.Empty(t, requestBody, "request body should be empty")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.skip {
				t.Skip()
			}
			h := mock.Hathora(t, mock.RespondsWithStatus(tt.responseStatus), mock.RespondsWithJSON([]byte(tt.responseBody)))
			app := commands.App()
			staticArgs := []string{
				"hathora",
				"-vvv",
				"--app-id",
				"test-app-id",
				"--token",
				"test-token",
				"--hathora-cloud-endpoint",
				h.Endpoint,
				"build",
			}
			testArgs := strings.Fields(tt.command)
			err := app.Run(context.Background(), append(staticArgs, testArgs...))
			assert.Nil(t, err, "command returned an error")
			request, body := h.ReceivedRequest()
			if tt.expectRequest != nil {
				require.NotNil(t, request, "request was nil")
				tt.expectRequest(t, request, body)
			}
		})
	}
}

func Test_Integration_BuildCommands_GlobalArgs(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	t.Parallel()

	tests := []struct {
		name           string
		command        string
		responseStatus int
		responseBody   string
		expectOutput   string
		expectRequest  func(t *testing.T, r *http.Request, requestBody *json.RawMessage)
	}{
		{
			name:           "use global args after domain-level command",
			command:        "build --app-id test-app-id --token test-token info --build-id bld-1",
			responseStatus: http.StatusOK,
			responseBody: `{
				"buildTag": "0.1.14-14c793",
				"regionalContainerTags": [
					{
						"containerTag": "string",
						"region": "Seattle"
					}
				],
				"imageSize": 0,
				"status": "created",
				"deletedAt": "2019-08-24T14:15:22Z",
				"finishedAt": "2019-08-24T14:15:22Z",
				"startedAt": "2019-08-24T14:15:22Z",
				"createdAt": "2019-08-24T14:15:22Z",
				"createdBy": "google-oauth2|107030234048588177467",
				"buildId": "bld-1",
				"appId": "app-af469a92-5b45-4565-b3c4-b79878de67d2"
			}`,
			expectRequest: func(t *testing.T, r *http.Request, requestBody *json.RawMessage) {
				assert.Equal(t, r.Method, http.MethodGet, "request method should be GET")
				assert.Equal(t, "/builds/v3/builds/bld-1", r.URL.Path, "request path should contain build id")
				assert.Empty(t, requestBody, "request body should be empty")
			},
		},
		{
			name:           "use global args after action-level command",
			command:        "build info --build-id bld-1 --app-id test-app-id --token test-token",
			responseStatus: http.StatusOK,
			responseBody: `{
				"buildTag": "0.1.14-14c793",
				"regionalContainerTags": [
					{
						"containerTag": "string",
						"region": "Seattle"
					}
				],
				"imageSize": 0,
				"status": "created",
				"deletedAt": "2019-08-24T14:15:22Z",
				"finishedAt": "2019-08-24T14:15:22Z",
				"startedAt": "2019-08-24T14:15:22Z",
				"createdAt": "2019-08-24T14:15:22Z",
				"createdBy": "google-oauth2|107030234048588177467",
				"buildId": "bld-1",
				"appId": "app-af469a92-5b45-4565-b3c4-b79878de67d2"
			}`,
			expectRequest: func(t *testing.T, r *http.Request, requestBody *json.RawMessage) {
				assert.Equal(t, r.Method, http.MethodGet, "request method should be GET")
				assert.Equal(t, "/builds/v3/builds/bld-1", r.URL.Path, "request path should contain build id")
				assert.Empty(t, requestBody, "request body should be empty")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := mock.Hathora(t, mock.RespondsWithStatus(tt.responseStatus), mock.RespondsWithJSON([]byte(tt.responseBody)))
			app := commands.App()
			staticArgs := []string{
				"hathora",
				"-vvv",
				"--hathora-cloud-endpoint",
				h.Endpoint,
			}
			testArgs := strings.Fields(tt.command)
			t.Log(append(staticArgs, testArgs...))
			err := app.Run(context.Background(), append(staticArgs, testArgs...))
			assert.Nil(t, err, "command returned an error")
			request, body := h.ReceivedRequest()
			if tt.expectRequest != nil {
				require.NotNil(t, request, "request was nil")
				tt.expectRequest(t, request, body)
			}
		})
	}
}
