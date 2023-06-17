package repository

import (
	"context"
	"github.com/uptrace/bun"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/recommended_content"
)

type recommendedContentRepository struct {
	DB *bun.DB
}

func NewRecommendedContentRepository(DB *bun.DB) recommended_content.RecommendedContentRepository {
	return &recommendedContentRepository{
		DB: DB,
	}
}

func (repo *recommendedContentRepository) GetByContentID(ctx context.Context, contentID int64) ([]entity.RecommendedContent, error) {
	recommended := []entity.RecommendedContent{}

	if err := repo.DB.NewSelect().Model(&recommended).Where("content_id = ?", contentID).Scan(ctx); err != nil {
		return nil, err
	}

	return recommended, nil
}
