package user_buy_product

import (
	"github.com/arifusr/online_store/app"
	"github.com/arifusr/online_store/middleware"
)

const Name string = "user_buy_product"

type UserBuyProduct struct {
}

func (userBuyProduct *UserBuyProduct) Install(a *app.App) error {
	svc := app.NewService(Name, nil, nil, UserBuyProductCreate)
	svc.AddMiddlewareCreate(middleware.JWT)
	svc.AddContextValue("Db", a.Db)
	if err := a.RegisterService(svc); err != nil {
		return err
	}
	return nil
}
