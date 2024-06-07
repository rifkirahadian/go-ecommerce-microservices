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

	router.POST("/product", middlewares.AuthMiddleware(), controllers.CreateProduct)

	router.Run("0.0.0.0:8082")
}
