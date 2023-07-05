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

func NewSuplemenResponse(suplemen *SuplemenResponse) SuplemenResponse {
	return SuplemenResponse{
		ID:       suplemen.ID,
		ParentID: suplemen.ParentID,
		Name:     suplemen.Name,
		Slug:     suplemen.Slug,
	}
}

func NewSuplemenArrayResponse(suplemen []*SuplemenResponse) []SuplemenResponse {
	suplemenRes := []SuplemenResponse{}

	for _, value := range suplemen {
		suplemenRes = append(suplemenRes, NewSuplemenResponse(value))
	}

	return suplemenRes
}
