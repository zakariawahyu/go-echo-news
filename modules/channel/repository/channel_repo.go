package repository

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/channel"
)

type ChannelRepository struct {
	DB *bun.DB
}

func NewChannelRepository(DB *bun.DB) channel.ChannelRepository {
	return &ChannelRepository{
		DB: DB,
	}
}

func (repo *ChannelRepository) GetAll(ctx context.Context) ([]*entity.Channel, error) {
	channel := []*entity.Channel{}

	if err := repo.DB.NewSelect().Model(&channel).Relation("Suplemens").Relation("SubChannels").Scan(ctx); err != nil {
		return nil, err
	}

	return channel, nil
}

func (repo *ChannelRepository) GetBySlugOrId(ctx context.Context, slug string) (*entity.Channel, error) {
	channel := &entity.Channel{}

	if err := repo.DB.NewSelect().Model(channel).Relation("Suplemens").Relation("SubChannels").Where("channel.slug = ?", slug).WhereOr("channel.id = ?", slug).Scan(ctx); err != nil {
		return nil, err
	}

	return channel, nil
}

func (repo *ChannelRepository) GetMetas(ctx context.Context, slug string) (interface{}, error) {
	channel := &entity.Channel{}

	if err := repo.DB.NewSelect().Model(channel).ColumnExpr("title, description").Where("slug = ?", slug).Scan(ctx); err != nil {
		return nil, err
	}

	data := echo.Map{
		"title":       channel.Title,
		"description": channel.Description,
	}

	return data, nil
}
