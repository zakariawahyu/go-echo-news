package repository

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/schedule"
	"github.com/zakariawahyu/go-echo-news/pkg/helpers"
)

type scheduleRedisRepository struct {
	Redis *redis.Client
}

func NewScheduleRedisRepository(Redis *redis.Client) schedule.ScheduleRedisRepository {
	return &scheduleRedisRepository{
		Redis: Redis,
	}
}

func (repo *scheduleRedisRepository) GetAllSchedule(ctx context.Context, key string) (interface{}, error) {
	scheduleBytes, err := repo.Redis.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	var data interface{}
	if err = json.Unmarshal(scheduleBytes, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func (repo *scheduleRedisRepository) GetSchedule(ctx context.Context, key string) (*entity.ScheduleResponse, error) {
	scheduleBytes, err := repo.Redis.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	schedule := &entity.ScheduleResponse{}
	if err = json.Unmarshal(scheduleBytes, schedule); err != nil {
		return nil, err
	}

	return schedule, nil
}

func (repo *scheduleRedisRepository) SetAllSchedule(ctx context.Context, key string, ttl int, data interface{}) error {
	scheduleBytes, err := json.Marshal(&data)
	if err != nil {
		return err
	}

	if err = repo.Redis.Set(ctx, key, scheduleBytes, helpers.TtlRedis(ttl)).Err(); err != nil {
		return err
	}

	return nil
}

func (repo *scheduleRedisRepository) SetSchedule(ctx context.Context, key string, ttl int, channel *entity.ScheduleResponse) error {
	scheduleBytes, err := json.Marshal(channel)
	if err != nil {
		return err
	}

	if err = repo.Redis.Set(ctx, key, scheduleBytes, helpers.TtlRedis(ttl)).Err(); err != nil {
		return err
	}

	return nil
}
