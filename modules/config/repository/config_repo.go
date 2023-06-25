package repository

import (
	"context"
	"github.com/uptrace/bun"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/config"
)

type configRepository struct {
	DB *bun.DB
}

func NewConfigRepository(DB *bun.DB) config.ConfigRepository {
	return &configRepository{
		DB: DB,
	}
}

func (repo *configRepository) GetAll(ctx context.Context) ([]*entity.ConfigResponse, error) {
	config := []*entity.ConfigResponse{}

	if err := repo.DB.NewSelect().Model(&config).Scan(ctx); err != nil {
		return nil, err
	}

	return config, nil
}
