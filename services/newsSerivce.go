package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/uzears/stock-suggester/config"
	"github.com/uzears/stock-suggester/models"
)

func FetchNews() ([]models.NewsArticle, error) {
	url := fmt.Sprintf("https://newsapi.org/v2/everything?q=stocks&apiKey=%s", config.NewsAPIKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var parsed map[string]interface{}
	json.Unmarshal(body, &parsed)

	var articles []models.NewsArticle
	for _, item := range parsed["artiscles"].([]interface{}) {
		data := item.(map[string]interface{})
		articles = append(articles, models.NewsArticle{
			Title: data["title"].(string),
			URL:   data["url"].(string),
		})
	}
	return articles, nil
}
