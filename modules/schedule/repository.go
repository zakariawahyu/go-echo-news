package schedule

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
)

type ScheduleRepository interface {
	GetAll(ctx context.Context) ([]*entity.ScheduleResponse, error)
	GetBySpecificKey(ctx context.Context, key string) (*entity.ScheduleResponse, error)
}

type ScheduleRedisRepository interface {
	GetAllSchedule(ctx context.Context, key string) (interface{}, error)
	GetSchedule(ctx context.Context, key string) (*entity.ScheduleResponse, error)
	SetAllSchedule(ctx context.Context, key string, ttl int, data interface{}) error
	SetSchedule(ctx context.Context, key string, ttl int, channel *entity.ScheduleResponse) error
}
