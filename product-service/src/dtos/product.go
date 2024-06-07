package dtos

type CreateProductDto struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Stock       uint   `json:"stock" binding:"required"`
}
