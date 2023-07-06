package repository

import (
	"context"
	"github.com/uptrace/bun"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/tag"
)

type tagRepository struct {
	DB *bun.DB
}

func NewTagRepository(DB *bun.DB) tag.TagRepository {
	return &tagRepository{
		DB: DB,
	}
}

func (repo *tagRepository) GetBySlugOrID(ctx context.Context, slug string) (*entity.Tag, error) {
	tag := &entity.Tag{}

	if err := repo.DB.NewSelect().Model(tag).Where("slug = ?", slug).WhereOr("id = ?", slug).Scan(ctx); err != nil {
		return nil, err
	}

	return tag, nil
}
