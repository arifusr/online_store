package user_login

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GenerateToken(c echo.Context) error {
	userLoginRequest := UserLoginRequest{}
	if err := c.Bind(&userLoginRequest); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return err
	}
	if err := c.Validate(userLoginRequest); err != nil {
		c.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
		}{err.Error()})
		return err
	}
	c.JSON(http.StatusOK, userLoginRequest)
	return nil
}
