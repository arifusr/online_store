package user_login

import (
	"fmt"
	"net/http"
	"time"

	"github.com/arifusr/online_store/app"
	"github.com/arifusr/online_store/modules/validator"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type MyClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func GenerateToken(c echo.Context) error {
	userLoginRequest := UserLoginRequest{}
	Validator := c.Get("Validator").(*validator.Validator)
	App := c.Get("App").(*app.App)
	if err := Validator.BindAndValidate(&userLoginRequest, c); err != nil {
		return err
	}

	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    App.Config.Name,
			ExpiresAt: time.Now().Add(time.Duration(1) * time.Hour).Unix(),
		},
		UserId: 1,
	}
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)
	signedToken, err := token.SignedString([]byte(App.Config.ConfigJWT.Secret))
	fmt.Print(err)
	c.JSON(http.StatusOK, signedToken)
	return nil
}
