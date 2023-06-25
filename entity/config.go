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

func NewConfigResponse(config *ConfigResponse) ConfigResponse {
	return ConfigResponse{
		Key:   config.Key,
		Value: config.Value,
	}
}

func NewConfigArrayResponse(config []*ConfigResponse) []ConfigResponse {
	configRes := []ConfigResponse{}

	for _, value := range config {
		configRes = append(configRes, NewConfigResponse(value))
	}

	return configRes
}
