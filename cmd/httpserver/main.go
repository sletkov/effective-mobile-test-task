package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/sletkov/effective-mobile-test-task/internal/app"
	_ "github.com/sletkov/effective-mobile-test-task/migrations"
)

var (
	configPath     string
	makeMigrations bool
	dropMigrations bool
)

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "config path")
	flag.BoolVar(&makeMigrations, "make-migrations", false, "make migrations")
	flag.BoolVar(&dropMigrations, "drop-migrations", false, "rollback migrations")

	flag.Parse()
}

// @title HTTP server
// @version 1.0
// @description HTTP Server for saving users

// @host localhost:9999
// @BasePath /api/v1/users

func main() {
	// Run server
	err := app.Run(configPath, makeMigrations, dropMigrations)

	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
