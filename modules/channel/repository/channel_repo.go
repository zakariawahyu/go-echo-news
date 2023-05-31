package repository

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
)

type ChannelRepository interface {
	GetAllChannel(ctx context.Context) ([]entity.Channel, error)
	GetChannel(ctx context.Context, slug string) (entity.Channel, error)
}
