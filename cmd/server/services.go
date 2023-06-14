package server

import (
	_channelServices "github.com/zakariawahyu/go-echo-news/modules/channel/services"
	_contentServices "github.com/zakariawahyu/go-echo-news/modules/content/services"
	"time"
)

type Services struct {
	contentServices _contentServices.ContentServices
	channelServices _channelServices.ChannelServices
}

func NewServices(repo *Repository, timeoutContext time.Duration) *Services {
	return &Services{
		contentServices: _contentServices.NewContentServices(repo.contentRepo, repo.contentRedisRepo, repo.recommendedRepo, timeoutContext),
		channelServices: _channelServices.NewChannelServices(repo.channelRepo, repo.channelRedisRepo, timeoutContext),
	}
}
