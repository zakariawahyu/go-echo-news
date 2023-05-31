package services

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
)

type ChannelServices interface {
	GetAllChannel(c context.Context) (channels []entity.ChannelResponse)
	GetChannel(c context.Context, slug string) entity.ChannelResponse
}
