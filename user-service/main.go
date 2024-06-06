package main

import (
	"shop/user-service/configs"
	"shop/user-service/src/controllers"
	"shop/user-service/src/models"

	"github.com/gin-gonic/gin"
)

func main() {
	db := configs.InitDB()
	db.AutoMigrate(models.User{})

	router := gin.Default()

	router.POST("/register", controllers.Register)

	router.Run("0.0.0.0:8081")
}
