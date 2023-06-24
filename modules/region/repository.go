package region

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
)

type RegionRepository interface {
	GetAll(ctx context.Context) ([]*entity.Region, error)
	GetBySlugOrId(ctx context.Context, slug string) (*entity.Region, error)
	GetMetas(ctx context.Context, slug string) (interface{}, error)
}
