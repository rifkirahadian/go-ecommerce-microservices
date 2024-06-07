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
