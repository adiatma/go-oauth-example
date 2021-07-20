package main

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tuan-krabs-github/go-oauth-example/handlers"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}

func main() {
	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", handlers.HomeHandler)        // public
	e.POST("/login", handlers.LoginHandler) // public

	r := e.Group("/api") // group path "api/*" need auth
	config := middleware.JWTConfig{
		SigningKey:  []byte("secret"),
		TokenLookup: "header:" + echo.HeaderAuthorization,
		AuthScheme:  "Bearer",
	}
	r.Use(middleware.JWTWithConfig(config))
	r.GET("/profile", handlers.ProfileHandler) // api/profile

	e.Logger.Fatal(e.Start(":1323"))
}
