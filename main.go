package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/aramirez3/projects/internal/database"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

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

	apiConfig.StartServer()
	// e := echo.New()
	// e.File("/", "public/index.html")
	// e.Static("/", "public")
	// e.GET("/healthz", handlerHealth)
	// e.POST("/admin/reset", apiConfig.handlerReset)
	// e.POST("/api/users", apiConfig.handlerCreateUser)
	// e.POST("/api/login", apiConfig.handlerLogin)
	// e.Logger.Fatal(e.Start(apiConfig.Addr))
}
