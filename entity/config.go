package entity

import "github.com/uptrace/bun"

type Config struct {
	bun.BaseModel `bun:"table:configs"`
	ID            int64  `bun:",pk" json:"id"`
	WebsiteID     int64  `json:"website_id"`
	Key           string `json:"key"`
	Value         string `json:"value"`
}

type ConfigResponse struct {
	bun.BaseModel `bun:"table:configs"`
	Key           string `json:"key"`
	Value         string `json:"value"`
}

func NewConfigResponse(config *Config) ConfigResponse {
	return ConfigResponse{
		Key:   config.Key,
		Value: config.Value,
	}
}
