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
					Name:          "default",
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
				RequestedMemoryMB: 555,
				RequestedCPU:      0.5,
			},
			expect: `
AppID  DeploymentID  BuildID  CreatedAt  CreatedBy  IdleTimeoutEnabled  RoomsPerProcess  DefaultContainerPort  AdditionalContainerPorts  Env        RequestedCPU  RequestedMemoryMB
appID  2             1        12:00AM    createdBy  true                3                default:3000/tcp      debug:4000/tcp            EULA=TRUE  0.500000      555.000000`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var deployment shared.DeploymentV2
			var envVar shared.DeploymentV2Env
			var containerPort shared.ContainerPort
			var timestamp time.Time
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
				output.WithFormatter(timestamp,
					func(t time.Time) string {
						return t.Format(time.Kitchen)
					},
				),
			)
			var buf bytes.Buffer
			actualErr := formatter.Write(tt.input, &buf)
			assert.NoError(t, actualErr)
			actual := buf.String()
			assert.Equal(t, tt.expect, actual)
		})
	}
}
