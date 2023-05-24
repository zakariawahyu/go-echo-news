package repository

import (
	"context"
	"github.com/uptrace/bun"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/helpers"
)

type ContentRepositoryImpl struct {
	Conn *bun.DB
}

func NewContentRepository(Conn *bun.DB) ContentRepository {
	return &ContentRepositoryImpl{Conn}
}

//var querySelectContent = `SELECT id, slug, content FROM contents_new WHERE is_active = true`
//
//func (repo *ContentRepositoryImpl) fetch(ctx context.Context, query string, args ...interface{}) (result []entity.Content, err error) {
//	rows, err := repo.Conn.QueryContext(ctx, query, args...)
//	if err != nil {
//		return nil, err
//	}
//
//	result = make([]entity.Content, 0)
//	for rows.Next() {
//		content := entity.Content{}
//		err = rows.Scan(
//			&content.ID,
//			&content.Slug,
//			&content.Content,
//		)
//
//		if err != nil {
//			return nil, err
//		}
//		result = append(result, content)
//	}
//
//	return result, nil
//}
//
//func (repo *ContentRepositoryImpl) count(ctx context.Context, query string) (total int64, err error) {
//
//	err = repo.Conn.QueryRow(query).Scan(&total)
//	if err != nil {
//		return
//	}
//
//	return total, nil
//}
//
//func (repo *ContentRepositoryImpl) Fetch(ctx context.Context, params *utils.Request) (res []entity.Content, total int64, err error) {
//	query := querySelectContent
//
//	query += ` ORDER BY name LIMIT ?,? `
//
//	res, err = repo.fetch(ctx, query, params.Offset, params.PerPage)
//
//	if err != nil {
//		return nil, 0, err
//	}
//
//	total, _ = repo.count(ctx, "SELECT COUNT(1) FROM contents_new")
//
//	return
//}

func (repo *ContentRepositoryImpl) GetBySlug(ctx context.Context, slug string) (entity.Content, error) {
	content := new(entity.Content)
	err := repo.Conn.NewSelect().Model(content).Relation("User").Relation("Region").Relation("Channel").Relation("SubChannel").Relation("Tags").Relation("Topics").Relation("Reporters").Relation("SubPhoto").Where("content.slug = ?", slug).Scan(ctx)
	if err != nil {
		return *content, helpers.ErrNotFound
	}

	return *content, nil
}
