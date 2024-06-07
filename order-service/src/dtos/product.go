package dtos

import "time"

type Product struct {
	ID          uint      `json:"id"`
	Name        string    `json:"Name"`
	Description string    `json:"Description"`
	UserID      int       `json:"UserId"`
	Price       float64   `json:"Price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
