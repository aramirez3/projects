package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (cfg *apiConfig) handlerReset(c echo.Context) error {
	if err := cfg.DB.DeleteUsers(c.Request().Context()); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
	return c.String(http.StatusNoContent, http.StatusText(http.StatusNoContent))
}
