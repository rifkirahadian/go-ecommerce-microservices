package controllers

import (
	"fmt"
	"net/http"
	"shop/warehouse-service/configs"
	"shop/warehouse-service/src/dtos"
	"shop/warehouse-service/src/models"
	"shop/warehouse-service/src/utils"
	"sync"

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

	var wg sync.WaitGroup
	errChan := make(chan error, body.Count)

	for i := 0; i < int(body.Count); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			productItem := models.ProductItem{
				ProductId:   body.ProductId,
				Code:        utils.RandStringBytes(6),
				UserId:      user.ID,
				IsAvailable: true,
			}
			if err := db.Create(&productItem).Error; err != nil {
				errChan <- err
			}
		}()
	}

	wg.Wait()
	close(errChan)

	// Check for errors
	for err := range errChan {
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
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
