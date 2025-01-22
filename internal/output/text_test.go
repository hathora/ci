package output_test

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/hathora/ci/internal/commands"
	sdk "github.com/hathora/cloud-sdk-go/hathoracloud"
	"github.com/hathora/cloud-sdk-go/hathoracloud/models/components"
)

func Test_DeploymentTextOutput(t *testing.T) {
	ts := time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name   string
		input  any
		expect [][]string
	}{
		{
			name: "single deployment",
			input: components.DeploymentV3{
				IdleTimeoutEnabled: true,
				RoomsPerProcess:    3,
				AdditionalContainerPorts: []components.ContainerPort{
					{
						TransportType: "tcp",
						Port:          4000,
						Name:          "debug",
					},
				},
				DefaultContainerPort: components.ContainerPort{
					TransportType: "tcp",
					Port:          3000,
					Name:          "default",
				},
				CreatedAt:    ts,
				CreatedBy:    "createdBy",
				AppID:        "appID",
				DeploymentID: "dep-2",
				BuildID:      "bld-1",
				Env: []components.DeploymentV3Env{
					{
						Name:  "EULA",
						Value: "TRUE",
					},
				},
				RequestedMemoryMB: 1024,
				RequestedCPU:      0.5,
			},
			expect: [][]string{
				{"DeploymentID", "BuildID", "CreatedAt", "IdleTimeoutEnabled", "RoomsPerProcess", "RequestedCPU", "RequestedMemory", "DefaultContainerPort", "AdditionalContainerPorts", "BuildTag", "DeploymentTag", "ExperimentalRequestedGPU"},
				{"dep-2", "bld-1", "2021-01-01T00:00:00Z", "true", "3", "0.5", "1.0", "GiB", "default:3000/tcp", "debug:4000/tcp", "null", "null", "null"},
			},
		},
		{
			name: "single deployment ptr",
			input: &components.DeploymentV3{
				IdleTimeoutEnabled: true,
				RoomsPerProcess:    3,
				AdditionalContainerPorts: []components.ContainerPort{
					{
						TransportType: "tcp",
						Port:          4000,
						Name:          "debug",
					},
				},
				DefaultContainerPort: components.ContainerPort{
					TransportType: "tcp",
					Port:          3000,
					Name:          "default",
				},
				CreatedAt:    ts,
				CreatedBy:    "createdBy",
				AppID:        "appID",
				DeploymentID: "dep-2",
				BuildID:      "bld-1",
				Env: []components.DeploymentV3Env{
					{
						Name:  "EULA",
						Value: "TRUE",
					},
				},
				RequestedMemoryMB: 1024,
				RequestedCPU:      0.5,
			},
			expect: [][]string{
				{"DeploymentID", "BuildID", "CreatedAt", "IdleTimeoutEnabled", "RoomsPerProcess", "RequestedCPU", "RequestedMemory", "DefaultContainerPort", "AdditionalContainerPorts", "BuildTag", "DeploymentTag", "ExperimentalRequestedGPU"},
				{"dep-2", "bld-1", "2021-01-01T00:00:00Z", "true", "3", "0.5", "1.0", "GiB", "default:3000/tcp", "debug:4000/tcp", "null", "null", "null"},
			},
		},
		{
			name: "multiple deployments",
			input: []components.DeploymentV3{
				{
					IdleTimeoutEnabled: true,
					RoomsPerProcess:    3,
					AdditionalContainerPorts: []components.ContainerPort{
						{
							TransportType: "tcp",
							Port:          4000,
							Name:          "debug",
						},
					},
					DefaultContainerPort: components.ContainerPort{
						TransportType: "tcp",
						Port:          3000,
						Name:          "default",
					},
					CreatedAt:    ts,
					CreatedBy:    "createdBy",
					AppID:        "appID",
					DeploymentID: "dep-2",
					BuildID:      "bld-1",
					Env: []components.DeploymentV3Env{
						{
							Name:  "EULA",
							Value: "TRUE",
						},
					},
					RequestedMemoryMB: 1024,
					RequestedCPU:      0.5,
				},
				{
					IdleTimeoutEnabled: true,
					RoomsPerProcess:    3,
					AdditionalContainerPorts: []components.ContainerPort{
						{
							TransportType: "tcp",
							Port:          4000,
							Name:          "debug",
						},
					},
					DefaultContainerPort: components.ContainerPort{
						TransportType: "tcp",
						Port:          3000,
						Name:          "default",
					},
					CreatedAt:    ts,
					CreatedBy:    "createdBy",
					AppID:        "appID",
					DeploymentID: "dep-2",
					BuildID:      "bld-1",
					Env: []components.DeploymentV3Env{
						{
							Name:  "EULA",
							Value: "TRUE",
						},
					},
					RequestedMemoryMB: 1024,
					RequestedCPU:      0.5,
				},
			},
			expect: [][]string{
				{"DeploymentID", "BuildID", "CreatedAt", "IdleTimeoutEnabled", "RoomsPerProcess", "RequestedCPU", "RequestedMemory", "DefaultContainerPort", "AdditionalContainerPorts", "BuildTag", "DeploymentTag", "ExperimentalRequestedGPU"},
				{"dep-2", "bld-1", "2021-01-01T00:00:00Z", "true", "3", "0.5", "1.0", "GiB", "default:3000/tcp", "debug:4000/tcp", "null", "null", "null"},
				{"dep-2", "bld-1", "2021-01-01T00:00:00Z", "true", "3", "0.5", "1.0", "GiB", "default:3000/tcp", "debug:4000/tcp", "null", "null", "null"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			formatter := commands.BuildTextFormatter()
			var buf bytes.Buffer
			actualErr := formatter.Write(tt.input, &buf)
			assert.NoError(t, actualErr)
			actualStr := strings.TrimSpace(buf.String())
			actualLines := strings.Split(actualStr, "\n")
			for i, line := range actualLines {
				actualLines[i] = strings.TrimSpace(line)
				if actualLines[i] == "" {
					continue
				}
				actualColumns := strings.Fields(actualLines[i])
				assert.Equal(t, tt.expect[i], actualColumns)
			}
		})
	}
}

func Test_BuildTextOutput(t *testing.T) {
	ts := time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name   string
		input  any
		expect [][]string
	}{
		{
			name: "single build",
			input: components.BuildV3{
				CreatedAt:  ts,
				CreatedBy:  "createdBy",
				BuildID:    "bld-1",
				ImageSize:  2048,
				Status:     "status",
				BuildTag:   sdk.String("v1.0.0"),
				StartedAt:  &ts,
				FinishedAt: nil,
			},
			expect: [][]string{
				{"BuildID", "BuildTag", "CreatedAt", "Status", "ImageSize", "StartedAt", "FinishedAt", "ContentHash", "OrgID"},
				{"bld-1", "v1.0.0", "2021-01-01T00:00:00Z", "status", "2.0", "KiB", "2021-01-01T00:00:00Z", "null", "null"},
			},
		},
		{
			name: "single build ptr",
			input: &components.BuildV3{
				CreatedAt:  ts,
				CreatedBy:  "createdBy",
				BuildID:    "bld-1",
				ImageSize:  2048,
				Status:     "status",
				BuildTag:   sdk.String("v1.0.0"),
				StartedAt:  &ts,
				FinishedAt: nil,
			},
			expect: [][]string{
				{"BuildID", "BuildTag", "CreatedAt", "Status", "ImageSize", "StartedAt", "FinishedAt", "ContentHash", "OrgID"},
				{"bld-1", "v1.0.0", "2021-01-01T00:00:00Z", "status", "2.0", "KiB", "2021-01-01T00:00:00Z", "null", "null"},
			},
		},
		{
			name: "multiple builds",
			input: []components.BuildV3{
				{
					CreatedAt:  ts,
					CreatedBy:  "createdBy",
					BuildID:    "bld-1",
					ImageSize:  2048,
					Status:     "status",
					BuildTag:   sdk.String("v1.0.0"),
					StartedAt:  &ts,
					FinishedAt: nil,
				},
				{
					CreatedAt:  ts,
					CreatedBy:  "createdBy",
					BuildID:    "bld-1",
					ImageSize:  2048,
					Status:     "status",
					BuildTag:   sdk.String("v1.0.0"),
					StartedAt:  &ts,
					FinishedAt: nil,
				},
			},
			expect: [][]string{
				{"BuildID", "BuildTag", "CreatedAt", "Status", "ImageSize", "StartedAt", "FinishedAt", "ContentHash", "OrgID"},
				{"bld-1", "v1.0.0", "2021-01-01T00:00:00Z", "status", "2.0", "KiB", "2021-01-01T00:00:00Z", "null", "null"},
				{"bld-1", "v1.0.0", "2021-01-01T00:00:00Z", "status", "2.0", "KiB", "2021-01-01T00:00:00Z", "null", "null"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			formatter := commands.BuildTextFormatter()
			var buf bytes.Buffer
			actualErr := formatter.Write(tt.input, &buf)
			assert.NoError(t, actualErr)
			actualStr := strings.TrimSpace(buf.String())
			actualLines := strings.Split(actualStr, "\n")
			for i, line := range actualLines {
				actualLines[i] = strings.TrimSpace(line)
				if actualLines[i] == "" {
					continue
				}
				actualColumns := strings.Fields(actualLines[i])
				assert.Equal(t, tt.expect[i], actualColumns)
			}
		})
	}
}
