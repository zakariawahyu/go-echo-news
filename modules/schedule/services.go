package schedule

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
)

type ScheduleServices interface {
	GetAllSchedule(ctx context.Context) interface{}
	GetScheduleBySpecificKey(ctx context.Context, key string) entity.ScheduleResponse
}
