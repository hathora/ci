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

func Test_DeploymentCommands_HelpText(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	t.Parallel()

	app := commands.App()
	err := app.Run(context.Background(), []string{"hathora", "deployment", "--help"})
	assert.Nil(t, err, "command returned an error")
}

func Test_Integration_DeploymentCommands_Happy(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	t.Parallel()

	tests := []struct {
		name           string
		skip           string
		command        string
		responseStatus int
		responseBody   string
		expectOutput   string
		expectRequest  func(t *testing.T, r *http.Request, requestBody *json.RawMessage)
	}{
		{
			name:           "get deployment info",
			command:        "info --deployment-id 1",
			responseStatus: http.StatusOK,
			responseBody: `{
				"idleTimeoutEnabled": true,
				"env": [
					{
						"value": "TRUE",
						"name": "EULA"
					}
				],
				"roomsPerProcess": 3,
				"additionalContainerPorts": [{
					"transportType": "tcp",
					"port": 4000,
					"name": "debug"
				}],
				"defaultContainerPort": {
					"transportType": "tcp",
					"port": 8000,
					"name": "default"
				},
				"createdAt": "2019-08-24T14:15:22Z",
				"createdBy": "google-oauth2|107030234048588177467",
				"requestedMemoryMB": 1024,
				"requestedCPU": 0.5,
				"deploymentId": 1,
				"buildId": 1,
				"appId": "app-af469a92-5b45-4565-b3c4-b79878de67d2"
			}`,
			expectRequest: func(t *testing.T, r *http.Request, requestBody *json.RawMessage) {
				assert.Equal(t, r.Method, http.MethodGet, "request method should be GET")
				assert.Equal(t, "/deployments/v2/test-app-id/info/1", r.URL.Path, "request path should contain app id and deplyoment id")
				assert.Empty(t, requestBody, "request body should be empty")
			},
		},
		{
			name:           "get latest deployment",
			command:        "latest",
			responseStatus: http.StatusOK,
			responseBody: `{
				"idleTimeoutEnabled": true,
				"env": [
					{
						"value": "TRUE",
						"name": "EULA"
					}
				],
				"roomsPerProcess": 3,
				"additionalContainerPorts": [{
					"transportType": "tcp",
					"port": 4000,
					"name": "debug"
				}],
				"defaultContainerPort": {
					"transportType": "tcp",
					"port": 8000,
					"name": "default"
				},
				"createdAt": "2019-08-24T14:15:22Z",
				"createdBy": "google-oauth2|107030234048588177467",
				"requestedMemoryMB": 1024,
				"requestedCPU": 0.5,
				"deploymentId": 1,
				"buildId": 1,
				"appId": "app-af469a92-5b45-4565-b3c4-b79878de67d2"
			}`,
			expectRequest: func(t *testing.T, r *http.Request, requestBody *json.RawMessage) {
				assert.Equal(t, r.Method, http.MethodGet, "request method should be GET")
				assert.Equal(t, "/deployments/v2/test-app-id/latest", r.URL.Path, "request path should contain app id")
				assert.Empty(t, requestBody, "request body should be empty")
			},
		},
		{
			name:           "get all deployments",
			command:        "list",
			responseStatus: http.StatusOK,
			responseBody: `[
				{
					"idleTimeoutEnabled": true,
					"env": [
						{
							"value": "TRUE",
							"name": "EULA"
						}
					],
					"roomsPerProcess": 3,
					"additionalContainerPorts": [{
						"transportType": "tcp",
						"port": 4000,
						"name": "debug"
					}],
					"defaultContainerPort": {
						"transportType": "tcp",
						"port": 8000,
						"name": "default"
					},
					"createdAt": "2019-08-24T14:15:22Z",
					"createdBy": "google-oauth2|107030234048588177467",
					"requestedMemoryMB": 1024,
					"requestedCPU": 0.5,
					"deploymentId": 1,
					"buildId": 1,
					"appId": "app-af469a92-5b45-4565-b3c4-b79878de67d2"
				}
			]`,
			expectRequest: func(t *testing.T, r *http.Request, requestBody *json.RawMessage) {
				assert.Equal(t, r.Method, http.MethodGet, "request method should be GET")
				assert.Equal(t, "/deployments/v2/test-app-id/list", r.URL.Path, "request path should contain app id")
				assert.Empty(t, requestBody, "request body should be empty")
			},
		},
		{
			name: "create a deployment",
			command: "create --build-id 1 --idle-timeout-enabled --rooms-per-process 3" +
				" --transport-type tcp --container-port 8000 --requested-memory-mb 1024 --requested-cpu 0.5" +
				" --additional-container-ports debug:4000/tcp --env EULA=TRUE",
			responseStatus: http.StatusCreated,
			responseBody: `{
				"idleTimeoutEnabled": true,
				"env": [
					{
						"value": "TRUE",
						"name": "EULA"
					}
				],
				"roomsPerProcess": 3,
				"additionalContainerPorts": [
					{
						"transportType": "tcp",
						"port": 4000,
						"name": "debug"
					}
				],
				"transportType": "tcp",
				"containerPort": 8000,
				"requestedMemoryMB": 1024,
				"requestedCPU": 0.5
			}`,
			expectRequest: func(t *testing.T, r *http.Request, requestBody *json.RawMessage) {
				assert.Equal(t, r.Method, http.MethodPost, "request method should be POST")
				assert.Equal(t, "/deployments/v2/test-app-id/create/1", r.URL.Path, "request path should contain app id and build id")
				assert.NotNil(t, requestBody, "request body should not be nil")
				assert.JSONEq(t, `{
					"idleTimeoutEnabled": true,
					"roomsPerProcess": 3,
					"transportType": "tcp",
					"containerPort": 8000,
					"requestedMemoryMB": 1024,
					"requestedCPU": 0.5,
					"additionalContainerPorts": [
						{
							"transportType": "tcp",
							"port": 4000,
							"name": "debug"
						}
					],
					"env": [
						{
							"value": "TRUE",
							"name": "EULA"
						}
					]
				}`, string(*requestBody), "request body should match expected")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.skip != "" {
				t.Skip(tt.skip)
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
				"deployment",
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

type mockResponse struct {
	status int
	body   string
}

func Test_Integration_DeploymentCommands_CreateFromLatest(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	t.Parallel()

	tests := []struct {
		name          string
		skip          string
		command       string
		responses     []mockResponse
		expectOutput  string
		expectRequest func(t *testing.T, r *http.Request, requestBody *json.RawMessage)
	}{
		{
			name:    "create a deployment from latest with no overrides",
			command: "create --from-latest",
			skip:    "this test is currently flakey when run in with other package tests",
			responses: []mockResponse{
				// get latest response
				{
					status: http.StatusOK,
					body: `{
						"idleTimeoutEnabled": true,
						"env": [
							{
								"value": "TRUE",
								"name": "EULA"
							}
						],
						"roomsPerProcess": 3,
						"additionalContainerPorts": [{
							"transportType": "tcp",
							"port": 4000,
							"name": "debug"
						}],
						"defaultContainerPort": {
							"transportType": "tcp",
							"port": 8000,
							"name": "default"
						},
						"createdAt": "2019-08-24T14:15:22Z",
						"createdBy": "google-oauth2|107030234048588177467",
						"requestedMemoryMB": 1024,
						"requestedCPU": 0.5,
						"deploymentId": 1,
						"buildId": 1,
						"appId": "app-af469a92-5b45-4565-b3c4-b79878de67d2"
					}`,
				},
				// create deployment response
				{
					status: http.StatusCreated,
					body: `{
						"idleTimeoutEnabled": true,
						"env": [
							{
								"value": "TRUE",
								"name": "EULA"
							}
						],
						"roomsPerProcess": 3,
						"additionalContainerPorts": [
							{
								"transportType": "tcp",
								"port": 4000,
								"name": "debug"
							}
						],
						"transportType": "tcp",
						"containerPort": 8000,
						"requestedMemoryMB": 1024,
						"requestedCPU": 0.5
					}`,
				},
			},
			expectOutput: `{
				"idleTimeoutEnabled": true,
				"env": [
					{
						"value": "TRUE",
						"name": "EULA"
					}
				],
				"roomsPerProcess": 3,
				"additionalContainerPorts": [
					{
						"transportType": "tcp",
						"port": 4000,
						"name": "debug"
					}
				],
				"transportType": "tcp",
				"containerPort": 8000,
				"requestedMemoryMB": 1024,
				"requestedCPU": 0.5
			}`,
			expectRequest: func(t *testing.T, r *http.Request, requestBody *json.RawMessage) {
				assert.Equal(t, r.Method, http.MethodPost, "request method should be POST")
				assert.Equal(t, "/deployments/v2/test-app-id/create/1", r.URL.Path, "request path should contain app id and build id")
				assert.NotNil(t, requestBody, "request body should not be nil")
				assert.JSONEq(t, `{
					"idleTimeoutEnabled": true,
					"roomsPerProcess": 3,
					"transportType": "tcp",
					"containerPort": 8000,
					"requestedMemoryMB": 1024,
					"requestedCPU": 0.5,
					"additionalContainerPorts": [
						{
							"transportType": "tcp",
							"port": 4000,
							"name": "debug"
						}
					],
					"env": [
						{
							"value": "TRUE",
							"name": "EULA"
						}
					]
				}`, string(*requestBody), "request body should match expected")
			},
		},
		{
			name: "create a deployment from latest with some overrides",
			command: "create --from-latest --build-id 3 --idle-timeout-enabled=false --rooms-per-process 5 " +
				" --env EULA=TRUE --env ADDITIONAL_ENV_VAR=1",
			skip: "this test is currently flakey when run in with other package tests",
			responses: []mockResponse{
				// get latest response
				{
					status: http.StatusOK,
					body: `{
						"idleTimeoutEnabled": true,
						"env": [
							{
								"value": "TRUE",
								"name": "EULA"
							}
						],
						"roomsPerProcess": 3,
						"additionalContainerPorts": [{
							"transportType": "tcp",
							"port": 4000,
							"name": "debug"
						}],
						"defaultContainerPort": {
							"transportType": "tcp",
							"port": 8000,
							"name": "default"
						},
						"createdAt": "2019-08-24T14:15:22Z",
						"createdBy": "google-oauth2|107030234048588177467",
						"requestedMemoryMB": 1024,
						"requestedCPU": 0.5,
						"deploymentId": 1,
						"buildId": 1,
						"appId": "app-af469a92-5b45-4565-b3c4-b79878de67d2"
					}`,
				},
				// create deployment response
				{
					status: http.StatusCreated,
					body: `{
						"idleTimeoutEnabled": false,
						"env": [
							{
								"value": "TRUE",
								"name": "EULA"
							},
							{
								"value": "1",
								"name": "ADDITIONAL_ENV_VAR"
							}
						],
						"roomsPerProcess": 5,
						"additionalContainerPorts": [
							{
								"transportType": "tcp",
								"port": 4000,
								"name": "debug"
							}
						],
						"transportType": "tcp",
						"containerPort": 8000,
						"requestedMemoryMB": 1024,
						"requestedCPU": 0.5
					}`,
				},
			},
			expectOutput: `{
				"idleTimeoutEnabled": false,
				"env": [
					{
						"value": "TRUE",
						"name": "EULA"
					},
					{
						"value": "1",
						"name": "ADDITIONAL_ENV_VAR"
					}
				],
				"roomsPerProcess": 5,
				"additionalContainerPorts": [
					{
						"transportType": "tcp",
						"port": 4000,
						"name": "debug"
					}
				],
				"transportType": "tcp",
				"containerPort": 8000,
				"requestedMemoryMB": 1024,
				"requestedCPU": 0.5
			}`,
			expectRequest: func(t *testing.T, r *http.Request, requestBody *json.RawMessage) {
				assert.Equal(t, r.Method, http.MethodPost, "request method should be POST")
				assert.Equal(t, "/deployments/v2/test-app-id/create/3", r.URL.Path, "request path should contain app id and build id")
				assert.NotNil(t, requestBody, "request body should not be nil")
				assert.JSONEq(t, `{
					"idleTimeoutEnabled": false,
					"roomsPerProcess": 5,
					"transportType": "tcp",
					"containerPort": 8000,
					"requestedMemoryMB": 1024,
					"requestedCPU": 0.5,
					"additionalContainerPorts": [
						{
							"transportType": "tcp",
							"port": 4000,
							"name": "debug"
						}
					],
					"env": [
						{
							"value": "TRUE",
							"name": "EULA"
						},
						{
							"value": "1",
							"name": "ADDITIONAL_ENV_VAR"
						}
					]
				}`, string(*requestBody), "request body should match expected")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.skip != "" {
				t.Skip(tt.skip)
			}

			var opts []mock.HathoraOption
			for _, r := range tt.responses {
				opts = append(opts, mock.RespondsWithStatus(r.status))
				opts = append(opts, mock.RespondsWithJSON([]byte(r.body)))
			}
			h := mock.Hathora(t, opts...)
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
				"deployment",
			}
			testArgs := strings.Fields(tt.command)
			t.Log(append(staticArgs, testArgs...))
			err := app.Run(context.Background(), append(staticArgs, testArgs...))
			assert.Nil(t, err, "command returned an error")
			request, body := h.ReceivedNthRequest(1)
			if tt.expectRequest != nil {
				require.NotNil(t, request, "request was nil")
				tt.expectRequest(t, request, body)
			}
		})
	}
}
