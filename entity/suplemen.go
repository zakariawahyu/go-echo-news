package entity

import (
	"github.com/uptrace/bun"
	"time"
)

type Suplemen struct {
	bun.BaseModel `bun:"table:suplemens"`
	ID            int64     `json:"id"`
	ParentID      int64     `json:"parent-id"`
	Type          string    `json:"type"`
	Name          string    `json:"name"`
	Slug          string    `json:"slug"`
	IsActive      bool      `json:"is_active"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type SuplemenResponse struct {
	bun.BaseModel `bun:"table:suplemens"`
	ID            int64  `json:"id"`
	ParentID      int64  `json:"parent-id"`
	Name          string `json:"name"`
	Slug          string `json:"slug"`
}
