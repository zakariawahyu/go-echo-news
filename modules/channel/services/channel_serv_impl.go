package services

import (
	"context"
	"fmt"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/channel/repository"
	"github.com/zakariawahyu/go-echo-news/pkg/exception"
	"github.com/zakariawahyu/go-echo-news/pkg/helpers"
	"time"
)

type ChannelServicesImpl struct {
	channelRepo    repository.ChannelRepository
	redisRepo      repository.ChannelRedisRepository
	contextTimeout time.Duration
}

func NewChannelServices(channelRepo repository.ChannelRepository, redisRepo repository.ChannelRedisRepository, timeout time.Duration) ChannelServices {
	return &ChannelServicesImpl{
		channelRepo:    channelRepo,
		redisRepo:      redisRepo,
		contextTimeout: timeout,
	}
}

func (serv *ChannelServicesImpl) GetAllChannel(c context.Context) (channels []entity.ChannelResponse) {
	ctx, cancel := context.WithTimeout(c, serv.contextTimeout)
	defer cancel()

	res, err := serv.channelRepo.GetAllChannel(ctx)
	exception.PanicIfNeeded(err)

	for _, channel := range res {
		channels = append(channels, entity.NewChannelResponse(&channel))
	}

	return channels
}

func (serv *ChannelServicesImpl) GetChannel(c context.Context, slug string) entity.ChannelResponse {
	ctx, cancel := context.WithTimeout(c, serv.contextTimeout)
	defer cancel()

	newBase, err := serv.redisRepo.GetChannel(ctx, helpers.KeyRedis(fmt.Sprintf("channel-%s", slug), ""))
	if newBase != nil {
		return entity.NewChannelResponse(newBase)
	}

	channel, err := serv.channelRepo.GetChannel(ctx, slug)
	exception.PanicIfNeeded(err)

	err = serv.redisRepo.SetChannel(ctx, helpers.KeyRedis(fmt.Sprintf("channel-%s", slug), ""), helpers.Slowly, channel)
	exception.PanicIfNeeded(err)

	return entity.NewChannelResponse(channel)
}
