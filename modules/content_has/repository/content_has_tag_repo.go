package repository

import (
	"context"
	"github.com/uptrace/bun"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/content_has"
)

type contentHasTagRepository struct {
	DB *bun.DB
}

func NewContentHasTagRepository(DB *bun.DB) content_has.ContentHasTagRepository {
	return &contentHasTagRepository{
		DB: DB,
	}
}

func (repo *contentHasTagRepository) GetByTagID(ctx context.Context, id string) ([]*entity.ContentHasTag, error) {
	contentHasTag := []*entity.ContentHasTag{}

	if err := repo.DB.NewSelect().Model(&contentHasTag).Where("tag_id = ?", id).Scan(ctx); err != nil {
		return nil, err
	}

	return contentHasTag, nil
}
