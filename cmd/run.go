package main

import (
	"context"
	"fmt"
	"os"

	"github.com/hathora/ci/internal/commands"
)

func main() {
	if err := commands.App().Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
