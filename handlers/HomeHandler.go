package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HomeHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"login":   "/login",
		"profile": "/api/profile",
	})
}
