package dtos

type CreateStockDto struct {
	ProductID uint `json:"product_id"`
	Count     uint `json:"count"`
}
