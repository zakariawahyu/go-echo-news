package services

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/item"
	"github.com/zakariawahyu/go-echo-news/pkg/helpers"
	"github.com/zakariawahyu/go-echo-news/pkg/logger"
	"time"
)

type itemServices struct {
	itemRepo       item.ItemRepository
	itemRedisRepo  item.ItemRedisRepository
	zapLogger      logger.Logger
	contextTimeout time.Duration
}

func NewItemServices(itemRepo item.ItemRepository, itemRedisRepo item.ItemRedisRepository, zapLogger logger.Logger, timeout time.Duration) item.ItemServices {
	return &itemServices{
		itemRepo:       itemRepo,
		itemRedisRepo:  itemRedisRepo,
		zapLogger:      zapLogger,
		contextTimeout: timeout,
	}
}

func (serv *itemServices) GetItemByTypes(ctx context.Context, types string) (items []entity.ItemResponse) {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	redisData, err := serv.itemRedisRepo.GetItem(c, helpers.KeyRedis("item", types))
	if redisData != nil {
		return entity.NewItemArrayResponse(redisData)
	}

	res, err := serv.itemRepo.GetByType(c, types)
	if err != nil {
		serv.zapLogger.Errorf("itemServ.GetItemByTypes.itemRepo.GetByType, err = %s", err)
		panic(err)
	}

	items = entity.NewItemArrayResponse(res)

	if err = serv.itemRedisRepo.SetItem(c, helpers.KeyRedis("item", types), helpers.Slowest, items); err != nil {
		serv.zapLogger.Errorf("itemServ.GetItemByTypes.itemRedisRepo.SetItem, err = %s", err)
		panic(err)
	}

	return items
}
