package user_buy_product

import (
	"errors"
	"net/http"

	"github.com/arifusr/online_store/modules/product"
	"github.com/arifusr/online_store/modules/user"
	"github.com/arifusr/online_store/modules/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RaceCondition(Db *gorm.DB, requestBody UserBuyProductCreateRequest) error {
	// get product
	Product := product.ProductModel{}
	Db.Model(product.ProductModel{}).First(&Product)

	// get user
	User := user.UserModel{}
	Db.Model(user.UserModel{}).First(&User)

	// naikkan versi product dan user

	newProductStock := Product.Stock - requestBody.Qty
	NewProductVersion := Product.Version + 1

	newUserSaldo := User.Saldo - (Product.Price * requestBody.Qty)
	newUserVersion := User.Version + 1

	// update menggunakan transaction

	tx := Db.Begin()

	if result := tx.Model(product.ProductModel{}).Where("id =?", Product.ID).Where("version =?", Product.Version).Updates(product.ProductModel{
		Stock:   newProductStock,
		Version: NewProductVersion,
	}); result.RowsAffected < 1 || result.Error != nil {
		tx.Rollback()
		// ulangi dari awal

		return errors.New("any error")
	}

	if result := tx.Model(user.UserModel{}).Where("id =?", User.ID).Where("version =?", User.Version).Updates(user.UserModel{
		Saldo:   newUserSaldo,
		Version: newUserVersion,
	}); result.RowsAffected < 1 || result.Error != nil {
		tx.Rollback()
		// ulangi dari awal
		return errors.New("any error")

	}

	if err := tx.Commit().Error; err != nil {
		// ulangi dari awal
		return errors.New("any error")
	}
	return nil
}

func UserBuyProductCreate(c echo.Context) error {
	Validator := c.Get("Validator").(*validator.Validator)
	Db := c.Get("Db").(*gorm.DB)
	// UserId := c.Get("UserId")
	requestBody := UserBuyProductCreateRequest{}
	Validator.BindAndValidate(&requestBody, c)

	if err := RaceCondition(Db, requestBody); err != nil {
		RaceCondition(Db, requestBody)
	}

	c.NoContent(http.StatusOK)
	return nil
}
