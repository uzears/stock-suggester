package services

import "github.com/uzears/stock-suggester/models"

func GenerateSuggestions() ([]models.Suggestion, error) {
	articles, err := FetchNews()
	if err != nil {
		return nil, err
	}

	var suggestions []models.Suggestion

	for _, article := range articles {
		sentiment, err := GetSentiment(article.Title)
		if err != nil {
			sentiment = "neutral"
		}

		suggestion := MapSentimentToAction(sentiment)

		suggestions = append(suggestions, models.Suggestion{
			Title:      article.Title,
			URL:        article.URL,
			Sentiment:  sentiment,
			Suggestion: suggestion,
			Ticker:     GetStockSymbol(article.Title),
		})
	}

	return suggestions, nil
}

func MapSentimentToAction(sentiment string) string {
	switch sentiment {
	case "positive":
		return "buy"
	case "negative":
		return "sell"
	default:
		return "hold"
	}
}
