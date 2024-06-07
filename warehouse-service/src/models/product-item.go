package models

import "time"

type ProductItem struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	ProductId   uint      `gorm:"column:product_id"`
	Code        string    `gorm:"column:code"`
	UserId      uint      `gorm:"column:user_id"`
	IsAvailable bool      `gorm:"column:is_available"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at"`
}
