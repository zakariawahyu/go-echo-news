package services

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/item"
	"github.com/zakariawahyu/go-echo-news/pkg/exception"
	"time"
)

type itemServices struct {
	itemRepo       item.ItemRepository
	contextTimeout time.Duration
}

func NewItemServices(itemRepo item.ItemRepository, timeout time.Duration) item.ItemServices {
	return &itemServices{
		itemRepo:       itemRepo,
		contextTimeout: timeout,
	}
}

func (serv *itemServices) GetItemByTypes(ctx context.Context, types string) (items []entity.ItemResponse) {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	res, err := serv.itemRepo.GetByType(c, types)
	exception.PanicIfNeeded(err)

	for _, item := range *res {
		items = append(items, entity.NewItemResponse(&item))
	}

	return items
}
