package services

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/suplemen"
	"github.com/zakariawahyu/go-echo-news/pkg/logger"
	"time"
)

type suplemenServices struct {
	suplemenRepo   suplemen.SuplemenRepository
	zapLogger      logger.Logger
	contextTimeout time.Duration
}

func NewSuplemenServices(suplemenRepo suplemen.SuplemenRepository, zapLogger logger.Logger, timeout time.Duration) suplemen.SuplemenServices {
	return &suplemenServices{
		suplemenRepo:   suplemenRepo,
		zapLogger:      zapLogger,
		contextTimeout: timeout,
	}
}

func (serv *suplemenServices) GetAllSuplemen(ctx context.Context) (suplemens []entity.SuplemenResponse) {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	res, err := serv.suplemenRepo.GetAll(c)
	if err != nil {
		serv.zapLogger.Errorf("suplemenServ.GetAllSuplemen.suplemenRepo.GetAll, err = %s", err)
		panic(err)
	}

	suplemens = entity.NewSuplemenArrayResponse(res)

	return suplemens
}

func (serv *suplemenServices) GetSuplemenBySlugOrId(ctx context.Context, slug string) entity.SuplemenResponse {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	suplemen, err := serv.suplemenRepo.GetBySlugOrId(c, slug)
	if err != nil {
		serv.zapLogger.Errorf("suplemenServ.GetSuplemenBySlugOrId.suplemenRepo.GetBySlugOrId, err = %s", err)
		panic(err)
	}

	return entity.NewSuplemenResponse(suplemen)
}
