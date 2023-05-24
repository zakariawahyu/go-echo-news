package entity

import "github.com/uptrace/bun"

type Region struct {
	bun.BaseModel `bun:"table:regions"`
	ID            int64  `bun:",pk" json:"id"`
	Name          string `json:"name"`
	Slug          string `json:"slug"`
}

type RegionResponse struct {
	bun.BaseModel `bun:"table:regions"`
	ID            int64  `bun:",pk" json:"id"`
	Name          string `json:"name"`
	Slug          string `json:"slug"`
}
