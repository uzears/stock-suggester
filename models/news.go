package models

type NewsArticle struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type Suggestion struct {
	Title      string `json:"title"`
	URL        string `json:"url"`
	Sentiment  string `json:"sentiment"`
	Suggestion string `json:"suggestion"`
	Ticker     string `json:"ticker"`
}
