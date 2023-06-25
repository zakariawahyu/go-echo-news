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
	GetMetas(ctx context.Context, key string) (interface{}, error)
	SetAllConfig(ctx context.Context, key string, ttl int, config []entity.ConfigResponse) error
	SetMetas(ctx context.Context, key string, ttl int, data interface{}) error
}
