package dtos

type ProductDto struct {
	ProductID int `json:"product_id" binding:"required"`
	Quantity  int `json:"quantity" binding:"required"`
}

type CreateProductsDto struct {
	Products []ProductDto `json:"products" binding:"required"`
}
