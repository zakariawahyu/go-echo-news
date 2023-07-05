package repository

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/suplemen"
	"github.com/zakariawahyu/go-echo-news/pkg/helpers"
)

type suplemenRedisRepository struct {
	Redis *redis.Client
}

func NewSuplemenRedisRepository(Redis *redis.Client) suplemen.SuplemenRedisRepository {
	return &suplemenRedisRepository{
		Redis: Redis,
	}
}

func (repo *suplemenRedisRepository) GetAllSuplemen(ctx context.Context, key string) ([]*entity.SuplemenResponse, error) {
	suplemenBytes, err := repo.Redis.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	suplemen := []*entity.SuplemenResponse{}
	if err = json.Unmarshal(suplemenBytes, &suplemen); err != nil {
		return nil, err
	}

	return suplemen, nil
}

func (repo *suplemenRedisRepository) GetSuplemen(ctx context.Context, key string) (*entity.SuplemenResponse, error) {
	suplemenBytes, err := repo.Redis.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	suplemen := &entity.SuplemenResponse{}
	if err = json.Unmarshal(suplemenBytes, suplemen); err != nil {
		return nil, err
	}

	return suplemen, nil
}

func (repo *suplemenRedisRepository) SetALlSuplemen(ctx context.Context, key string, ttl int, suplemen []entity.SuplemenResponse) error {
	suplemenBytes, err := json.Marshal(&suplemen)
	if err != nil {
		return err
	}

	if err = repo.Redis.Set(ctx, key, suplemenBytes, helpers.TtlRedis(ttl)).Err(); err != nil {
		return err
	}

	return nil
}

func (repo *suplemenRedisRepository) SetSuplemen(ctx context.Context, key string, ttl int, suplemen *entity.SuplemenResponse) error {
	suplemenBytes, err := json.Marshal(&suplemen)
	if err != nil {
		return err
	}

	if err = repo.Redis.Set(ctx, key, suplemenBytes, helpers.TtlRedis(ttl)).Err(); err != nil {
		return err
	}

	return nil
}
