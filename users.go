package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aramirez3/projects/internal/auth"
	"github.com/aramirez3/projects/internal/database"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type User struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

func (cfg *apiConfig) handlerCreateUser(c echo.Context) error {
	type payload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	body := new(payload)
	if err := StrictUnmarshal(body, c); err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	hash, err := auth.HashPassword(body.Password)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}
	params := database.CreateUserParams{
		ID:             uuid.New(),
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
		Email:          body.Email,
		HashedPassword: hash,
	}

	dbUser, err := cfg.DB.CreateUser(c.Request().Context(), params)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	fmt.Printf("user created: %v\n", dbUser.Email)

	return c.JSON(http.StatusCreated, dbUser)
}
