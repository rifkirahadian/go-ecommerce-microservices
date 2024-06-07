package clients

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"shop/warehouse-service/src/dtos"
)

var (
	authServiceUrl = "http://localhost:8081"
)

func ValidateToken(token string) (bool, dtos.User, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/auth/user", authServiceUrl), nil)
	if err != nil {
		return false, dtos.User{}, err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, dtos.User{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, dtos.User{}, fmt.Errorf("validation service returned status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, dtos.User{}, err
	}

	var result struct {
		Data dtos.User `json:"data"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return false, dtos.User{}, err
	}

	return true, result.Data, nil
}
