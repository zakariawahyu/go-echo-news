package server

import (
	"github.com/zakariawahyu/go-echo-news/modules/channel"
	_channelServices "github.com/zakariawahyu/go-echo-news/modules/channel/services"
	"github.com/zakariawahyu/go-echo-news/modules/content"
	_contentServices "github.com/zakariawahyu/go-echo-news/modules/content/services"
	"time"
)

type Services struct {
	contentServices content.ContentServices
	channelServices channel.ChannelServices
}

func NewServices(repo *Repository, timeoutContext time.Duration) *Services {
	return &Services{
		contentServices: _contentServices.NewContentServices(repo.contentRepo, repo.contentRedisRepo, repo.recommendedRepo, timeoutContext),
		channelServices: _channelServices.NewChannelServices(repo.channelRepo, repo.channelRedisRepo, timeoutContext),
	}
}
