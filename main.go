package main

import (
	"database/sql"
	"embed"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/aramirez3/projects/internal/database"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var staticFiles embed.FS

func main() {

	apiConfig := newAPIConfig()

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		log.Fatal("Unable to connect to the db.")
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	db, err := sql.Open("libsql", dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)
	apiConfig.DB = dbQueries
	log.Println("Connected to db")

	router := http.NewServeMux()
	server := http.Server{
		Addr:    apiConfig.Addr,
		Handler: RequestLoggerMiddleware(router),
	}
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		f, err := staticFiles.Open("public/index.html")
		if err != nil {
			fmt.Println("could not open public/index.html")
			log.Println(err)
			return
		}
		defer f.Close()
		if _, err := io.Copy(w, f); err != nil {
			log.Fatal(err)
			return
		}
	})

	fmt.Printf("Server started at %s\n", apiConfig.Addr)
	server.ListenAndServe()
	// e := echo.New()
	// e.File("/", "public/index.html")
	// e.Static("/", "public")
	// e.GET("/healthz", handlerHealth)
	// e.POST("/admin/reset", apiConfig.handlerReset)
	// e.POST("/api/users", apiConfig.handlerCreateUser)
	// e.POST("/api/login", apiConfig.handlerLogin)
	// e.Logger.Fatal(e.Start(apiConfig.Addr))
}
