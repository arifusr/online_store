package invoice

import (
	"github.com/arifusr/online_store/modules/product"
	"github.com/arifusr/online_store/modules/user"
	"gorm.io/gorm"
)

const InvoiceTableName string = "invoice"

type InvoiceModel struct {
	ID         int `gorm:"type:int;primaryKey;autoIncrement"`
	UserID     int
	User       user.UserModel
	ProductID  int
	Product    product.ProductModel
	Qty        int
	StockAwal  int
	StockAkhir int
	gorm.Model
}

func (M *InvoiceModel) TableName() string {
	return InvoiceTableName
}
