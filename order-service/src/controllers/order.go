package controllers

import (
	"net/http"
	"shop/order-service/configs"
	"shop/order-service/src/dtos"
	"shop/order-service/src/models"
	"shop/order-service/src/services"
	"shop/order-service/src/utils"
	"sync"

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

	order := models.Order{
		Code:   utils.RandStringBytes(10),
		Total:  total,
		UserId: user.ID,
		Status: "Pending",
	}
	db.Create(&order)

	var wg sync.WaitGroup
	errChan := make(chan error, len(body.Products))

	for i := 0; i < int(len(body.Products)); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			orderDetail := models.OrderDetail{
				ProductId: body.Products[i].ProductID,
				Quantity:  body.Products[i].ProductID,
				OrderId:   order.ID,
			}
			if err := db.Create(&orderDetail).Error; err != nil {
				errChan <- err
			}
		}()
	}

	wg.Wait()
	close(errChan)

	ctx.JSON(http.StatusCreated, gin.H{"error": "Order created"})
}
