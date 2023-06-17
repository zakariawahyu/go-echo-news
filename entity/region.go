package entity

import (
	"github.com/uptrace/bun"
	"time"
)

type Region struct {
	bun.BaseModel `bun:"table:regions"`
	ID            int64     `bun:",pk" json:"id"`
	Name          string    `json:"name"`
	Slug          string    `json:"slug"`
	Image         string    `json:"image"`
	Logo          string    `json:"logo"`
	Ordering      int64     `json:"ordering"`
	IsActive      bool      `json:"is_active"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type ContentRegionResponse struct {
	bun.BaseModel `bun:"table:regions"`
	ID            int64  `bun:",pk" json:"id"`
	Name          string `json:"name"`
	Slug          string `json:"slug"`
}

type RegionResponse struct {
	bun.BaseModel `bun:"table:regions"`
	ID            int64     `bun:",pk" json:"id"`
	Name          string    `json:"name"`
	Slug          string    `json:"slug"`
	Image         string    `json:"image"`
	Logo          string    `json:"logo"`
	Ordering      int64     `json:"ordering"`
	IsActive      bool      `json:"is_active"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func NewRegionResponse(region *Region) RegionResponse {
	return RegionResponse{
		ID:          region.ID,
		Name:        region.Name,
		Slug:        region.Slug,
		Image:       region.Image,
		Logo:        region.Logo,
		Ordering:    region.Ordering,
		IsActive:    region.IsActive,
		Title:       region.Title,
		Description: region.Description,
		CreatedAt:   region.CreatedAt,
		UpdatedAt:   region.UpdatedAt,
	}
}
