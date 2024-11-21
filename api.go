package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aramirez3/projects/internal/auth"
	"github.com/aramirez3/projects/internal/database"
	"github.com/joho/godotenv"
)

type apiConfig struct {
	DB     *database.Queries
	Addr   string
	Secret string
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

	secret := os.Getenv("SECRET")
	if port == "" {
		log.Fatalf("SECRET environment variable is not set. Default set to %s", port)
	}

	return &apiConfig{
		Addr:   ":" + port,
		Secret: secret,
	}
}

func (cfg *apiConfig) StartServer() {
	router := http.NewServeMux()
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: RequestLoggerMiddleware(router),
	}

	router.Handle("/", http.FileServer(http.Dir("public/")))
	router.HandleFunc("/healthz", handlerHealth)

	fmt.Printf("Server started at %s\n", cfg.Addr)
	server.ListenAndServe()
}

func RequestLoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v \"%v\"", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}

func (cfg *apiConfig) AuthRequiredMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := auth.GetBearerToken(r.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		}
		tokenUserID, err := auth.ValidateJWT(token, cfg.Secret)
	}
}
