package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/aramirez3/projects/internal/auth"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type authUser struct {
	Id        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	Updated   time.Time `json:"updated_at"`
}

func (cfg *apiConfig) handlerLogin(c echo.Context) error {
	type payload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	body := payload{}
	if err := c.Bind(body); err != nil {
		return c.String(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}
	fmt.Printf("Payload had been validated. This will query the db: %v\n", body)
	dbUser, err := cfg.DB.GetUserByEmail(c.Request().Context(), body.Email)
	if err != nil || dbUser.HashedPassword == "" {
		fmt.Println(err)
		return c.String(http.StatusNotFound, "User email was not found")
	}
	err = auth.CheckPasswordHash(body.Password, "")

	return c.String(http.StatusOK, http.StatusText(http.StatusOK))
}
