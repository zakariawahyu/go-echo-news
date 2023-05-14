package repository

import (
	"context"
	"database/sql"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/helpers"
	"github.com/zakariawahyu/go-echo-news/utils"
)

type ContentRepositoryImpl struct {
	Conn *sql.DB
}

func NewContentRepository(Conn *sql.DB) ContentRepository {
	return &ContentRepositoryImpl{Conn}
}

var querySelectContent = `SELECT id, slug, content FROM contents_new WHERE is_active = true`

func (repo *ContentRepositoryImpl) fetch(ctx context.Context, query string, args ...interface{}) (result []entity.Content, err error) {
	rows, err := repo.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	result = make([]entity.Content, 0)
	for rows.Next() {
		content := entity.Content{}
		err = rows.Scan(
			&content.ID,
			&content.Slug,
			&content.Content,
		)

		if err != nil {
			return nil, err
		}
		result = append(result, content)
	}

	return result, nil
}

func (repo *ContentRepositoryImpl) count(ctx context.Context, query string) (total int64, err error) {

	err = repo.Conn.QueryRow(query).Scan(&total)
	if err != nil {
		return
	}

	return total, nil
}

func (repo *ContentRepositoryImpl) Fetch(ctx context.Context, params *utils.Request) (res []entity.Content, total int64, err error) {
	query := querySelectContent

	query += ` ORDER BY name LIMIT ?,? `

	res, err = repo.fetch(ctx, query, params.Offset, params.PerPage)

	if err != nil {
		return nil, 0, err
	}

	total, _ = repo.count(ctx, "SELECT COUNT(1) FROM contents_new")

	return
}

func (repo *ContentRepositoryImpl) GetBySlug(ctx context.Context, slug string) (res entity.Content, err error) {
	query := querySelectContent + ` AND slug = ?`

	list, err := repo.fetch(ctx, query, slug)
	if err != nil {
		return entity.Content{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, helpers.ErrNotFound
	}

	return
}
