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

	parts := strings.Split(s, "=")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid env var format: %s", s)
	}

	return &shared.DeploymentConfigV3Env{
		Name:  strings.TrimSpace(parts[0]),
		Value: strings.TrimSpace(parts[1]),
	}, nil
}

func MapEnvToEnvConfig(input []shared.DeploymentV3Env) []shared.DeploymentConfigV3Env {
	output := make([]shared.DeploymentConfigV3Env, 0)

	for _, config := range input {
		output = append(output, shared.DeploymentConfigV3Env(config))
	}

	return output
}
