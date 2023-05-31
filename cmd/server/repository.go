package server

import (
	_channelRepository "github.com/zakariawahyu/go-echo-news/modules/channel/repository"
	_contentRepository "github.com/zakariawahyu/go-echo-news/modules/content/repository"
	_recommendedRepository "github.com/zakariawahyu/go-echo-news/modules/recommended/repository"
	"github.com/zakariawahyu/go-echo-news/pkg"
)

type Repository struct {
	contentRepo     _contentRepository.ContentRepository
	recommendedRepo _recommendedRepository.RecommendedRepository
	channelRepo     _channelRepository.ChannelRepository
}

func NewRepository(DB *pkg.Conn) *Repository {
	return &Repository{
		contentRepo:     _contentRepository.NewContentRepository(DB.Mysql),
		recommendedRepo: _recommendedRepository.NewRecommendedRepository(DB.Mysql),
		channelRepo:     _channelRepository.NewChannelRepository(DB.Mysql),
	}
}
