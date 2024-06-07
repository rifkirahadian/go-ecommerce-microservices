package dtos

type User struct {
	Email string `json:"email"`
	Exp   int64  `json:"exp"`
	ID    uint   `json:"id"`
	Name  string `json:"name"`
}
