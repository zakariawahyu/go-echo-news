package services

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/channel/repository"
	"github.com/zakariawahyu/go-echo-news/pkg/exception"
	"time"
)

type ChannelServicesImpl struct {
	channelRepo    repository.ChannelRepository
	contextTimeout time.Duration
}

func NewChannelServices(channelRepo repository.ChannelRepository, timeout time.Duration) ChannelServices {
	return &ChannelServicesImpl{
		channelRepo:    channelRepo,
		contextTimeout: timeout,
	}
}

func (serv *ChannelServicesImpl) GetAllChannel(c context.Context) (channels []entity.ChannelResponse) {
	ctx, cancel := context.WithTimeout(c, serv.contextTimeout)
	defer cancel()

	res, err := serv.channelRepo.GetAllChannel(ctx)
	exception.PanicIfNeeded(err)

	for _, channel := range res {
		channels = append(channels, entity.NewChannelResponse(channel))
	}

	return channels
}

func (serv *ChannelServicesImpl) GetChannel(c context.Context, slug string) entity.ChannelResponse {
	ctx, cancel := context.WithTimeout(c, serv.contextTimeout)
	defer cancel()

	res, err := serv.channelRepo.GetChannel(ctx, slug)
	exception.PanicIfNeeded(err)

	return entity.NewChannelResponse(res)
}
