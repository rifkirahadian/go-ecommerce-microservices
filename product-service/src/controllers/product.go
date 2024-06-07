package controllers

import (
	"net/http"
	"shop/product-service/configs"
	"shop/product-service/src/dtos"
	"shop/product-service/src/models"
	"shop/product-service/src/utils"

	"github.com/gin-gonic/gin"
)

func CreateProduct(ctx *gin.Context) {
	var body dtos.CreateProductDto
	authUser, _ := ctx.Get("user")
	user, _ := authUser.(dtos.User)

	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.ValidationResponse(ctx, err)
		return
	}
	db := configs.InitDB()

	product := models.Product{
		Name:        body.Name,
		Description: body.Description,
		UserId:      user.ID,
		Price:       body.Price,
	}

	db.Create(&product)

	ctx.IndentedJSON(http.StatusCreated, product)
}

func ListProduct(ctx *gin.Context) {
	db := configs.InitDB()
	var products []models.Product
	db.Find(&products)

	ctx.IndentedJSON(http.StatusCreated, products)
}

func GetProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	db := configs.InitDB()
	var product models.Product
	productQuery := db.First(&product, id)
	if productQuery.RowsAffected == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"data": product})
}
