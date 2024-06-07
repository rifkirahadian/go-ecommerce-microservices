package clients

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"shop/order-service/src/dtos"
	"sync"
)

var (
	productServiceUrl = "http://localhost:8082"
)

type GetProductResponse struct {
	Data dtos.Product `json:"data"`
}

func GetProductDetailByID(productID uint, token string, ch chan<- *dtos.Product, errCh chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()

	url := fmt.Sprintf("%s/product/%d", productServiceUrl, productID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		errCh <- fmt.Errorf("failed to create request: %w", err)
		return
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errCh <- fmt.Errorf("failed to send request: %w", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest {
		errCh <- errors.New("product id not found")
		return
	}

	if resp.StatusCode != http.StatusOK {
		errCh <- fmt.Errorf("received unexpected status %d from server", resp.StatusCode)
		return
	}

	var response GetProductResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		errCh <- fmt.Errorf("failed to decode response: %w", err)
		return
	}

	ch <- &response.Data
}
