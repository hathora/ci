package shorthand

import (
	"fmt"
	"strings"

	"github.com/hathora/ci/internal/sdk/models/shared"
)

func ParseDeploymentEnvVar(s string) (*shared.DeploymentConfigV2Env, error) {
	if s == "" {
		return nil, fmt.Errorf("env var cannot be empty")
	}

	parts := strings.Split(s, "=")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid env var format: %s", s)
	}

	return &shared.DeploymentConfigV2Env{
		Name:  strings.TrimSpace(parts[0]),
		Value: strings.TrimSpace(parts[1]),
	}, nil
}

func MapEnvToEnvConfig(input []shared.DeploymentV2Env) []shared.DeploymentConfigV2Env {
	output := make([]shared.DeploymentConfigV2Env, 0)

	for _, config := range input {
		newEnvVar := shared.DeploymentConfigV2Env{
			Name: config.Name,
			Value: config.Value,
		}

		output = append(output, newEnvVar)
	}

	return output
}