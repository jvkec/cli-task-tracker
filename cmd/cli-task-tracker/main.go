package main

import (
	"fmt"
	"os"

	"github.com/jvkec/cli-task-tracker/internal/app"
)

func main() {
	application := app.New()

	if err := application.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
