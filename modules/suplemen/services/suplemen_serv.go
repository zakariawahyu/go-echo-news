package services

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/suplemen"
	"github.com/zakariawahyu/go-echo-news/pkg/helpers"
	"github.com/zakariawahyu/go-echo-news/pkg/logger"
	"time"
)

type suplemenServices struct {
	suplemenRepo      suplemen.SuplemenRepository
	suplemenRedisRepo suplemen.SuplemenRedisRepository
	zapLogger         logger.Logger
	contextTimeout    time.Duration
}

func NewSuplemenServices(suplemenRepo suplemen.SuplemenRepository, suplemenRedisRepo suplemen.SuplemenRedisRepository, zapLogger logger.Logger, timeout time.Duration) suplemen.SuplemenServices {
	return &suplemenServices{
		suplemenRepo:      suplemenRepo,
		suplemenRedisRepo: suplemenRedisRepo,
		zapLogger:         zapLogger,
		contextTimeout:    timeout,
	}
}

func (serv *suplemenServices) GetAllSuplemen(ctx context.Context) (suplemens []entity.SuplemenResponse) {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	redisData, err := serv.suplemenRedisRepo.GetAllSuplemen(c, helpers.KeyRedisType("suplemen", "all"))
	if redisData != nil {
		return entity.NewSuplemenArrayResponse(redisData)
	}

	res, err := serv.suplemenRepo.GetAll(c)
	if err != nil {
		serv.zapLogger.Errorf("suplemenServ.GetAllSuplemen.suplemenRepo.GetAll, err = %s", err)
		panic(err)
	}

	suplemens = entity.NewSuplemenArrayResponse(res)

	if err = serv.suplemenRedisRepo.SetALlSuplemen(c, helpers.KeyRedisType("suplemen", "all"), helpers.Faster, suplemens); err != nil {
		serv.zapLogger.Errorf("suplemenServ.GetAllSuplemen.suplemenRedisRepo.SetALlSuplemen, err = %s", err)
		panic(err)
	}

	return suplemens
}

func (serv *suplemenServices) GetSuplemenBySlugOrId(ctx context.Context, slug string) entity.SuplemenResponse {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	redisData, err := serv.suplemenRedisRepo.GetSuplemen(c, helpers.KeyRedisType("suplemen", slug))
	if redisData != nil {
		return entity.NewSuplemenResponse(redisData)
	}

	suplemen, err := serv.suplemenRepo.GetBySlugOrId(c, slug)
	if err != nil {
		serv.zapLogger.Errorf("suplemenServ.GetSuplemenBySlugOrId.suplemenRepo.GetBySlugOrId, err = %s", err)
		panic(err)
	}

	if err = serv.suplemenRedisRepo.SetSuplemen(c, helpers.KeyRedisType("suplemen", slug), helpers.Faster, suplemen); err != nil {
		serv.zapLogger.Errorf("suplemenServ.GetSuplemenBySlugOrId.suplemenRedisRepo.SetSuplemen, err = %s", err)
		panic(err)
	}

	return entity.NewSuplemenResponse(suplemen)
}
