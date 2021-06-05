package product

type ProductCreateRequest struct {
	Name string `json:"name" validate:"required"`
}

const ProductTableName string = "product"

type ProductModel struct {
	ID    int    `gorm:"type:int;primaryKey;autoIncrement"`
	Name  string `gorm:"type:varchar(255);not null"`
	Stock int    `gorm:"type:int;not null;default:0"`
	Price int    `gorm:"type:int;not null;default:0"`
}

func (M *ProductModel) TableName() string {
	return ProductTableName
}
