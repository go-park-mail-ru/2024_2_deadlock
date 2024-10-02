package domain

type ArticleID int

type Article struct {
	ID       ArticleID `json:"id"`
	Title    string    `json:"title"`
	MediaURL string    `json:"media-url"`
	Body     string    `json:"body"`
}
