package repository

import (
	"context"
	"github.com/uptrace/bun"
	"github.com/zakariawahyu/go-echo-news/entity"
)

type ChannelRepositoryImpl struct {
	DB *bun.DB
}

func NewChannelRepository(DB *bun.DB) ChannelRepository {
	return &ChannelRepositoryImpl{
		DB: DB,
	}
}

func (repo *ChannelRepositoryImpl) GetAllChannel(ctx context.Context) ([]entity.Channel, error) {
	channel := []entity.Channel{}

	err := repo.DB.NewSelect().Model(&channel).Relation("Suplemens").Relation("SubChannels").Scan(ctx)
	if err != nil {
		return channel, err
	}

	return channel, nil
}

func (repo *ChannelRepositoryImpl) GetChannel(ctx context.Context, slug string) (entity.Channel, error) {
	channel := entity.Channel{}

	err := repo.DB.NewSelect().Model(&channel).Relation("Suplemens").Relation("SubChannels").Where("channel.slug = ?", slug).WhereOr("channel.id = ?", slug).Scan(ctx)
	if err != nil {
		return channel, err
	}

	return channel, nil
}
