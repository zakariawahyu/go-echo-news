package entity

import (
	"github.com/uptrace/bun"
	"time"
)

type Item struct {
	bun.BaseModel `bun:"table:items"`
	ID            int64     `bun:",pk" json:"id"`
	Type          string    `json:"type"`
	Title         string    `json:"title"`
	Slug          string    `json:"slug"`
	Description   string    `json:"description"`
	Image         string    `json:"image"`
	Url           string    `json:"url"`
	Ordering      int64     `json:"ordering"`
	IsActive      bool      `json:"is_active"`
	CreatedAt     time.Time `bun:"created" json:"created_at"`
}

type ItemResponse struct {
	bun.BaseModel `bun:"table:items"`
	ID            int64  `bun:",pk" json:"id"`
	Type          string `json:"type"`
	Title         string `json:"title"`
	Image         string `json:"image"`
	Url           string `json:"url"`
}

func NewItemResponse(item *Item) ItemResponse {
	return ItemResponse{
		ID:    item.ID,
		Type:  item.Type,
		Title: item.Title,
		Url:   item.Url,
		Image: item.Image,
	}
}
