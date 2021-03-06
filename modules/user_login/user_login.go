package user_login

import (
	"github.com/arifusr/online_store/app"
)

const Name string = "user_login"

type UserLogin struct {
}

func (userLogin *UserLogin) Install(a *app.App) error {
	svc := app.NewService(Name, nil, nil, GenerateToken)
	svc.AddContextValue("Db", a.Db)
	if err := a.RegisterService(svc); err != nil {
		return err
	}
	return nil
}
