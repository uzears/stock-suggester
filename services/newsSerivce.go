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
	url := fmt.Sprintf("https://newsapi.org/v2/everything?q=indian+stocks+order&apiKey=%s", config.NewsAPIKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal(body, &parsed); err != nil {
		return nil, err
	}

	rawArticles, ok := parsed["articles"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected format for 'articles'")
	}

	var articles []models.NewsArticle
	for _, item := range rawArticles {
		data, ok := item.(map[string]interface{})
		if !ok {
			continue
		}

		title, _ := data["title"].(string)
		url, _ := data["url"].(string)

		articles = append(articles, models.NewsArticle{
			Title: title,
			URL:   url,
		})
	}

	return articles, nil
}
