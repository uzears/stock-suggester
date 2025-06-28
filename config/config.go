package config

import "os"

var NewsAPIKey = os.Getenv("NEWS_API_KEY")
