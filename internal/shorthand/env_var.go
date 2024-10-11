package shorthand

import (
	"fmt"
	"strings"

	"github.com/hathora/ci/internal/sdk/models/shared"
)

func ParseDeploymentEnvVar(s string) (*shared.DeploymentConfigV3Env, error) {
	if s == "" {
		return nil, fmt.Errorf("env var cannot be empty")
	}

	parts := strings.SplitN(s, "=", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid env var format: %s", s)
	}

	key := strings.TrimSpace(parts[0])
	value := strings.TrimSpace(parts[1])
	value = TrimSingleOuterQuote(value)

	return &shared.DeploymentConfigV3Env{
		Name:  key,
		Value: value,
	}, nil
}

func TrimSingleOuterQuote(s string) string {
	if len(s) > 0 {
		lastChar := s[len(s)-1]
		if lastChar == '"' {
			s = strings.Trim(s, `"`)
		}
		if lastChar == '\'' {
			s = strings.Trim(s, `'`)
		}
	}

	return s
}

func MapEnvToEnvConfig(input []shared.DeploymentV3Env) []shared.DeploymentConfigV3Env {
	output := make([]shared.DeploymentConfigV3Env, 0)

	for _, config := range input {
		output = append(output, shared.DeploymentConfigV3Env(config))
	}

	return output
}
