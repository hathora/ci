package shorthand

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/hathora/ci/internal/sdk/models/components"
)

func ParseDeploymentEnvVar(s string) (*components.DeploymentConfigV3Env, error) {
	if s == "" {
		return nil, fmt.Errorf("env var cannot be empty")
	}

	parts := strings.SplitN(s, "=", 2)
	if len(parts) != 2 {
		return nil, NewErrInvalidEnvVarFormat(s)
	}

	parts[1] = strings.TrimSpace(parts[1])
	// remove any enclosing quotes
	if strings.HasPrefix(parts[1], `"`) && strings.HasSuffix(parts[1], `"`) {
		parts[1] = strings.TrimPrefix(strings.TrimSuffix(parts[1], `"`), `"`)
	}

	return &components.DeploymentConfigV3Env{
		Name:  strings.TrimSpace(parts[0]),
		Value: parts[1],
	}, nil
}

var reEnvs = regexp.MustCompile(`\b[A-Za-z0-9_]+="[^"]*"|\b[A-Za-z0-9_]+=[^" ]+`)

func ParseConfigFileVars(input string) ([]string, error) {
	input = strings.TrimPrefix(strings.TrimSuffix(input, "]"), "[")
	if len(input) == 0 {
		return nil, nil
	}
	matches := reEnvs.FindAllString(input, -1)
	if len(matches) == 0 {
		return nil, NewErrInvalidEnvVarFormat(input)
	}
	return matches, nil
}

func MapEnvToEnvConfig(input []components.DeploymentV3Env) []components.DeploymentConfigV3Env {
	output := make([]components.DeploymentConfigV3Env, 0)

	for _, config := range input {
		output = append(output, components.DeploymentConfigV3Env(config))
	}

	return output
}

func NewErrInvalidEnvVarFormat(env string) error {
	return fmt.Errorf("invalid env var format: %s", env)
}
