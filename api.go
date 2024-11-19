package main

import (
	"log"
	"os"

	"github.com/aramirez3/projects/internal/database"
	"github.com/joho/godotenv"
)

type apiConfig struct {
	DB   *database.Queries
	Addr string
}

func newAPIConfig() *apiConfig {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("error loading .env file: %v\n", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
		log.Fatal("PORT environment variable is not set. Default set to %s", port)
	}
	return &apiConfig{
		Addr: ":" + port,
	}
}