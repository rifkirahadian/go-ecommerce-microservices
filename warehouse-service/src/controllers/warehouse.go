package controllers

import (
	"net/http"
	"shop/warehouse-service/configs"
	"shop/warehouse-service/src/dtos"
	"shop/warehouse-service/src/models"
	"shop/warehouse-service/src/utils"

	"github.com/gin-gonic/gin"
)

func CreateWarehouse(ctx *gin.Context) {
	var body dtos.CreateWarehouseDto

	authUser, _ := ctx.Get("user")
	user, _ := authUser.(dtos.User)

	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.ValidationResponse(ctx, err)
		return
	}
	db := configs.InitDB()

	warehouse := models.Warehouse{
		Name:     body.Name,
		UserId:   user.ID,
		IsActive: true,
	}

	db.Create(&warehouse)

	ctx.IndentedJSON(http.StatusCreated, warehouse)
}

func ListWarehouse(ctx *gin.Context) {
	db := configs.InitDB()
	var warehouses []models.Warehouse
	db.Find(&warehouses)

	ctx.IndentedJSON(http.StatusOK, gin.H{"data": warehouses})
}

func WarehouseStatus(ctx *gin.Context) {
	var body dtos.UpdateWarehouseStatusDto

	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.ValidationResponse(ctx, err)
		return
	}

	db := configs.InitDB()
	var warehouse models.Warehouse
	warehouseQuery := db.First(&warehouse, body.WarehouseId)
	if warehouseQuery.RowsAffected == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Warehouse not found"})
		return
	}
	warehouse.IsActive = body.IsActive
	db.Save(&warehouse)

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Warehouse status updated"})
}
