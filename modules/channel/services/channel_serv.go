package services

import (
	"context"
	"fmt"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/channel"
	"github.com/zakariawahyu/go-echo-news/pkg/exception"
	"github.com/zakariawahyu/go-echo-news/pkg/helpers"
	"time"
)

type channelServices struct {
	channelRepo    channel.ChannelRepository
	redisRepo      channel.ChannelRedisRepository
	contextTimeout time.Duration
}

func NewChannelServices(channelRepo channel.ChannelRepository, redisRepo channel.ChannelRedisRepository, timeout time.Duration) channel.ChannelServices {
	return &channelServices{
		channelRepo:    channelRepo,
		redisRepo:      redisRepo,
		contextTimeout: timeout,
	}
}

func (serv *channelServices) GetAllChannel(ctx context.Context) (channels []entity.ChannelResponse) {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	res, err := serv.channelRepo.GetAll(c)
	exception.PanicIfNeeded(err)

	for _, channel := range res {
		channels = append(channels, entity.NewChannelResponse(&channel))
	}

	return channels
}

func (serv *channelServices) GetChannelBySlugOrId(ctx context.Context, slug string) entity.ChannelResponse {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	newBase, err := serv.redisRepo.GetChannel(c, helpers.KeyRedis(fmt.Sprintf("channel-%s", slug), ""))
	if newBase != nil {
		return entity.NewChannelResponse(newBase)
	}

	channel, err := serv.channelRepo.GetBySlugOrId(c, slug)
	exception.PanicIfNeeded(err)

	err = serv.redisRepo.SetChannel(c, helpers.KeyRedis(fmt.Sprintf("channel-%s", slug), ""), helpers.Slowly, channel)
	exception.PanicIfNeeded(err)

	return entity.NewChannelResponse(channel)
}
