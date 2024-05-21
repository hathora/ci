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
