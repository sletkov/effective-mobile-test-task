package app

import (
	"context"
	"database/sql"
	"log/slog"
	"net"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"

	"github.com/sletkov/effective-mobile-test-task/config"
	v1 "github.com/sletkov/effective-mobile-test-task/internal/controller/http/v1"
	"github.com/sletkov/effective-mobile-test-task/internal/repository/postgres"
	"github.com/sletkov/effective-mobile-test-task/internal/service"
	httptransport "github.com/sletkov/effective-mobile-test-task/internal/transport/http"
)

// Run application
func Run(config *config.Config) error {
	db, err := initDB(config.DatabaseURL)

	if err != nil {
		return err
	}

	// Up migrations
	if err := goose.Up(db, "../../migrations"); err != nil {
		return err
	}

	repo := postgres.New(db)

	transport := httptransport.New(http.DefaultClient)

	service := service.New(repo, transport)

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	controller := v1.New(*service, *logger)

	router := controller.InitRoutes(context.Background())

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
