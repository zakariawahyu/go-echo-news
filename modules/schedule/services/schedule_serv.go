package services

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/schedule"
	"github.com/zakariawahyu/go-echo-news/pkg/logger"
	"time"
)

type scheduleServices struct {
	scheduleRepo   schedule.ScheduleRepository
	zapLogger      logger.Logger
	contextTimeout time.Duration
}

func NewScheduleServices(scheduleRepo schedule.ScheduleRepository, zapLogger logger.Logger, timeout time.Duration) schedule.ScheduleServices {
	return &scheduleServices{
		scheduleRepo:   scheduleRepo,
		zapLogger:      zapLogger,
		contextTimeout: timeout,
	}
}

func (serv *scheduleServices) GetAllSchedule(ctx context.Context) interface{} {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	res, err := serv.scheduleRepo.GetAll(c)
	if err != nil {
		serv.zapLogger.Errorf("scheduleServ.GetAllSchedule.scheduleRepo.GetAll, err = %s", err)
		panic(err)
	}

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
	if err != nil {
		serv.zapLogger.Errorf("scheduleServ.GetScheduleBySpecificKey.scheduleRepo.GetBySpecificKey, err = %s", err)
		panic(err)
	}

	return entity.NewScheduleResponse(schedule)
}
