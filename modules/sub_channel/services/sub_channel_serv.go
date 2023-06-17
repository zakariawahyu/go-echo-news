package services

import (
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/sub_channel"
	"github.com/zakariawahyu/go-echo-news/pkg/exception"
	"golang.org/x/net/context"
	"time"
)

type subChannelServices struct {
	subChannelRepo sub_channel.SubChannelRepository
	contextTimeout time.Duration
}

func NewSubChannelServices(subChannelRepo sub_channel.SubChannelRepository, timeout time.Duration) sub_channel.SubChannelServices {
	return &subChannelServices{
		subChannelRepo: subChannelRepo,
		contextTimeout: timeout,
	}
}

func (serv *subChannelServices) GetAllSubChannel(ctx context.Context) (subChannels []entity.SubChannelResponse) {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	res, err := serv.subChannelRepo.GetAll(c)
	exception.PanicIfNeeded(err)

	for _, subChannel := range *res {
		subChannels = append(subChannels, entity.NewSubChannelResponse(&subChannel))
	}

	return subChannels
}

func (serv *subChannelServices) GetSubChannelBySlugOrId(ctx context.Context, slug string) entity.SubChannelResponse {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	subChannel, err := serv.subChannelRepo.GetBySlugOrId(c, slug)
	exception.PanicIfNeeded(err)

	return entity.NewSubChannelResponse(subChannel)
}
