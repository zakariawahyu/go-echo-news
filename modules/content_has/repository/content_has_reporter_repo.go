package repository

import (
	"context"
	"github.com/uptrace/bun"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/content_has"
)

type contentHasReporterRepository struct {
	DB *bun.DB
}

func NewContentHasReporterRepository(DB *bun.DB) content_has.ContentHasReporterRepository {
	return &contentHasReporterRepository{
		DB: DB,
	}
}

func (repo *contentHasReporterRepository) GetByReporterID(ctx context.Context, id string) (*entity.ContentHasReporter, error) {
	contentHasReporter := &entity.ContentHasReporter{}

	if err := repo.DB.NewSelect().Model(contentHasReporter).Where("tag_id = ?", id).Scan(ctx); err != nil {
		return nil, err
	}

	return contentHasReporter, nil
}
