package models

import "time"

type Product struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `gorm:"column:name"`
	Description string    `gorm:"column:description"`
	UserId      uint      `gorm:"column:user_id"`
	Price       float32   `gorm:"column:price"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at"`
}
