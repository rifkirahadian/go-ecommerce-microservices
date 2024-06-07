package dtos

type CreateWarehouseDto struct {
	Name string `json:"name" binding:"required"`
}
