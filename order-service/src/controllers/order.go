package controllers

import (
	"net/http"
	"shop/order-service/configs"
	"shop/order-service/src/dtos"
	"shop/order-service/src/services"
	"shop/order-service/src/utils"

	"github.com/gin-gonic/gin"
)

func Order(ctx *gin.Context) {
	var body dtos.OrderDto
	authUser, _ := ctx.Get("user")
	user, _ := authUser.(dtos.User)
	authToken, _ := ctx.Get("token")
	token, _ := authToken.(string)

	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.ValidationResponse(ctx, err)
		return
	}
	db := configs.InitDB()

	total, err := services.CalculateTotalPrice(body, token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, errCreate := services.CreateOrder(body, user, total, db)
	if errCreate != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errCreate.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"error": "Order created"})
}
