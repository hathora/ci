package main

import (
	"log"
	"os"

	"github.com/hathora/ci/internal/commands"
)

func main() {
	if err := commands.App().Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
