package dtos

type CreateProductItemDto struct {
	ProductId   uint `json:"product_id" binding:"required"`
	WarehouseId uint `json:"warehouse_id" binding:"required"`
	Count       uint `json:"count" binding:"required"`
}

type TransferStockDto struct {
	ProductItemId []uint `json:"product_item_id" binding:"required"`
	WarehouseId   uint   `json:"warehouse_id" binding:"required"`
}
