package main

import (
	"context"
	"log"
	"os"

	"github.com/hathora/ci/internal/commands"
)

func main() {
	if err := commands.App().Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
