package item

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
)

type ItemRepository interface {
	GetByType(ctx context.Context, types string) ([]*entity.ItemResponse, error)
}

type ItemRedisRepository interface {
	GetItem(ctx context.Context, key string) ([]*entity.ItemResponse, error)
	SetItem(ctx context.Context, key string, ttl int, config []entity.ItemResponse) error
}
