package main

import (
	"shop/product-service/configs"
	"shop/product-service/src/controllers"
	"shop/product-service/src/middlewares"
	"shop/product-service/src/models"

	"github.com/gin-gonic/gin"
)

func main() {
	db := configs.InitDB()
	db.AutoMigrate(models.Product{})

	router := gin.Default()

	router.Use(middlewares.AuthMiddleware())
	router.POST("/product", controllers.CreateProduct)
	router.GET("/product", controllers.ListProduct)
	router.GET("/product/:id", controllers.GetProduct)

	router.Run("0.0.0.0:8082")
}
