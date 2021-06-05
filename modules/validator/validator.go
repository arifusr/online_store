package validator

import (
	"github.com/arifusr/online_store/app"
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func (v *Validator) Install(a *app.App) error {
	v.validator = validator.New()
	a.Echo.Validator = v
	return nil
}
