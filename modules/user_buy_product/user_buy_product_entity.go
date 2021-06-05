package user_buy_product

type UserBuyProductCreateRequest struct {
	ProductId int `json:"product_id" validate:"required"`
	Qty       int `json:"qty" validate:"required"`
}
