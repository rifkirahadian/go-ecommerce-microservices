package dtos

type CreateProductItemDto struct {
	ProductId uint `json:"product_id" binding:"required"`
	Count     uint `json:"count" binding:"required"`
}
