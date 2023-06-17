package channel

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
)

type ChannelRepository interface {
	GetAll(ctx context.Context) ([]entity.Channel, error)
	GetBySlugOrId(ctx context.Context, slug string) (*entity.Channel, error)
}

type ChannelRedisRepository interface {
	GetChannel(ctx context.Context, key string) (*entity.Channel, error)
	SetChannel(ctx context.Context, key string, ttl int, channel *entity.Channel) error
}
