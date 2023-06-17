package repository

import (
	"context"
	"github.com/uptrace/bun"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/content"
)

type contentRepository struct {
	DB *bun.DB
}

func NewContentRepository(DB *bun.DB) content.ContentRepository {
	return &contentRepository{DB}
}

func (repo *contentRepository) GetBySlugOrId(ctx context.Context, slug string) (*entity.Content, error) {
	content := &entity.Content{}

	if err := repo.DB.NewSelect().Model(content).Relation("User").Relation("Region").Relation("Channel").Relation("SubChannel").Relation("Tags").Relation("Topics").Relation("Reporters").Relation("SubPhotos").Where("content.slug = ?", slug).WhereOr("content.id = ?", slug).Scan(ctx); err != nil {
		return nil, err
	}

	return content, nil
}
