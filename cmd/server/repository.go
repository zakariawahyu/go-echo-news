package server

import (
	_channelRepository "github.com/zakariawahyu/go-echo-news/modules/channel/repository"
	_contentRepository "github.com/zakariawahyu/go-echo-news/modules/content/repository"
	_recommendedRepository "github.com/zakariawahyu/go-echo-news/modules/recommended/repository"
	"github.com/zakariawahyu/go-echo-news/pkg/db"
)

type Repository struct {
	contentRepo      _contentRepository.ContentRepository
	contentRedisRepo _contentRepository.ContentRedisRepository
	recommendedRepo  _recommendedRepository.RecommendedRepository
	channelRepo      _channelRepository.ChannelRepository
	channelRedisRepo _channelRepository.ChannelRedisRepository
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
