package email

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HunterResponse struct {
	Data struct {
		Status string `json:"status"`
	} `json:"data"`
}

func CheckEmailWithHunter(email, apiKey string) (bool, error) {
	apiURL := fmt.Sprintf("https://api.hunter.io/v2/email-verifier?email=%s&api_key=%s", email, apiKey)
	resp, err := http.Get(apiURL)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var result HunterResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	return result.Data.Status == "valid", nil
}
