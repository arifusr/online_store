package user_login

import (
	"net/http"

	"github.com/arifusr/online_store/modules/validator"
	"github.com/labstack/echo/v4"
)

func GenerateToken(c echo.Context) error {
	userLoginRequest := UserLoginRequest{}
	Validator := c.Get("Validator").(*validator.Validator)
	if err := Validator.BindAndValidate(&userLoginRequest, c); err != nil {
		return err
	}
	c.JSON(http.StatusOK, userLoginRequest)
	return nil
}
