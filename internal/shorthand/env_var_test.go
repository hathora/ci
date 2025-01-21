package shorthand_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/hathora/ci/internal/shorthand"
	"github.com/hathora/cloud-sdk-go/models/components"
)

func Test_DeploymentEnvVarShorthand(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expect    *components.DeploymentConfigV3Env
		expectErr bool
	}{
		{
			name:  "name and value",
			input: "NAME=VALUE",
			expect: &components.DeploymentConfigV3Env{
				Name:  "NAME",
				Value: "VALUE",
			},
		},
		{
			name:  "name and value with spaces",
			input: "NAME = VALUE ",
			expect: &components.DeploymentConfigV3Env{
				Name:  "NAME",
				Value: "VALUE",
			},
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
		{
			name:  "nested flag",
			input: `KEY=-SomeFlag="With Spaces,And Commas"`,
			expect: &components.DeploymentConfigV3Env{
				Name:  "KEY",
				Value: `-SomeFlag="With Spaces,And Commas"`,
			},
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

func TestParseConfigFileVars(t *testing.T) {
	input := `[KEY1=VAL1 KEY3="VAL 3 WITH SPACES" KEY2=VAL2]`
	out, err := shorthand.ParseConfigFileVars(input)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 3, len(out))
	input = `[KEY1 =SHOULD_ERROR]`
	_, err = shorthand.ParseConfigFileVars(input)
	if err == nil {
		t.Error("expected error for invalid env var format")
	}
}
