package services

import (
	"shop/order-service/src/clients"
	"shop/order-service/src/dtos"
	"sync"
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
