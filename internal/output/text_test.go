package output_test

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/hathora/ci/internal/output"
	"github.com/hathora/ci/internal/sdk/models/shared"
	"github.com/stretchr/testify/assert"
)

func Test_DeploymentTextOutput(t *testing.T) {
	ts := time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name   string
		input  shared.DeploymentV2
		expect string
	}{
		{
			name: "simple deployment",
			input: shared.DeploymentV2{
				IdleTimeoutEnabled: true,
				RoomsPerProcess:    3,
				AdditionalContainerPorts: []shared.ContainerPort{
					{
						TransportType: "tcp",
						Port:          4000,
						Name:          "debug",
					},
				},
				DefaultContainerPort: shared.ContainerPort{
					TransportType: "tcp",
					Port:          3000,
				},
				CreatedAt:    ts,
				CreatedBy:    "createdBy",
				AppID:        "appID",
				DeploymentID: 2,
				BuildID:      1,
				Env: []shared.DeploymentV2Env{
					{
						Name:  "EULA",
						Value: "TRUE",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var deployment shared.DeploymentV2
			var envVar shared.DeploymentV2Env
			var containerPort shared.ContainerPort
			formatter := output.TextFormat(
				output.WithFieldOrder(deployment,
					"AppID",
					"DeploymentID",
					"BuildID",
					"CreatedAt",
					"CreatedBy",
					"IdleTimeoutEnabled",
					"RoomsPerProcess",
					"DefaultContainerPort",
					"AdditionalContainerPorts",
					"Env",
				),
				output.WithFormatter(envVar,
					func(e shared.DeploymentV2Env) string {
						return fmt.Sprintf("%s=%s", e.Name, e.Value)
					},
				),
				output.WithFormatter(containerPort,
					func(cp shared.ContainerPort) string {
						return fmt.Sprintf("%s:%d/%s", cp.Name, cp.Port, cp.TransportType)
					},
				),
			)
			var buf bytes.Buffer
			actualErr := formatter.Write(tt.input, &buf)
			assert.NoError(t, actualErr)
			actual := string(buf.Bytes())
			assert.Equal(t, tt.expect, actual)
		})
	}
}
