package entity

type Content struct {
	ID      int64  `json:"id"`
	Slug    string `json:"slug"`
	Content string `json:"content"`
}
