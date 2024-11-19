package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/aramirez3/projects/internal/database"
	"github.com/labstack/echo/v4"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func main() {

	apiConfig := newAPIConfig()

	fmt.Println("hello world")
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		log.Println("Unable to connect to the db.")
		log.Println("DATABASE_URL environment variable is not set")
	} else {
		db, err := sql.Open("libsql", dbUrl)
		if err != nil {
			log.Fatal(err)
		}
		dbQueries := database.New(db)
		apiConfig.DB = dbQueries
		log.Println("Connected to db")

		e := echo.New()
		e.File("/", "public/index.html")
		e.Static("/", "public")
		e.GET("/healthz", handlerHealth)
		e.POST("/admin/reset", apiConfig.handlerReset)
		e.POST("/api/users", apiConfig.handlerCreateUser)
		e.POST("/api/login", apiConfig.handlerLogin)
		e.Logger.Fatal(e.Start(apiConfig.Addr))
	}
}
