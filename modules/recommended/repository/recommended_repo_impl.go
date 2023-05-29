package repository

import (
	"context"
	"github.com/uptrace/bun"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/pkg/helpers"
)

type RecommendedRepositoryImpl struct {
	DB *bun.DB
}

func NewRecommendedRepository(DB *bun.DB) RecommendedRepository {
	return &RecommendedRepositoryImpl{
		DB: DB,
	}
}

func (repo *RecommendedRepositoryImpl) GetByContentID(ctx context.Context, contentID int64) ([]entity.RecommendedContent, error) {
	recommended := []entity.RecommendedContent{}

	err := repo.DB.NewSelect().Model(&recommended).Where("content_id = ?", contentID).Scan(ctx)
	if err != nil {
		return recommended, helpers.ErrNotFound
	}

	return recommended, nil
}
