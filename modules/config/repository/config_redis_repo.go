package repository

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/config"
	"github.com/zakariawahyu/go-echo-news/pkg/helpers"
)

type configRedisRepository struct {
	Redis *redis.Client
}

func NewConfigRedisRepo(Redis *redis.Client) config.ConfigRedisRepository {
	return &configRedisRepository{
		Redis: Redis,
	}
}

func (repo *configRedisRepository) GetAllConfig(ctx context.Context, key string) ([]*entity.ConfigResponse, error) {
	configBytes, err := repo.Redis.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	configs := []*entity.ConfigResponse{}
	if err = json.Unmarshal(configBytes, &configs); err != nil {
		return nil, err
	}

	return configs, nil
}

func (repo *configRedisRepository) GetMetas(ctx context.Context, key string) (interface{}, error) {
	configBytes, err := repo.Redis.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}
	var data interface{}
	if err = json.Unmarshal(configBytes, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func (repo *configRedisRepository) SetAllConfig(ctx context.Context, key string, ttl int, config []entity.ConfigResponse) error {
	configBytes, err := json.Marshal(&config)
	if err != nil {
		return err
	}

	if err = repo.Redis.Set(ctx, key, configBytes, helpers.TtlRedis(ttl)).Err(); err != nil {
		return err
	}

	return nil
}

func (repo *configRedisRepository) SetMetas(ctx context.Context, key string, ttl int, data interface{}) error {
	dataBytes, err := json.Marshal(&data)
	if err != nil {
		return err
	}

	if err = repo.Redis.Set(ctx, key, dataBytes, helpers.TtlRedis(ttl)).Err(); err != nil {
		return err
	}

	return nil
}
