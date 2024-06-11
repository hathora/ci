package shorthand_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/hathora/ci/internal/sdk/models/shared"
	"github.com/hathora/ci/internal/shorthand"
)

func Test_DeploymentEnvVarShorthand(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expect    *shared.DeploymentConfigV2Env
		expectErr bool
	}{
		{
			name:  "name and value",
			input: "NAME=VALUE",
			expect: &shared.DeploymentConfigV2Env{
				Name:  "NAME",
				Value: "VALUE",
			},
		},
		{
			name:  "name and value with spaces",
			input: "NAME = VALUE ",
			expect: &shared.DeploymentConfigV2Env{
				Name:  "NAME",
				Value: "VALUE",
			},
		},
		{
			name:      "name and value and extra",
			input:     "NAME=VALUE=EXTRA",
			expectErr: true,
		},
		{
			name:      "key only",
			input:     "KEY",
			expectErr: true,
		},
		{
			name:      "empty string",
			input:     "",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := shorthand.ParseDeploymentEnvVar(tt.input)
			if tt.expectErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expect, actual)
		})
	}

}
