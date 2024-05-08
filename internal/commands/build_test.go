package commands_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"

	"github.com/hathora/ci/internal/commands"
)

func Test_HelpText(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
	}{
		{"build help", "build"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := cli.NewApp()
			app.Commands = cli.Commands{commands.Build}
			err := app.Run([]string{"appname", tt.input})
			assert.Nil(t, err, "Commands should not return an error")

			// TODO check that the help text was printed
		})
	}
}
