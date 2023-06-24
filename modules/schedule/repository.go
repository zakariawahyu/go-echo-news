package schedule

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
)

type ScheduleRepository interface {
	GetAll(ctx context.Context) ([]*entity.Schedule, error)
	GetBySpecificKey(ctx context.Context, key string) (*entity.Schedule, error)
}
