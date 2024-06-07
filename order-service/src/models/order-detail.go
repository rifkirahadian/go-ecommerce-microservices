package models

import "time"

type OrderDetail struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	ProductId uint      `gorm:"column:product_id"`
	OrderId   uint      `gorm:"column:order_id"`
	Quantity  uint      `gorm:"column:quantity"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}
