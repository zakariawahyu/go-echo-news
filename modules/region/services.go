package region

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
)

type RegionServices interface {
	GetAllRegion(ctx context.Context) (regions []entity.RegionResponse)
	GetRegionBySlugOrId(ctx context.Context, slug string) entity.RegionResponse
}
