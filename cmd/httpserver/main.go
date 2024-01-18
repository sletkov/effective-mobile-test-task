package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sletkov/effective-mobile-test-task/config"
	"github.com/sletkov/effective-mobile-test-task/internal/app"
	_ "github.com/sletkov/effective-mobile-test-task/migrations"
)

const envPath = "../../.env"

func main() {
	// Load config from .env
	err := godotenv.Load(envPath)

	if err != nil {
		log.Fatalf("cannot load .env file: %s", err.Error())
	}

	host := os.Getenv("SERVER_HOST")
	port := os.Getenv("SERVER_PORT")
	dbURL := os.Getenv("DB_URL")

	config := config.New(host, port, dbURL)

	// Run server
	err = app.Run(config)

	if err != nil {
		log.Fatal(err)
	}
}
