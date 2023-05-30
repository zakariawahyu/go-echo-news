package repository

import (
	"context"
	"github.com/uptrace/bun"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/pkg/helpers"
)

type ContentRepositoryImpl struct {
	DB *bun.DB
}

func NewContentRepository(DB *bun.DB) ContentRepository {
	return &ContentRepositoryImpl{DB}
}

func (repo *ContentRepositoryImpl) GetByID(ctx context.Context, slug string) (entity.Content, error) {
	content := entity.Content{}

	err := repo.DB.NewSelect().Model(&content).Relation("User").Relation("Region").Relation("Channel").Relation("SubChannel").Relation("Tags").Relation("Topics").Relation("Reporters").Relation("SubPhoto").Where("content.id = ?", slug).Scan(ctx)
	if err != nil {
		return content, helpers.ErrNotFound
	}

	return content, nil
}

func (repo *ContentRepositoryImpl) GetBySlug(ctx context.Context, id string) (entity.Content, error) {
	content := entity.Content{}

	err := repo.DB.NewSelect().Model(&content).Relation("User").Relation("Region").Relation("Channel").Relation("SubChannel").Relation("Tags").Relation("Topics").Relation("Reporters").Relation("SubPhoto").Where("content.slug = ?", id).Scan(ctx)
	if err != nil {
		return content, helpers.ErrNotFound
	}

	return content, nil
}
