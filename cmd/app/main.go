// Package main is the entry point for the TUI application.
package main

import (
	"fmt"
	"os"

	"github.com/studerw/tui-e2e-demo/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
