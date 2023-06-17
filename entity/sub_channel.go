package entity

import (
	"github.com/uptrace/bun"
	"time"
)

type SubChannel struct {
	bun.BaseModel `bun:"table:sub_channels"`
	ID            int64                   `bun:",pk" json:"id"`
	ChannelID     int64                   `json:"channel_id"`
	Name          string                  `json:"name"`
	Slug          string                  `json:"slug"`
	Image         string                  `json:"image"`
	Logo          string                  `json:"logo"`
	Ordering      int64                   `json:"ordering"`
	IsActive      bool                    `json:"is_active"`
	Title         string                  `json:"title"`
	Description   string                  `json:"description"`
	Channel       *ContentChannelResponse `bun:"rel:has-one,join:channel_id=id" json:"channel"`
	CreatedAt     time.Time               `json:"created_at"`
	UpdatedAt     time.Time               `json:"updated_at"`
}

type ContentSubChannelResponse struct {
	bun.BaseModel `bun:"table:sub_channels"`
	ID            int64  `bun:",pk" json:"id"`
	Name          string `json:"name"`
	Slug          string `json:"slug"`
}

type SubChannelResponse struct {
	bun.BaseModel `bun:"table:sub_channels"`
	ID            int64                   `bun:",pk" json:"id"`
	ChannelID     int64                   `json:"channel_id"`
	Name          string                  `json:"name"`
	Slug          string                  `json:"slug"`
	Image         string                  `json:"image"`
	Logo          string                  `json:"logo"`
	Ordering      int64                   `json:"ordering"`
	IsActive      bool                    `json:"is_active"`
	Title         string                  `json:"title"`
	Description   string                  `json:"description"`
	Channel       *ContentChannelResponse `bun:"rel:has-one,join:channel_id=id" json:"channel"`
	CreatedAt     time.Time               `json:"created_at"`
	UpdatedAt     time.Time               `json:"updated_at"`
}

func NewSubChannelResponse(subChannel *SubChannel) SubChannelResponse {
	return SubChannelResponse{
		ID:          subChannel.ID,
		ChannelID:   subChannel.ChannelID,
		Name:        subChannel.Name,
		Slug:        subChannel.Slug,
		Image:       subChannel.Image,
		Logo:        subChannel.Logo,
		Ordering:    subChannel.Ordering,
		IsActive:    subChannel.IsActive,
		Title:       subChannel.Title,
		Description: subChannel.Description,
		Channel:     subChannel.Channel,
		CreatedAt:   subChannel.CreatedAt,
		UpdatedAt:   subChannel.UpdatedAt,
	}
}
