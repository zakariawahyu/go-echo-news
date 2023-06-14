package entity

import (
	"github.com/uptrace/bun"
	"time"
)

type Channel struct {
	bun.BaseModel `bun:"table:channels"`
	ID            int64                 `bun:",pk" json:"id"`
	Name          string                `json:"name"`
	Slug          string                `json:"slug"`
	Ordering      int                   `json:"ordering"`
	Image         string                `json:"image"`
	Logo          string                `json:"logo"`
	Title         string                `json:"title"`
	Description   string                `json:"description"`
	Suplemens     []*SuplemenResponse   `bun:"rel:has-many,join:id=parent_id" json:"suplemens"`
	SubChannels   []*SubChannelResponse `bun:"rel:has-many,join:id=channel_id" json:"sub_channels"`
	IsActive      bool                  `json:"is_active"`
	CreatedBy     int64                 `json:"created_by"`
	CreatedAt     time.Time             `json:"created_at"`
	UpdatedAt     time.Time             `json:"updated_at"`
}

type ContentChannelResponse struct {
	bun.BaseModel `bun:"table:channels"`
	ID            int64  `bun:",pk" json:"id"`
	Name          string `json:"name"`
	Slug          string `json:"slug"`
}

type ChannelResponse struct {
	bun.BaseModel `bun:"table:channels"`
	ID            int64                 `bun:",pk" json:"id"`
	Name          string                `json:"name"`
	Slug          string                `json:"slug"`
	Image         string                `json:"image"`
	Logo          string                `json:"logo"`
	Title         string                `json:"title"`
	Description   string                `json:"description"`
	Suplemens     []*SuplemenResponse   `bun:"rel:has-many,join:id=parent_id" json:"suplemens"`
	SubChannels   []*SubChannelResponse `bun:"rel:has-many,join:id=channel_id" json:"sub_channels"`
}

func NewChannelResponse(channel *Channel) ChannelResponse {
	return ChannelResponse{
		ID:          channel.ID,
		Name:        channel.Name,
		Slug:        channel.Slug,
		Image:       channel.Image,
		Logo:        channel.Logo,
		Title:       channel.Title,
		Description: channel.Description,
		Suplemens:   channel.Suplemens,
		SubChannels: channel.SubChannels,
	}
}
