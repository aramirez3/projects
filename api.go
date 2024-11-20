package main

import (
	"log"
	"net/http"
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
		port = "3000"
		log.Fatalf("PORT environment variable is not set. Default set to %s", port)
	}
	return &apiConfig{
		Addr: ":" + port,
	}
}

func RequestLoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v \"%v\"", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}
