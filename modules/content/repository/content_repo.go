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

	if err := repo.DB.NewSelect().Model(content).
		Relation("User").
		Relation("Region").
		Relation("Channel").
		Relation("SubChannel").
		Relation("Tags").
		Relation("Topics").
		Relation("Reporters").
		Relation("SubPhotos").
		Where("content.slug = ?", slug).
		WhereOr("content.id = ?", slug).
		Where("content.is_active", true).
		Scan(ctx); err != nil {
		return nil, err
	}

	return content, nil
}

func (repo *contentRepository) GetAllHome(ctx context.Context, limit int, offset int) (*[]entity.ContentRowResponse, error) {
	content := &[]entity.ContentRowResponse{}

	if err := repo.DB.NewSelect().Model(content).
		Relation("Region").
		Relation("Channel").
		Relation("SubChannel").
		Relation("SubPhotos").
		Relation("Tags").
		Where("content_row_response.is_active = ?", true).
		Where("headline_type != ?", 1).
		Where("is_national = ?", true).
		WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Where("ads_position is null").WhereOr("ads_position = 0")
		}).
		Order("published_date desc").
		Limit(limit).
		Offset(offset).
		Scan(ctx); err != nil {
		return nil, err
	}

	return content, nil
}

func (repo *contentRepository) GetAllChannel(ctx context.Context, type_id int64, limit int, offset int) (*[]entity.ContentRowResponse, error) {
	content := &[]entity.ContentRowResponse{}

	if err := repo.DB.NewSelect().Model(content).
		Relation("Region").
		Relation("Channel").
		Relation("SubChannel").
		Relation("SubPhotos").
		Relation("Tags").
		Where("type_id = ?", type_id).
		Where("content_row_response.is_active = ?", true).
		Where("headline_type != ?", 1).
		WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Where("ads_position is null").WhereOr("ads_position = 0")
		}).
		Order("published_date desc").
		Limit(limit).
		Offset(offset).
		Scan(ctx); err != nil {
		return nil, err
	}

	return content, nil
}

func (repo *contentRepository) GetAllSubChannel(ctx context.Context, type_child_id int64, limit int, offset int) (*[]entity.ContentRowResponse, error) {
	content := &[]entity.ContentRowResponse{}

	if err := repo.DB.NewSelect().Model(content).
		Relation("Region").
		Relation("Channel").
		Relation("SubChannel").
		Relation("SubPhotos").
		Relation("Tags").
		Where("type_child_id = ?", type_child_id).
		Where("content_row_response.is_active = ?", true).
		Where("headline_type != ?", 1).
		WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Where("ads_position is null").WhereOr("ads_position = 0")
		}).
		Order("published_date desc").
		Limit(limit).
		Offset(offset).
		Scan(ctx); err != nil {
		return nil, err
	}

	return content, nil
}
