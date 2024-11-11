package main

import (
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"hexagonal-architecture-example/internal/adapters/handlers/http"
	"hexagonal-architecture-example/internal/adapters/handlers/infrastructure"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	serve := http.Serve{
		Port: os.Getenv("PORT"),
		Postgres: infrastructure.Sql{
			Host:     os.Getenv("STORAGE_HOST"),
			Port:     os.Getenv("STORAGE_PORT"),
			User:     os.Getenv("STORAGE_USER"),
			Password: os.Getenv("STORAGE_PASSWORD"),
			DBName:   os.Getenv("STORAGE_DB"),
		},
	}

	http.ListenServe(serve)
}
