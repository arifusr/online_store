package user_buy_product

import (
	"fmt"

	"github.com/arifusr/online_store/modules/validator"
	"github.com/labstack/echo/v4"
)

func UserBuyProductCreate(c echo.Context) error {
	Validator := c.Get("Validator").(*validator.Validator)
	requestBody := UserBuyProductCreateRequest{}
	Validator.BindAndValidate(&requestBody, c)
	fmt.Print(requestBody)
	return nil
}
