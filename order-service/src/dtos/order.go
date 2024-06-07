package dtos

type ProductDto struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  uint `json:"quantity" binding:"required"`
}

type OrderDto struct {
	Products []ProductDto `json:"products" binding:"required"`
}
