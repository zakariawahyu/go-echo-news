package repository

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/content"
	"github.com/zakariawahyu/go-echo-news/pkg/helpers"
)

type contentRedisRepository struct {
	Redis *redis.Client
}

func NewContentRedisRepository(Redis *redis.Client) content.ContentRedisRepository {
	return &contentRedisRepository{
		Redis: Redis,
	}
}

func (repo *contentRedisRepository) GetAllContentRow(ctx context.Context, key string) ([]*entity.ContentRowResponse, error) {
	contentBytes, err := repo.Redis.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	content := []*entity.ContentRowResponse{}
	if err = json.Unmarshal(contentBytes, &content); err != nil {
		return nil, err
	}

	return content, nil
}

func (repo *contentRedisRepository) GetContent(ctx context.Context, key string) (*entity.Content, error) {
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

func (repo *contentRedisRepository) SetALlContentRow(ctx context.Context, key string, ttl int, content []entity.ContentRowResponse) error {
	contentBytes, err := json.Marshal(&content)
	if err != nil {
		return err
	}

	if err = repo.Redis.Set(ctx, key, contentBytes, helpers.TtlRedis(ttl)).Err(); err != nil {
		return err
	}

	return nil
}

func (repo *contentRedisRepository) SetContent(ctx context.Context, key string, ttl int, content *entity.Content) error {
	contentBytes, err := json.Marshal(&content)
	if err != nil {
		return err
	}

	if err = repo.Redis.Set(ctx, key, contentBytes, helpers.TtlRedis(ttl)).Err(); err != nil {
		return err
	}

	return nil
}
