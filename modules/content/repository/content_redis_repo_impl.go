package repository

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/pkg/helpers"
)

type ContentRedisRepositoryImpl struct {
	Redis *redis.Client
}

func NewContentRedisRepository(Redis *redis.Client) ContentRedisRepository {
	return &ContentRedisRepositoryImpl{
		Redis: Redis,
	}
}

func (repo *ContentRedisRepositoryImpl) GetContent(ctx context.Context, key string) (*entity.Content, error) {
	contentBytes, err := repo.Redis.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	content := &entity.Content{}
	if err = json.Unmarshal(contentBytes, content); err != nil {
		return nil, err
	}

	return content, nil
}

func (repo *ContentRedisRepositoryImpl) SetContent(ctx context.Context, key string, ttl int, content *entity.Content) error {
	contentBytes, err := json.Marshal(&content)
	if err != nil {
		return err
	}

	if err = repo.Redis.Set(ctx, key, contentBytes, helpers.TtlRedis(ttl)).Err(); err != nil {
		return err
	}

	return nil
}
