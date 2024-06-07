package controllers

import (
	"fmt"
	"net/http"
	"shop/warehouse-service/configs"
	"shop/warehouse-service/src/dtos"
	"shop/warehouse-service/src/models"
	"shop/warehouse-service/src/utils"

	"github.com/gin-gonic/gin"
)

func CreateProductStock(ctx *gin.Context) {
	var body dtos.CreateProductItemDto
	authUser, _ := ctx.Get("user")
	user, _ := authUser.(dtos.User)

	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.ValidationResponse(ctx, err)
		return
	}
	db := configs.InitDB()

	for i := 0; i < int(body.Count); i++ {
		productItem := models.ProductItem{
			ProductId:   body.ProductId,
			Code:        utils.RandStringBytes(6),
			UserId:      user.ID,
			IsAvailable: true,
		}
		db.Create(&productItem)
	}

	ctx.IndentedJSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("%d products added", body.Count)})
}

func ListStock(ctx *gin.Context) {
	productId := ctx.Query("productId")
	db := configs.InitDB()
	var products []models.ProductItem
	db.Find(&products, "product_id = ?", productId)

	ctx.IndentedJSON(http.StatusOK, gin.H{"data": products})
}
