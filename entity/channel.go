package entity

import (
	"github.com/uptrace/bun"
	"time"
)

type Channel struct {
	bun.BaseModel `bun:"table:channels"`
	ID            int64     `bun:",pk" json:"id"`
	Name          string    `json:"name"`
	Slug          string    `json:"slug"`
	Ordering      int       `json:"ordering"`
	Image         string    `json:"image"`
	Logo          string    `json:"logo"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	IsActive      bool      `json:"is_active"`
	CreatedBy     int64     `json:"created_by"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type ChannelResponse struct {
	bun.BaseModel `bun:"table:channels"`
	ID            int64  `bun:",pk" json:"id"`
	Name          string `json:"name"`
	Slug          string `json:"slug"`
}
