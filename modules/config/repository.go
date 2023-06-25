package config

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
)

type ConfigRepository interface {
	GetAll(ctx context.Context) ([]*entity.ConfigResponse, error)
}

type ConfigRedisRepository interface {
	GetAllConfig(ctx context.Context, key string) ([]*entity.ConfigResponse, error)
	SetAllConfig(ctx context.Context, key string, ttl int, config []entity.ConfigResponse) error
}
