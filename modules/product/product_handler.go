package product

import (
	"net/http"

	"github.com/arifusr/online_store/app"

	"github.com/labstack/echo/v4"
)

func ProductFindAll(c echo.Context) error {
	return nil
}

func ProductCreate(c echo.Context) error {
	App := c.Get("App").(*app.App)
	App.Db.Create(&ProductModel{Name: "yakult", Price: 8000, Stock: 5})
	respose := struct {
		Status string `json:"status"`
	}{Status: "ok"}
	c.JSON(http.StatusOK, respose)
	return nil
}
