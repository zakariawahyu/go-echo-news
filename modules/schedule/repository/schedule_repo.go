package repository

import (
	"context"
	"github.com/uptrace/bun"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/schedule"
)

type scheduleRepository struct {
	DB *bun.DB
}

func NewScheduleRepository(DB *bun.DB) schedule.ScheduleRepository {
	return &scheduleRepository{
		DB: DB,
	}
}

func (repo *scheduleRepository) GetAll(ctx context.Context) ([]*entity.ScheduleResponse, error) {
	schedule := []*entity.ScheduleResponse{}

	if err := repo.DB.NewSelect().Model(&schedule).Scan(ctx); err != nil {
		return nil, err
	}

	return schedule, nil
}

func (repo *scheduleRepository) GetBySpecificKey(ctx context.Context, key string) (*entity.ScheduleResponse, error) {
	schedule := &entity.ScheduleResponse{}

	if err := repo.DB.NewSelect().Model(schedule).Where("specific_key = ? ", key).Scan(ctx); err != nil {
		return nil, err
	}

	return schedule, nil
}
