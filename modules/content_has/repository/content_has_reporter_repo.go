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

func (repo *contentHasReporterRepository) GetByReporterIDLimited(ctx context.Context, id int64, limit int) ([]*entity.ContentHasReporter, error) {
	contentHasReporter := []*entity.ContentHasReporter{}

	if err := repo.DB.NewSelect().Model(&contentHasReporter).Where("tag_id = ?", id).Order("id desc").Limit(limit).Scan(ctx); err != nil {
		return nil, err
	}

	return contentHasReporter, nil
}
