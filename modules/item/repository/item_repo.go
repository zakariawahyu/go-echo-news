package repository

import (
	"context"
	"github.com/uptrace/bun"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/item"
)

type itemRepository struct {
	DB *bun.DB
}

func NewItemRepository(DB *bun.DB) item.ItemRepository {
	return &itemRepository{
		DB: DB,
	}
}

func (repo *itemRepository) GetByType(ctx context.Context, types string) (*[]entity.Item, error) {
	item := &[]entity.Item{}

	if err := repo.DB.NewSelect().Model(item).Where("type = ?", types).Scan(ctx); err != nil {
		return nil, err
	}

	return item, nil
}
