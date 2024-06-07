package controllers

import (
	"net/http"
	"shop/product-service/configs"
	"shop/product-service/src/clients"
	"shop/product-service/src/dtos"
	"shop/product-service/src/models"
	"shop/product-service/src/utils"

	"github.com/gin-gonic/gin"
)

func CreateProduct(ctx *gin.Context) {
	var body dtos.CreateProductDto
	authUser, _ := ctx.Get("user")
	bearerToken, _ := ctx.Get("token")
	user, _ := authUser.(dtos.User)
	token, _ := bearerToken.(string)

	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.ValidationResponse(ctx, err)
		return
	}
	db := configs.InitDB()

	product := models.Product{
		Name:        body.Name,
		Description: body.Description,
		UserId:      user.ID,
	}

	db.Create(&product)
	clients.CreateProductStock(product.ID, body.Stock, token)

	ctx.IndentedJSON(http.StatusCreated, product)
}

func ListProduct(ctx *gin.Context) {
	db := configs.InitDB()
	var products []models.Product
	db.Find(&products)

	ctx.IndentedJSON(http.StatusCreated, products)
}
