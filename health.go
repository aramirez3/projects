package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func handlerHealth(c echo.Context) error {
	return c.String(http.StatusOK, http.StatusText(http.StatusOK))
}
