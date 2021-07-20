package handlers

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/tuan-krabs-github/go-oauth-example/entities"
)

func LoginHandler(c echo.Context) (err error) {
	l := new(entities.Login)

	if err := c.Bind(l); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(l); err != nil {
		return err
	}

	if l.Username != "adiatma" || l.Password != "secret" {
		return echo.ErrUnauthorized
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		ExpiresAt: 1500,
	})

	t, err := token.SignedString([]byte("secret"))

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"result": echo.Map{
			"user":  l,
			"token": t,
		},
	})
}
