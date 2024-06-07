package dtos

type CreateProductDto struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Stock       uint32 `json:"stock" binding:"required"`
}

type User struct {
	Email string `json:"email"`
	Exp   int64  `json:"exp"`
	ID    uint   `json:"id"`
	Name  string `json:"name"`
}
