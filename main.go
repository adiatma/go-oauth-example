package main

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	Login struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}

func HomeHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

func LoginHandler(c echo.Context) (err error) {
	l := new(Login)

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

func ProfileHandler(c echo.Context) error {
	// user := c.Get("user").(*jwt.Token)
	return c.JSON(http.StatusOK, echo.Map{
		"name": "test",
	})
}

func main() {
	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", HomeHandler)        // public
	e.POST("/login", LoginHandler) // public

	r := e.Group("/api") // group path "api/*" need authentication
	config := middleware.JWTConfig{
		SigningKey:  []byte("secret"),
		TokenLookup: "header:" + echo.HeaderAuthorization,
		AuthScheme:  "Bearer",
	}
	r.Use(middleware.JWTWithConfig(config))
	r.GET("/profile", ProfileHandler) // api/profile

	e.Logger.Fatal(e.Start(":1323"))
}
