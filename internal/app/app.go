package app

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"

	"github.com/sletkov/effective-mobile-test-task/config"
	v1 "github.com/sletkov/effective-mobile-test-task/internal/controller/http/v1"
	"github.com/sletkov/effective-mobile-test-task/internal/repository/postgres"
	"github.com/sletkov/effective-mobile-test-task/internal/service"
	httptransport "github.com/sletkov/effective-mobile-test-task/internal/transport/http"
)

// Run application
func Run(configPath string, makeMigrations, dropMigrations bool) error {

	// Initialize logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	slog.SetDefault(logger)

	// Read config from .env
	slog.Debug("reading config...")

	var config config.Config

	err := cleanenv.ReadConfig(configPath, &config)

	if err != nil {
		return fmt.Errorf("reading config: %w", err)
	}

	// Initialize database
	slog.Debug("initializing database...")
	db, err := initDB(config.DatabaseURL)

	if err != nil {
		return fmt.Errorf("initializing db: %w", err)
	}

	// Migrations control
	if makeMigrations {
		// Make migrations
		slog.Debug("making migrations...")

		if err := goose.Up(db, "migrations"); err != nil {
			return fmt.Errorf("making migrations: %w", err)
		}
	} else if dropMigrations {
		// Rollback migrations
		slog.Debug("rollback migrations...")

		if err := goose.Down(db, "migrations"); err != nil {
			return fmt.Errorf("making migrations: %w", err)
		}
	}

	repo := postgres.New(db)

	transport := httptransport.New(http.DefaultClient)

	service := service.New(repo, transport)

	controller := v1.New(service)

	router := controller.InitRoutes(context.Background())

	slog.Debug("starting server...")

	return http.ListenAndServe(net.JoinHostPort(config.Host, config.Port), router)
}

// Initialize postgres database
func initDB(url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)

	if err != nil {
		return nil, err
	}

	return db, nil
}
