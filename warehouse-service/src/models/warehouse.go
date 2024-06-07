package models

import "time"

type Warehouse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `gorm:"column:name"`
	UserId    uint      `gorm:"column:user_id"`
	IsActive  bool      `gorm:"column:is_active"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}
