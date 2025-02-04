package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"

	"github.com/hathora/ci/internal/commands"
)

func main() {
	app := commands.App()
	args := commands.NormalizeArgs(app, os.Args)
	if err := app.Run(context.Background(), args); err != nil {
		red := color.New(color.FgRed)
		errStr := fmt.Sprintf("%v", err)
		errorLines := strings.Split(errStr, "\n")
		fmt.Println()
		fmt.Println()
		for _, line := range errorLines {
			red.Fprintf(os.Stderr, "[ERROR]\t%v\n", line)
		}
		os.Exit(1)
	}
}
