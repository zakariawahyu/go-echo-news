package item

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
)

type ItemRepository interface {
	GetByType(ctx context.Context, types string) (*[]entity.Item, error)
}
