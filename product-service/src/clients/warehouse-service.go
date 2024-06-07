package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"shop/product-service/src/dtos"
)

var (
	warehouseServiceUrl = "http://localhost:8083"
)

func CreateProductStock(productId uint, count uint, token string) {
	// Create the payload
	payload := dtos.CreateStockDto{
		ProductID: productId,
		Count:     count,
	}

	// Marshal the payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("Error marshaling payload: %v\n", err)
		return
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/stock", warehouseServiceUrl), bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	// Add the Authorization header with the bearer token
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	// Use an HTTP client to send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Unexpected status code: %d\n", resp.StatusCode)
		return
	}
}
