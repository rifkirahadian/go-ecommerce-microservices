package main

import (
	"shop/warehouse-service/configs"
	"shop/warehouse-service/src/controllers"
	"shop/warehouse-service/src/middlewares"
	"shop/warehouse-service/src/models"

	"github.com/gin-gonic/gin"
)

func main() {
	db := configs.InitDB()
	db.AutoMigrate(models.ProductItem{})
	db.AutoMigrate(models.Warehouse{})

	router := gin.Default()

	router.Use(middlewares.AuthMiddleware())
	router.POST("/stock", controllers.CreateProductStock)
	router.GET("/stock", controllers.ListStock)
	router.POST("/warehouse", controllers.CreateWarehouse)
	router.GET("/warehouse", controllers.ListWarehouse)
	router.POST("/stock-transfer", controllers.TransferStock)

	router.Run("0.0.0.0:8083")
}
