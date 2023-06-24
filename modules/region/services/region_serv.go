package services

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/region"
	"github.com/zakariawahyu/go-echo-news/pkg/exception"
	"time"
)

type regionServices struct {
	regionRepo     region.RegionRepository
	contextTimeout time.Duration
}

func NewRegionServices(regionRepo region.RegionRepository, timeout time.Duration) region.RegionServices {
	return &regionServices{
		regionRepo:     regionRepo,
		contextTimeout: timeout,
	}
}

func (serv *regionServices) GetAllRegion(ctx context.Context) (regions []entity.RegionResponse) {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	res, err := serv.regionRepo.GetAll(c)
	exception.PanicIfNeeded(err)

	for _, region := range res {
		regions = append(regions, entity.NewRegionResponse(region))
	}

	return regions
}

func (serv *regionServices) GetRegionBySlugOrId(ctx context.Context, slug string) entity.RegionResponse {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	region, err := serv.regionRepo.GetBySlugOrId(c, slug)
	exception.PanicIfNeeded(err)

	return entity.NewRegionResponse(region)
}
