package services

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func GetSentiment(text string) (string, error) {
	payload := map[string]string{"text": text}
	jsonData, _ := json.Marshal(payload)

	resp, err := http.Post("http://localhost:5000/sentiment", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "neutral", err
	}
	defer resp.Body.Close()

	var result map[string]string
	json.NewDecoder(resp.Body).Decode(&result)

	return result["sentiment"], nil
}
