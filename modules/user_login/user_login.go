package user_login

import (
	"github.com/arifusr/online_store/app"
)

const Name string = "user_login"

type UserLogin struct {
}

func (userLogin *UserLogin) Install(a *app.App) error {
	svc := app.NewService(Name, a, nil, nil, GenerateToken)
	if err := a.RegisterService(svc); err != nil {
		return err
	}
	return nil
}
