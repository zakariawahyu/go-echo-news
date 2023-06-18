package item

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
)

type ItemServices interface {
	GetItemByTypes(ctx context.Context, types string) (items []entity.ItemResponse)
}
