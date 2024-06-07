package main

import (
	"shop/warehouse-service/configs"
	"shop/warehouse-service/src/middlewares"
	"shop/warehouse-service/src/models"

	"github.com/gin-gonic/gin"
)

func main() {
	db := configs.InitDB()
	db.AutoMigrate(models.ProductItem{})

	router := gin.Default()

	router.Use(middlewares.AuthMiddleware())
	// router.POST("/product", controllers.CreateProduct)
	// router.GET("/product", controllers.ListProduct)

	router.Run("0.0.0.0:8082")
}
