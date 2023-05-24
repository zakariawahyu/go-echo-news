package entity

type Topic struct {
	ID   int64  `bun:",pk" json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}
