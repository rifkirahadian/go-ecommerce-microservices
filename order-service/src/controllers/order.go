package controllers

import (
	"shop/order-service/configs"
	"shop/order-service/src/dtos"
	"shop/order-service/src/models"
	"shop/order-service/src/utils"
	"sync"

	"github.com/gin-gonic/gin"
)

func Order(ctx *gin.Context) {
	var body dtos.OrderDto
	authUser, _ := ctx.Get("user")
	user, _ := authUser.(dtos.User)

	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.ValidationResponse(ctx, err)
		return
	}
	db := configs.InitDB()

	total := 0
	for i := 0; i < len(body.Products); i++ {
		var product
	}

	order := models.Order{
		Code:   utils.RandStringBytes(10),
		Total:  0,
		UserId: user.ID,
		Status: "Pending",
	}

	var wg sync.WaitGroup
	errChan := make(chan error, len(body.Products))

	for i := 0; i < int(len(body.Products)); i++ {
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
}
