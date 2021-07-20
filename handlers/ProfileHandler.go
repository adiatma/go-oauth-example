package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ProfileHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"result": echo.Map{
			"message": "Ok",
		},
	})
}
