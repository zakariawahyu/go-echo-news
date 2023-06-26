package services

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/schedule"
	"github.com/zakariawahyu/go-echo-news/pkg/helpers"
	"github.com/zakariawahyu/go-echo-news/pkg/logger"
	"time"
)

type scheduleServices struct {
	scheduleRepo      schedule.ScheduleRepository
	scheduleRedisRepo schedule.ScheduleRedisRepository
	zapLogger         logger.Logger
	contextTimeout    time.Duration
}

func NewScheduleServices(scheduleRepo schedule.ScheduleRepository, scheduleRedisRepo schedule.ScheduleRedisRepository, zapLogger logger.Logger, timeout time.Duration) schedule.ScheduleServices {
	return &scheduleServices{
		scheduleRepo:      scheduleRepo,
		scheduleRedisRepo: scheduleRedisRepo,
		zapLogger:         zapLogger,
		contextTimeout:    timeout,
	}
}

func (serv *scheduleServices) GetAllSchedule(ctx context.Context) interface{} {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	redisData, err := serv.scheduleRedisRepo.GetAllSchedule(c, helpers.KeyRedisType("live_stream", ""))
	if redisData != nil {
		return redisData
	}

	res, err := serv.scheduleRepo.GetAll(c)
	if err != nil {
		serv.zapLogger.Errorf("scheduleServ.GetAllSchedule.scheduleRepo.GetAll, err = %s", err)
		panic(err)
	}

	collections := make(map[string][]*entity.ScheduleResponse)
	for _, b := range res {
		collections[b.SpecificKey] = append(collections[b.SpecificKey], b)
	}

	if err = serv.scheduleRedisRepo.SetAllSchedule(c, helpers.KeyRedisType("live_stream", ""), helpers.Slowest, collections); err != nil {
		serv.zapLogger.Errorf("scheduleServ.GetAllSchedule.scheduleRedisRepo.SetAllSchedule, err = %s", err)
		panic(err)
	}

	return collections
}

func (serv *scheduleServices) GetScheduleBySpecificKey(ctx context.Context, key string) entity.ScheduleResponse {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	redisData, err := serv.scheduleRedisRepo.GetSchedule(c, helpers.KeyRedisType("live_stream", key))
	if redisData != nil {
		return entity.NewScheduleResponse(redisData)
	}

	schedule, err := serv.scheduleRepo.GetBySpecificKey(c, key)
	if err != nil {
		serv.zapLogger.Errorf("scheduleServ.GetScheduleBySpecificKey.scheduleRepo.GetBySpecificKey, err = %s", err)
		panic(err)
	}

	if err = serv.scheduleRedisRepo.SetSchedule(c, helpers.KeyRedisType("live_stream", key), helpers.Slowest, schedule); err != nil {
		serv.zapLogger.Errorf("scheduleServ.GetScheduleBySpecificKey.scheduleRedisRepo.SetSchedule, err = %s", err)
		panic(err)
	}

	return entity.NewScheduleResponse(schedule)
}
