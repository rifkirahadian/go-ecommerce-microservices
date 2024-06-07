package dtos

type CreateWarehouseDto struct {
	Name string `json:"name" binding:"required"`
}

type UpdateWarehouseStatusDto struct {
	WarehouseId uint `json:"warehouse_id" binding:"required"`
	IsActive    bool `json:"is_active"`
}
