package repository

import (
	"context"
	"github.com/uptrace/bun"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/pkg/helpers"
)

type ContentRepositoryImpl struct {
	Conn *bun.DB
}

func NewContentRepository(Conn *bun.DB) ContentRepository {
	return &ContentRepositoryImpl{Conn}
}

func (repo *ContentRepositoryImpl) GetBySlug(ctx context.Context, slug string) (entity.Content, error) {
	content := entity.Content{}

	err := repo.Conn.NewSelect().Model(&content).Relation("User").Relation("Region").Relation("Channel").Relation("SubChannel").Relation("Tags").Relation("Topics").Relation("Reporters").Relation("SubPhoto").Where("content.slug = ?", slug).Scan(ctx)
	if err != nil {
		return content, helpers.ErrNotFound
	}

	return content, nil
}
