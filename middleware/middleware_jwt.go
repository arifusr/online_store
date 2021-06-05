package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func JWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorizationHeader := c.Request().Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			c.NoContent(http.StatusUnauthorized)
			return errors.New("no bearer token")
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
		jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			return nil, nil

		})
		return next(c)
	}
}
