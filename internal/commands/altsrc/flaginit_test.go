package altsrc

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v3"
)

func Test_InitializeFlagSources(t *testing.T) {
	t.Parallel()
	tests := []struct {
		skip         string
		name         string
		defaultValue string
		sources      []cli.ValueSource
		args         []string
		envVars      map[string]string
		expectValue  any
	}{
		{
			name: "file flag beats env var when first in chain",
			args: []string{},
			sources: []cli.ValueSource{
				&mockFlagInitializedValueSource{
					value: "file",
				},
				cli.EnvVar("TEST"),
			},
			envVars: map[string]string{
				"TEST": "ENV_VAR",
			},
			expectValue: "file",
		},
		{
			name: "env var beats file flag when first in chain",
			args: []string{},
			sources: []cli.ValueSource{
				cli.EnvVar("TEST"),
				&mockFlagInitializedValueSource{
					value: "file",
				},
			},
			envVars: map[string]string{
				"TEST": "ENV_VAR",
			},
			expectValue: "ENV_VAR",
		},
		{
			name: "passed arg wins when file flag is first in chain",
			args: []string{"--flag", "arg"},
			sources: []cli.ValueSource{
				&mockFlagInitializedValueSource{
					value: "file",
				},
				cli.EnvVar("TEST"),
			},
			envVars: map[string]string{
				"TEST": "ENV_VAR",
			},
			expectValue: "arg",
		},
		{
			name: "passed arg wins by alias when file flag is first in chain",
			args: []string{"-f", "arg"},
			sources: []cli.ValueSource{
				&mockFlagInitializedValueSource{
					value: "file",
				},
				cli.EnvVar("TEST"),
			},
			envVars: map[string]string{
				"TEST": "ENV_VAR",
			},
			expectValue: "arg",
		},
		{
			name: "passed arg wins when env var is first in chain",
			args: []string{"--flag", "arg"},
			sources: []cli.ValueSource{
				cli.EnvVar("TEST"),
				&mockFlagInitializedValueSource{
					value: "file",
				},
			},
			envVars: map[string]string{
				"TEST": "ENV_VAR",
			},
			expectValue: "arg",
		},
		{
			name: "passed arg wins by alias when env var is first in chain",
			args: []string{"-f", "arg"},
			sources: []cli.ValueSource{
				cli.EnvVar("TEST"),
				&mockFlagInitializedValueSource{
					value: "file",
				},
			},
			envVars: map[string]string{
				"TEST": "ENV_VAR",
			},
			expectValue: "arg",
		},
		{
			name: "passed arg wins when env var is last in chain and matches passed arg",
			args: []string{"--flag", "arg"},
			sources: []cli.ValueSource{
				&mockFlagInitializedValueSource{
					value: "file",
				},
				cli.EnvVar("TEST"),
			},
			envVars: map[string]string{
				"TEST": "arg",
			},
			expectValue: "arg",
		},
		{
			name: "passed arg wins by alias when env var is last in chain and matches passed arg",
			args: []string{"-f", "arg"},
			sources: []cli.ValueSource{
				&mockFlagInitializedValueSource{
					value: "file",
				},
				cli.EnvVar("TEST"),
			},
			envVars: map[string]string{
				"TEST": "arg",
			},
			expectValue: "arg",
		},
		{
			name: "default value wins when nothing in chain supplied",
			args: []string{},
			sources: []cli.ValueSource{
				&mockFlagInitializedValueSource{
					value: "",
				},
			},
			defaultValue: "default",
			expectValue:  "default",
		},
		{
			name: "file wins when flag has default but no passed value",
			args: []string{},
			sources: []cli.ValueSource{
				&mockFlagInitializedValueSource{
					value: "file",
				},
			},
			defaultValue: "default",
			expectValue:  "file",
		},
		{
			name: "passed arg wins when flag has default, file is set, and arg is set to default",
			args: []string{"--flag", "default"},
			sources: []cli.ValueSource{
				&mockFlagInitializedValueSource{
					value: "file",
				},
			},
			defaultValue: "default",
			expectValue:  "default",
		},
		{
			name: "passed arg wins when flag has default, file is set, and arg is set to default by alias",
			args: []string{"-f", "default"},
			sources: []cli.ValueSource{
				&mockFlagInitializedValueSource{
					value: "file",
				},
			},
			defaultValue: "default",
			expectValue:  "default",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.skip != "" {
				t.Skip(tt.skip)
			}
			for k, v := range tt.envVars {
				currentValue := os.Getenv(k)
				t.Cleanup(func() {
					os.Setenv(k, currentValue)
				})
				os.Setenv(k, v)
			}
			var flagValue any
			cmd := &cli.Command{
				Name: "test",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "flag",
						Aliases: []string{"f"},
						Value:   tt.defaultValue,
						Sources: cli.NewValueSourceChain(tt.sources...),
					},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					t.Helper()

					err := InitializeValueSourcesFromFlags(ctx, cmd, tt.args)
					require.NoError(t, err)
					flagValue = cmd.Value("flag")

					return nil
				},
			}

			err := cmd.Run(context.Background(), append([]string{"test"}, tt.args...))
			if err != nil {
				return 
			}

			assert.Equal(t, fmt.Sprintf("%v", tt.expectValue), flagValue)
		})
	}
}

type mockFlagInitializedValueSource struct {
	value       string
	initialized bool
}

func (m *mockFlagInitializedValueSource) Lookup() (string, bool) {
	if !m.initialized || m.value == "" {
		return "", false
	}
	return m.value, true
}

func (m *mockFlagInitializedValueSource) Initialize(cmd *cli.Command) error {
	m.initialized = true
	return nil
}

func (m *mockFlagInitializedValueSource) String() string {
	return fmt.Sprintf("mockFlagInitializedValueSource{value:%q}", m.value)
}

func (m *mockFlagInitializedValueSource) GoString() string {
	return fmt.Sprintf("&mockFlagInitializedValueSource{value:%q,initialized:%t}", m.value, m.initialized)
}
