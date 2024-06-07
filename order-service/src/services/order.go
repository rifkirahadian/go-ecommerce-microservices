package services

import (
	"shop/order-service/src/clients"
	"shop/order-service/src/dtos"
	"shop/order-service/src/models"
	"shop/order-service/src/utils"
	"sync"

	"gorm.io/gorm"
)

func CalculateTotalPrice(orderDto dtos.OrderDto, token string) (float64, error) {
	var totalPrice float64
	productCh := make(chan *dtos.Product, len(orderDto.Products))
	errCh := make(chan error, len(orderDto.Products))
	var wg sync.WaitGroup

	for _, order := range orderDto.Products {
		wg.Add(1)
		go clients.GetProductDetailByID(order.ProductID, token, productCh, errCh, &wg)
	}

	// Close channels once all goroutines are done
	go func() {
		wg.Wait()
		close(productCh)
		close(errCh)
	}()

	for product := range productCh {
		for _, order := range orderDto.Products {
			if product.ID == order.ProductID {
				totalPrice += product.Price * float64(order.Quantity)
			}
		}
	}

	if len(errCh) > 0 {
		return 0, <-errCh
	}

	return totalPrice, nil
}

func CreateOrder(body dtos.OrderDto, user dtos.User, total float64, db *gorm.DB) (models.Order, error) {
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

	if len(errChan) > 0 {
		return order, <-errChan
	}

	return order, nil
}
