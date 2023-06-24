package services

import (
	"context"
	"fmt"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/channel"
	"github.com/zakariawahyu/go-echo-news/pkg/exception"
	"github.com/zakariawahyu/go-echo-news/pkg/helpers"
	"github.com/zakariawahyu/go-echo-news/pkg/logger"
	"time"
)

type channelServices struct {
	channelRepo    channel.ChannelRepository
	redisRepo      channel.ChannelRedisRepository
	zapLogger      logger.Logger
	contextTimeout time.Duration
}

func NewChannelServices(channelRepo channel.ChannelRepository, redisRepo channel.ChannelRedisRepository, zapLogger logger.Logger, timeout time.Duration) channel.ChannelServices {
	return &channelServices{
		channelRepo:    channelRepo,
		redisRepo:      redisRepo,
		zapLogger:      zapLogger,
		contextTimeout: timeout,
	}
}

func (serv *channelServices) GetAllChannel(ctx context.Context) (channels []entity.ChannelResponse) {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	res, err := serv.channelRepo.GetAll(c)
	if err != nil {
		serv.zapLogger.Errorf("channelServ.GetAllChannel.channelRepo.GetAll, err = %s", err)
	}
	exception.PanicIfNeeded(err)

	for _, channel := range res {
		channels = append(channels, entity.NewChannelResponse(channel))
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
	if err != nil {
		serv.zapLogger.Errorf("channelServ.GetChannelBySlugOrId.channelRepo.GetBySlugOrId, err = %s", err)
	}
	exception.PanicIfNeeded(err)

	err = serv.redisRepo.SetChannel(c, helpers.KeyRedis(fmt.Sprintf("channel-%s", slug), ""), helpers.Slowly, channel)
	if err != nil {
		serv.zapLogger.Errorf("channelServ.GetChannelBySlugOrId.redisRepo.SetChannel, err = %s", err)
	}
	exception.PanicIfNeeded(err)

	return entity.NewChannelResponse(channel)
}
