package user

import (
	"github.com/arifusr/online_store/app"
)

const Name string = "user"

type User struct {
}

func (user *User) Install(a *app.App) error {
	svc := app.NewService(Name, a, nil, nil, nil)
	if err := a.RegisterService(svc); err != nil {
		return err
	}
	if err := a.RegisterModel(&UserModel{}); err != nil {
		return err
	}
	return nil
}
