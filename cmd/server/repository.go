package server

import (
	"github.com/zakariawahyu/go-echo-news/modules/channel"
	_channelRepository "github.com/zakariawahyu/go-echo-news/modules/channel/repository"
	"github.com/zakariawahyu/go-echo-news/modules/config"
	_configRepo "github.com/zakariawahyu/go-echo-news/modules/config/repository"
	"github.com/zakariawahyu/go-echo-news/modules/content"
	_contentRepository "github.com/zakariawahyu/go-echo-news/modules/content/repository"
	"github.com/zakariawahyu/go-echo-news/modules/content_has"
	_contentHas "github.com/zakariawahyu/go-echo-news/modules/content_has/repository"
	"github.com/zakariawahyu/go-echo-news/modules/item"
	_itemRepo "github.com/zakariawahyu/go-echo-news/modules/item/repository"
	"github.com/zakariawahyu/go-echo-news/modules/recommended_content"
	_recommendedContentRepository "github.com/zakariawahyu/go-echo-news/modules/recommended_content/repository"
	"github.com/zakariawahyu/go-echo-news/modules/region"
	_regionRepo "github.com/zakariawahyu/go-echo-news/modules/region/repository"
	"github.com/zakariawahyu/go-echo-news/modules/schedule"
	_scheduleRepo "github.com/zakariawahyu/go-echo-news/modules/schedule/repository"
	"github.com/zakariawahyu/go-echo-news/modules/sub_channel"
	_subChannelRepository "github.com/zakariawahyu/go-echo-news/modules/sub_channel/repository"
	"github.com/zakariawahyu/go-echo-news/modules/suplemen"
	_suplemenRepo "github.com/zakariawahyu/go-echo-news/modules/suplemen/repository"
	"github.com/zakariawahyu/go-echo-news/modules/tag"
	_tagRepo "github.com/zakariawahyu/go-echo-news/modules/tag/repository"
	"github.com/zakariawahyu/go-echo-news/modules/topic"
	_topicRepo "github.com/zakariawahyu/go-echo-news/modules/topic/repository"
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
	configRedisRep         config.ConfigRedisRepository
	itemRepo               item.ItemRepository
	itemRedisRepo          item.ItemRedisRepository
	scheduleRepo           schedule.ScheduleRepository
	scheduleRedisRepo      schedule.ScheduleRedisRepository
	suplemenRepo           suplemen.SuplemenRepository
	suplemenRedisRepo      suplemen.SuplemenRedisRepository
	tagRepo                tag.TagRepository
	topicRepo              topic.TopicRepository
	contentHasTag          content_has.ContentHasTagRepository
	contentHasTopic        content_has.ContentHasTopicRepository
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
		configRedisRep:         _configRepo.NewConfigRedisRepo(DB.Redis),
		itemRepo:               _itemRepo.NewItemRepository(DB.Mysql),
		itemRedisRepo:          _itemRepo.NewItemRedisRepository(DB.Redis),
		scheduleRepo:           _scheduleRepo.NewScheduleRepository(DB.Mysql),
		scheduleRedisRepo:      _scheduleRepo.NewScheduleRedisRepository(DB.Redis),
		suplemenRepo:           _suplemenRepo.NewSuplemenRepository(DB.Mysql),
		suplemenRedisRepo:      _suplemenRepo.NewSuplemenRedisRepository(DB.Redis),
		tagRepo:                _tagRepo.NewTagRepository(DB.Mysql),
		topicRepo:              _topicRepo.NewTopicRepository(DB.Mysql),
		contentHasTag:          _contentHas.NewContentHasTagRepository(DB.Mysql),
		contentHasTopic:        _contentHas.NewContentHasTopicRepository(DB.Mysql),
	}
}
