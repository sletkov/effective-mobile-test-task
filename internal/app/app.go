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

	"github.com/sletkov/effective-mobile-test-task/internal/config"
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
	slog.Info("reading config")

	var config config.Config

	err := cleanenv.ReadConfig(configPath, &config)

	if err != nil {
		return fmt.Errorf("reading config: %w", err)
	}

	slog.Info("config was read successfully")

	// Initialize database
	slog.Info("initializing db")
	db, err := initDB(config.DatabaseURL)

	if err != nil {
		return fmt.Errorf("initializing db: %w", err)
	}

	slog.Info("db was initialized successfully")

	// Migrations control
	if makeMigrations {
		// Make migrations
		slog.Info("making migrations")

		if err := goose.Up(db, "migrations"); err != nil {
			return fmt.Errorf("making migrations: %w", err)
		}

		slog.Info("migrations were made successfully")

	} else if dropMigrations {
		// Rollback migrations
		slog.Info("rollback migrations")

		if err := goose.Down(db, "migrations"); err != nil {
			return fmt.Errorf("making migrations: %w", err)
		}

		slog.Info("migrations were rollbacked successfully")
	}

	repo := postgres.New(db)

	transport := httptransport.New(http.DefaultClient)

	service := service.New(repo, transport)

	controller := v1.New(service)

	router := controller.InitRoutes(context.Background())

	slog.Debug(fmt.Sprintf("http server started on port: %s", config.Port))

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
