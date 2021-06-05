package validator

import (
	"net/http"

	"github.com/arifusr/online_store/app"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func (v *Validator) BindAndValidate(i interface{}, c echo.Context) error {
	if err := c.Bind(i); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return err
	}
	if err := c.Validate(i); err != nil {
		c.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
		}{err.Error()})
		return err
	}
	return nil
}

func (v *Validator) Install(a *app.App) error {
	v.validator = validator.New()
	a.Echo.Validator = v
	for _, svc := range a.Services {
		svc.AddContextValue("Validator", v)
		svc.AddContextValue("App", a)
	}
	return nil
}
