package commands

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v3"
)

func TestNormalizeArgs(t *testing.T) {
	cmd := &cli.Command{
		Flags: []cli.Flag{
			idleTimeoutFlag,
		},
	}
	tests := []struct {
		input, expected []string
	}{
		{
			input:    []string{"main", "--config", "abcd", "--idle-timeout-enabled", "false"},
			expected: []string{"main", "--config", "abcd", "--idle-timeout-enabled=false"},
		},
		{
			input:    []string{"main", "--config", "abcd", "--idle-timeout-enabled", "--next-flag"},
			expected: []string{"main", "--config", "abcd", "--idle-timeout-enabled", "--next-flag"},
		},
		{
			input:    []string{"main", "--config", "abcd", "--idle-timeout-enabled", "-n", "abc"},
			expected: []string{"main", "--config", "abcd", "--idle-timeout-enabled", "-n", "abc"},
		},
		{
			input:    []string{"main", "--config", "abcd", "--idle-timeout-enabled", "T", "-n", "abc"},
			expected: []string{"main", "--config", "abcd", "--idle-timeout-enabled=T", "-n", "abc"},
		},
		{
			input:    []string{"main", "--config", "abcd", "--idle-timeout-enabled"},
			expected: []string{"main", "--config", "abcd", "--idle-timeout-enabled"},
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := NormalizeArgs(cmd, test.input)
			assert.Equal(t, fmt.Sprintf("%v", got), fmt.Sprintf("%v", test.expected))
		})
	}
}
