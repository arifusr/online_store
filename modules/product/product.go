package product

import (
	"github.com/arifusr/online_store/app"
)

const Name string = "product"

type Product struct {
}

func (product *Product) Install(a *app.App) error {
	svc := app.NewService(Name, a, ProductFindAll, nil, ProductCreate)
	if err := a.RegisterService(svc); err != nil {
		return err
	}
	if err := a.RegisterModel(&ProductModel{}); err != nil {
		return err
	}
	return nil
}
