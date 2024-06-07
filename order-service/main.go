package main

import (
	"shop/order-service/configs"
	"shop/order-service/src/middlewares"
	"shop/order-service/src/models"

	"github.com/gin-gonic/gin"
)

func main() {
	db := configs.InitDB()
	db.AutoMigrate(models.Order{})
	db.AutoMigrate(models.OrderDetail{})

	router := gin.Default()

	router.Use(middlewares.AuthMiddleware())
	// router.POST("/product", controllers.CreateProduct)
	// router.GET("/product", controllers.ListProduct)

	router.Run("0.0.0.0:8084")
}
