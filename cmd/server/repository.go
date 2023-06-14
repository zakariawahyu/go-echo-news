package server

import (
	"github.com/zakariawahyu/go-echo-news/modules/channel"
	_channelRepository "github.com/zakariawahyu/go-echo-news/modules/channel/repository"
	"github.com/zakariawahyu/go-echo-news/modules/content"
	_contentRepository "github.com/zakariawahyu/go-echo-news/modules/content/repository"
	"github.com/zakariawahyu/go-echo-news/modules/recommended"
	_recommendedRepository "github.com/zakariawahyu/go-echo-news/modules/recommended/repository"
	"github.com/zakariawahyu/go-echo-news/pkg/db"
)

type Repository struct {
	contentRepo      content.ContentRepository
	contentRedisRepo content.ContentRedisRepository
	recommendedRepo  recommended.RecommendedRepository
	channelRepo      channel.ChannelRepository
	channelRedisRepo channel.ChannelRedisRepository
}

func NewRepository(DB *db.Conn) *Repository {
	return &Repository{
		contentRepo:      _contentRepository.NewContentRepository(DB.Mysql),
		contentRedisRepo: _contentRepository.NewContentRedisRepository(DB.Redis),
		recommendedRepo:  _recommendedRepository.NewRecommendedRepository(DB.Mysql),
		channelRepo:      _channelRepository.NewChannelRepository(DB.Mysql),
		channelRedisRepo: _channelRepository.NewChannelRedisRepository(DB.Redis),
	}
}
