package models

import "time"

type Order struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Code      string    `gorm:"column:code"`
	Status    string    `gorm:"column:status"`
	Total     float32   `gorm:"column:user_id"`
	UserId    uint      `gorm:"column:user_id"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}
