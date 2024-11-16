package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aramirez3/projects/internal/database"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type apiConfig struct {
	DB   *database.Queries
	Addr string
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("error loading .env file: %v\n", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	apiConfig := apiConfig{
		Addr: ":" + port,
	}

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
		e.File("/", "html/index.html")
		e.Static("/", "html")
		e.GET("/healthz", handlerHealth)
		e.POST("/api/users", apiConfig.handlerCreateUser)
		e.POST("/admin/reset", apiConfig.handlerReset)
		e.Logger.Fatal(e.Start(apiConfig.Addr))
	}
}

func (cfg *apiConfig) handlerReset(c echo.Context) error {
	if err := cfg.DB.DeleteUsers(c.Request().Context()); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
	return c.String(http.StatusNoContent, http.StatusText(http.StatusNoContent))
}
