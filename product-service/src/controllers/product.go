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
		Stock:       body.Stock,
		UserId:      user.ID,
	}

	db.Create(&product)

	ctx.IndentedJSON(http.StatusCreated, product)
}
