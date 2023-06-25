package repository

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/item"
	"github.com/zakariawahyu/go-echo-news/pkg/helpers"
)

type itemRedisRepository struct {
	Redis *redis.Client
}

func NewItemRedisRepository(Redis *redis.Client) item.ItemRedisRepository {
	return &itemRedisRepository{
		Redis: Redis,
	}
}

func (repo *itemRedisRepository) GetItem(ctx context.Context, key string) ([]*entity.ItemResponse, error) {
	itemBytes, err := repo.Redis.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	items := []*entity.ItemResponse{}
	if err = json.Unmarshal(itemBytes, &items); err != nil {
		return nil, err
	}

	return items, nil
}

func (repo *itemRedisRepository) SetItem(ctx context.Context, key string, ttl int, config []entity.ItemResponse) error {
	itemBytes, err := json.Marshal(&config)
	if err != nil {
		return err
	}

	if err = repo.Redis.Set(ctx, key, itemBytes, helpers.TtlRedis(ttl)).Err(); err != nil {
		return err
	}

	return nil
}
