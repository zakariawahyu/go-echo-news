package repository

import (
	"context"
	"github.com/uptrace/bun"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/suplemen"
)

type suplemenRepository struct {
	DB *bun.DB
}

func NewSuplemenRepository(DB *bun.DB) suplemen.SuplemenRepository {
	return &suplemenRepository{
		DB: DB,
	}
}

func (repo *suplemenRepository) GetAll(ctx context.Context) ([]*entity.SuplemenResponse, error) {
	suplemen := []*entity.SuplemenResponse{}

	if err := repo.DB.NewSelect().Model(&suplemen).Where("is_active = true").Order("parent_id ASC").Scan(ctx); err != nil {
		return nil, err
	}

	return suplemen, nil
}

func (repo *suplemenRepository) GetBySlugOrId(ctx context.Context, slug string) (*entity.SuplemenResponse, error) {
	suplemen := &entity.SuplemenResponse{}

	if err := repo.DB.NewSelect().Model(suplemen).Where("is_active = true").Where("slug = ? ", slug).WhereOr("id = ?", slug).Scan(ctx); err != nil {
		return nil, err
	}

	return suplemen, nil
}
