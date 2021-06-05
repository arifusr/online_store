package product

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ProductFindAll(c echo.Context) error {
	return nil
}

func ProductCreate(c echo.Context) error {
	Db := c.Get("Db").(*gorm.DB)
	Db.Create(&ProductModel{Name: "yakult", Price: 8000, Stock: 5})
	respose := struct {
		Status string `json:"status"`
	}{Status: "ok"}
	c.JSON(http.StatusOK, respose)
	return nil
}
