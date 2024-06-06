package configs

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	var err error
	db, err := gorm.Open(sqlite.Open("user-service.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
