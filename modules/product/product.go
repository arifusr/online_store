package product

import (
	"github.com/arifusr/online_store/app"
	"github.com/arifusr/online_store/middleware"
)

const Name string = "product"

type Product struct {
}

func (product *Product) Install(a *app.App) error {
	svc := app.NewService(Name, ProductFindAll, nil, ProductCreate)
	svc.AddMiddlewareCreate(middleware.JWT)
	svc.AddContextValue("Db", a.Db)
	if err := a.RegisterService(svc); err != nil {
		return err
	}
	if err := a.RegisterModel(&ProductModel{}); err != nil {
		return err
	}
	return nil
}
