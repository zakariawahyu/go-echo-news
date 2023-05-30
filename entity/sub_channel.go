package entity

import "github.com/uptrace/bun"

type SubChannel struct {
	bun.BaseModel `bun:"table:sub_channels"`
	ID            int64  `bun:",pk" json:"id"`
	Name          string `json:"name"`
	Slug          string `json:"slug"`
}

type SubChannelResponse struct {
	bun.BaseModel `bun:"table:sub_channels"`
	ID            int64  `bun:",pk" json:"id"`
	Name          string `json:"name"`
	Slug          string `json:"slug"`
}