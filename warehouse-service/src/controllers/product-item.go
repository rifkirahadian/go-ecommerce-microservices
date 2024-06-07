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

	var warehouse models.Warehouse
	warehouseQuery := db.Find(&warehouse, body.WarehouseId)
	if warehouseQuery.RowsAffected == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Warehouse not found"})
		return
	}

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
				WarehouseId: body.WarehouseId,
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
	warehouseId := ctx.Query("warehouseId")
	db := configs.InitDB()
	var products []models.ProductItem
	if productId != "" {
		db.Find(&products, "product_id = ?", productId)
	}
	if warehouseId != "" {
		db.Find(&products, "warehouse_id = ?", warehouseId)
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"data": products})
}

func TransferStock(ctx *gin.Context) {
	var body dtos.TransferStockDto
	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.ValidationResponse(ctx, err)
		return
	}

	db := configs.InitDB()

	var warehouse models.Warehouse
	warehouseQuery := db.Find(&warehouse, body.WarehouseId)
	if warehouseQuery.RowsAffected == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Warehouse not found"})
		return
	}

	db.Model(models.ProductItem{}).Where("id in (?)", body.ProductItemId).Updates(models.ProductItem{WarehouseId: body.WarehouseId})

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Stock transferred"})
}
