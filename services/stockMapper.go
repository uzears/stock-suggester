package services

import "strings"

func GetStockSymbol(title string) string {
	mapping := map[string]string{
		"Apple":  "AAPL",
		"Tesla":  "TSLA",
		"Amazon": "AMZN",
		"Google": "GOOGL",
	}

	for name, ticker := range mapping {
		if strings.Contains(strings.ToLower(title), strings.ToLower(name)) {
			return ticker
		}
	}

	return "UNKNOWN"
}
