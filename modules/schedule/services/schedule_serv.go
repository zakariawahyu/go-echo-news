package services

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/schedule"
	"github.com/zakariawahyu/go-echo-news/pkg/exception"
	"time"
)

type scheduleServices struct {
	scheduleRepo   schedule.ScheduleRepository
	contextTimeout time.Duration
}

func NewScheduleServices(scheduleRepo schedule.ScheduleRepository, timeout time.Duration) schedule.ScheduleServices {
	return &scheduleServices{
		scheduleRepo:   scheduleRepo,
		contextTimeout: timeout,
	}
}

func (serv *scheduleServices) GetAllSchedule(ctx context.Context) interface{} {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	res, err := serv.scheduleRepo.GetAll(c)
	exception.PanicIfNeeded(err)

	collections := make(map[string][]*entity.Schedule)
	for _, b := range res {
		collections[b.SpecificKey] = append(collections[b.SpecificKey], b)
	}

	return collections
}

func (serv *scheduleServices) GetScheduleBySpecificKey(ctx context.Context, key string) entity.ScheduleResponse {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	schedule, err := serv.scheduleRepo.GetBySpecificKey(c, key)
	exception.PanicIfNeeded(err)

	return entity.NewScheduleResponse(schedule)
}
