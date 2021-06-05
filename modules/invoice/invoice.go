package invoice

import (
	"github.com/arifusr/online_store/app"
)

const Name string = "invoice"

type Invoice struct {
}

func (inv *Invoice) Install(a *app.App) error {
	svc := app.NewService(Name, a, nil, nil, nil)
	if err := a.RegisterService(svc); err != nil {
		return err
	}
	return nil
}
