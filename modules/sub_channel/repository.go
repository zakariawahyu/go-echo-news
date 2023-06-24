package sub_channel

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
)

type SubChannelRepository interface {
	GetAll(ctx context.Context) ([]*entity.SubChannel, error)
	GetBySlugOrId(ctx context.Context, slug string) (*entity.SubChannel, error)
	GetMetas(ctx context.Context, slug string) (interface{}, error)
}
