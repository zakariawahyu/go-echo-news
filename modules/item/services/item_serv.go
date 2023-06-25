package services

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/item"
	"github.com/zakariawahyu/go-echo-news/pkg/logger"
	"time"
)

type itemServices struct {
	itemRepo       item.ItemRepository
	zapLogger      logger.Logger
	contextTimeout time.Duration
}

func NewItemServices(itemRepo item.ItemRepository, zapLogger logger.Logger, timeout time.Duration) item.ItemServices {
	return &itemServices{
		itemRepo:       itemRepo,
		zapLogger:      zapLogger,
		contextTimeout: timeout,
	}
}

func (serv *itemServices) GetItemByTypes(ctx context.Context, types string) (items []entity.ItemResponse) {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	res, err := serv.itemRepo.GetByType(c, types)
	if err != nil {
		serv.zapLogger.Errorf("itemServ.GetItemByTypes.itemRepo.GetByType, err = %s", err)
		panic(err)
	}

	for _, item := range res {
		items = append(items, entity.NewItemResponse(item))
	}

	return items
}
