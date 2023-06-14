package repository

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
)

type ChannelRepository interface {
	GetAllChannel(ctx context.Context) ([]entity.Channel, error)
	GetChannel(ctx context.Context, slug string) (*entity.Channel, error)
}

type ChannelRedisRepository interface {
	GetChannel(ctx context.Context, key string) (*entity.Channel, error)
	SetChannel(ctx context.Context, key string, ttl int, channel *entity.Channel) error
}
