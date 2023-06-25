package services

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/region"
	"github.com/zakariawahyu/go-echo-news/pkg/logger"
	"time"
)

type regionServices struct {
	regionRepo     region.RegionRepository
	zapLogger      logger.Logger
	contextTimeout time.Duration
}

func NewRegionServices(regionRepo region.RegionRepository, zapLogger logger.Logger, timeout time.Duration) region.RegionServices {
	return &regionServices{
		regionRepo:     regionRepo,
		zapLogger:      zapLogger,
		contextTimeout: timeout,
	}
}

func (serv *regionServices) GetAllRegion(ctx context.Context) (regions []entity.RegionResponse) {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	res, err := serv.regionRepo.GetAll(c)
	if err != nil {
		serv.zapLogger.Errorf("regionServ.GetAllRegion.regionRepo.GetAll, err = %s", err)
		panic(err)
	}

	for _, region := range res {
		regions = append(regions, entity.NewRegionResponse(region))
	}

	return regions
}

func (serv *regionServices) GetRegionBySlugOrId(ctx context.Context, slug string) entity.RegionResponse {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	region, err := serv.regionRepo.GetBySlugOrId(c, slug)
	if err != nil {
		serv.zapLogger.Errorf("regionServ.GetRegionBySlugOrId.regionRepo.GetBySlugOrId, err = %s", err)
		panic(err)
	}

	return entity.NewRegionResponse(region)
}
