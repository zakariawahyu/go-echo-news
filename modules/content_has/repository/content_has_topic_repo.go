package repository

import (
	"context"
	"github.com/uptrace/bun"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/content_has"
)

type contentHasTopicRepository struct {
	DB *bun.DB
}

func NewContentHasTopicRepository(DB *bun.DB) content_has.ContentHasTopicRepository {
	return &contentHasTopicRepository{
		DB: DB,
	}
}

func (repo *contentHasTopicRepository) GetByTopicID(ctx context.Context, id string) (*entity.ContentHasTopic, error) {
	contentHasTopic := &entity.ContentHasTopic{}

	if err := repo.DB.NewSelect().Model(contentHasTopic).Where("topic_id = ?", id).Scan(ctx); err != nil {
		return nil, err
	}

	return contentHasTopic, nil
}
