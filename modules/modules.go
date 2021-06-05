package modules

import (
	"github.com/arifusr/online_store/app"
	"github.com/arifusr/online_store/modules/database"
	"github.com/arifusr/online_store/modules/env"
	"github.com/arifusr/online_store/modules/product"
	"github.com/arifusr/online_store/modules/user"
	"github.com/arifusr/online_store/modules/user_buy_product"
	"github.com/arifusr/online_store/modules/user_login"
	"github.com/arifusr/online_store/modules/validator"
)

type Modules struct {
}

func (module *Modules) Install(a *app.App) error {
	// install env
	Env := env.Env{}
	if err := a.Use(&Env); err != nil {
		return err
	}

	// install database sqlite

	Db := database.DatabaseSqlite{}
	if err := a.Use(&Db); err != nil {
		return err
	}

	// install database postgres
	// Db := database.DatabasePostgres{}
	// if err := a.Use(&Db); err != nil {
	// 	return err
	// }

	//install product
	Product := product.Product{}
	if err := a.Use(&Product); err != nil {
		return err
	}

	//install user
	User := user.User{}
	if err := a.Use(&User); err != nil {
		return err
	}

	//install user_login
	UserLogin := user_login.UserLogin{}
	if err := a.Use(&UserLogin); err != nil {
		return err
	}

	//install user_buy_product
	UserBuyProduct := user_buy_product.UserBuyProduct{}
	if err := a.Use(&UserBuyProduct); err != nil {
		return err
	}

	// install validator
	Validator := validator.Validator{}
	if err := a.Use(&Validator); err != nil {
		return err
	}

	return nil

}
