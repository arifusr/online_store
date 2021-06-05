package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/arifusr/online_store/app"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func JWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		App := c.Get("App").(*app.App)
		authorizationHeader := c.Request().Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			c.NoContent(http.StatusUnauthorized)
			return errors.New("no bearer token")
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			return []byte(App.Config.ConfigJWT.Secret), nil

		})
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.NoContent(http.StatusUnauthorized)
			return errors.New("token not valid")
		}
		c.Set("UserId", int(claims["user_id"].(float64)))
		return next(c)
	}
}
