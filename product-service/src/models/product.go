package models

import "time"

type Product struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `gorm:"column:name"`
	Description string    `gorm:"column:description"`
	Stock       uint32    `gorm:"column:stock"`
	UserId      uint      `gorm:"column:user_id"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at"`
}
