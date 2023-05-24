package entity

type SubPhoto struct {
	ID        int64  `bun:",pk" json:"id"`
	ContentID int64  `json:"content_id"`
	Content   string `json:"content"`
	Image     string `json:"image"`
}
