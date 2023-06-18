package server

import (
	"github.com/zakariawahyu/go-echo-news/modules/channel"
	_channelRepository "github.com/zakariawahyu/go-echo-news/modules/channel/repository"
	"github.com/zakariawahyu/go-echo-news/modules/config"
	_configRepo "github.com/zakariawahyu/go-echo-news/modules/config/repository"
	"github.com/zakariawahyu/go-echo-news/modules/content"
	_contentRepository "github.com/zakariawahyu/go-echo-news/modules/content/repository"
	"github.com/zakariawahyu/go-echo-news/modules/item"
	_itemRepo "github.com/zakariawahyu/go-echo-news/modules/item/repository"
	"github.com/zakariawahyu/go-echo-news/modules/recommended_content"
	_recommendedContentRepository "github.com/zakariawahyu/go-echo-news/modules/recommended_content/repository"
	"github.com/zakariawahyu/go-echo-news/modules/region"
	_regionRepo "github.com/zakariawahyu/go-echo-news/modules/region/repository"
	"github.com/zakariawahyu/go-echo-news/modules/sub_channel"
	_subChannelRepository "github.com/zakariawahyu/go-echo-news/modules/sub_channel/repository"
	"github.com/zakariawahyu/go-echo-news/pkg/db"
)

type Repository struct {
	contentRepo            content.ContentRepository
	contentRedisRepo       content.ContentRedisRepository
	recommendedContentRepo recommended_content.RecommendedContentRepository
	channelRepo            channel.ChannelRepository
	channelRedisRepo       channel.ChannelRedisRepository
	subChannelRepo         sub_channel.SubChannelRepository
	regionRepo             region.RegionRepository
	configRepo             config.ConfigRepository
	itemRepo               item.ItemRepository
}

func NewRepository(DB *db.Conn) *Repository {
	return &Repository{
		contentRepo:            _contentRepository.NewContentRepository(DB.Mysql),
		contentRedisRepo:       _contentRepository.NewContentRedisRepository(DB.Redis),
		recommendedContentRepo: _recommendedContentRepository.NewRecommendedContentRepository(DB.Mysql),
		channelRepo:            _channelRepository.NewChannelRepository(DB.Mysql),
		channelRedisRepo:       _channelRepository.NewChannelRedisRepository(DB.Redis),
		subChannelRepo:         _subChannelRepository.NewSubChannelRepository(DB.Mysql),
		regionRepo:             _regionRepo.NewRegionRepository(DB.Mysql),
		configRepo:             _configRepo.NewConfigRepository(DB.Mysql),
		itemRepo:               _itemRepo.NewItemRepository(DB.Mysql),
	}
}
