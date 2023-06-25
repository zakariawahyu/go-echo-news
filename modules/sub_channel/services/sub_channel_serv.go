package services

import (
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/sub_channel"
	"github.com/zakariawahyu/go-echo-news/pkg/logger"
	"golang.org/x/net/context"
	"time"
)

type subChannelServices struct {
	subChannelRepo sub_channel.SubChannelRepository
	zapLogger      logger.Logger
	contextTimeout time.Duration
}

func NewSubChannelServices(subChannelRepo sub_channel.SubChannelRepository, zapLogger logger.Logger, timeout time.Duration) sub_channel.SubChannelServices {
	return &subChannelServices{
		subChannelRepo: subChannelRepo,
		zapLogger:      zapLogger,
		contextTimeout: timeout,
	}
}

func (serv *subChannelServices) GetAllSubChannel(ctx context.Context) (subChannels []entity.SubChannelResponse) {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	res, err := serv.subChannelRepo.GetAll(c)
	if err != nil {
		serv.zapLogger.Errorf("subChannelServ.GetAllSubChannel.subChannelRepo.GetAll, err = %s", err)
		panic(err)
	}

	for _, subChannel := range res {
		subChannels = append(subChannels, entity.NewSubChannelResponse(subChannel))
	}

	return subChannels
}

func (serv *subChannelServices) GetSubChannelBySlugOrId(ctx context.Context, slug string) entity.SubChannelResponse {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	subChannel, err := serv.subChannelRepo.GetBySlugOrId(c, slug)
	if err != nil {
		serv.zapLogger.Errorf("subChannelServ.GetSubChannelBySlugOrId.subChannelRepo.GetBySlugOrId, err = %s", err)
		panic(err)
	}

	return entity.NewSubChannelResponse(subChannel)
}
