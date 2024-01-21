package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/sletkov/effective-mobile-test-task/internal/app"
	_ "github.com/sletkov/effective-mobile-test-task/migrations"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "config path")

	flag.Parse()
}

func main() {
	// Run server
	err := app.Run(configPath)

	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
