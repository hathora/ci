package shorthand_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/hathora/ci/internal/sdk/models/components"
	"github.com/hathora/ci/internal/shorthand"
)

func Test_ContainerPortShorthand(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     string
		expect    *components.ContainerPort
		expectErr bool
	}{
		{
			name:  "port only",
			input: "8080",
			expect: &components.ContainerPort{
				TransportType: components.TransportTypeTCP,
				Port:          8080,
				Name:          "",
			},
		},
		{
			name:  "port with space",
			input: " 8080\t",
			expect: &components.ContainerPort{
				TransportType: components.TransportTypeTCP,
				Port:          8080,
				Name:          "",
			},
		},
		{
			name:  "port and transport",
			input: "8080/udp",
			expect: &components.ContainerPort{
				TransportType: components.TransportTypeUDP,
				Port:          8080,
				Name:          "",
			},
		},
		{
			name:  "port and transport with space",
			input: "\t8080/udp",
			expect: &components.ContainerPort{
				TransportType: components.TransportTypeUDP,
				Port:          8080,
				Name:          "",
			},
		},
		{
			name:  "port, transport, and name",
			input: "my-port:8080/udp",
			expect: &components.ContainerPort{
				TransportType: components.TransportTypeUDP,
				Port:          8080,
				Name:          "my-port",
			},
		},
		{
			name:  "port, transport, and name with space",
			input: "my-port:5555/tls\t",
			expect: &components.ContainerPort{
				TransportType: components.TransportTypeTLS,
				Port:          5555,
				Name:          "my-port",
			},
		},
		{
			name:  "port and name",
			input: "my-port:8080",
			expect: &components.ContainerPort{
				TransportType: components.TransportTypeTCP,
				Port:          8080,
				Name:          "my-port",
			},
		},
		{
			name:      "name only",
			input:     "my-port",
			expectErr: true,
		},
		{
			name:      "name and transport",
			input:     "my-port/udp",
			expectErr: true,
		},
		{
			name:      "transport only",
			input:     "/udp",
			expectErr: true,
		},
		{
			name:      "empty string",
			input:     "",
			expectErr: true,
		},
		{
			name:      "port too low",
			input:     "0",
			expectErr: true,
		},
		{
			name:      "port too high",
			input:     "65536",
			expectErr: true,
		},
		{
			name:      "invalid transport",
			input:     "8080/invalid",
			expectErr: true,
		},
		{
			name:      "invalid name",
			input:     "my-port!:8080/udp",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := shorthand.ParseContainerPort(tt.input)
			if tt.expectErr {
				assert.Nil(t, actual)
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expect, actual)
			}
		})
	}

}
