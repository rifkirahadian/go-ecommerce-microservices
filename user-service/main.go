package main

import (
	"shop/user-service/configs"
	"shop/user-service/src/controllers"
	"shop/user-service/src/middlewares"
	"shop/user-service/src/models"

	"github.com/gin-gonic/gin"
)

func main() {
	db := configs.InitDB()
	db.AutoMigrate(models.User{})

	router := gin.Default()

	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.GET("/auth/user", middlewares.JWTMiddleware(), controllers.UserAuth)

	router.Run("0.0.0.0:8081")
}
