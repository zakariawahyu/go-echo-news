package repository

import (
	"context"
	"github.com/uptrace/bun"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/region"
)

type regionRepository struct {
	DB *bun.DB
}

func NewRegionRepository(DB *bun.DB) region.RegionRepository {
	return &regionRepository{
		DB: DB,
	}
}

func (repo *regionRepository) GetAll(ctx context.Context) (*[]entity.Region, error) {
	region := &[]entity.Region{}

	if err := repo.DB.NewSelect().Model(region).Scan(ctx); err != nil {
		return nil, err
	}

	return region, nil
}

func (repo *regionRepository) GetBySlugOrId(ctx context.Context, slug string) (*entity.Region, error) {
	region := &entity.Region{}

	if err := repo.DB.NewSelect().Model(region).Where("region.slug = ? ", slug).WhereOr("region.id = ?", slug).Scan(ctx); err != nil {
		return nil, err
	}

	return region, nil
}
