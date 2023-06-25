package server

import (
	"github.com/zakariawahyu/go-echo-news/modules/channel"
	_channelServices "github.com/zakariawahyu/go-echo-news/modules/channel/services"
	"github.com/zakariawahyu/go-echo-news/modules/config"
	_configServices "github.com/zakariawahyu/go-echo-news/modules/config/services"
	"github.com/zakariawahyu/go-echo-news/modules/content"
	_contentServices "github.com/zakariawahyu/go-echo-news/modules/content/services"
	"github.com/zakariawahyu/go-echo-news/modules/item"
	_itemServices "github.com/zakariawahyu/go-echo-news/modules/item/services"
	"github.com/zakariawahyu/go-echo-news/modules/region"
	_regionServices "github.com/zakariawahyu/go-echo-news/modules/region/services"
	"github.com/zakariawahyu/go-echo-news/modules/schedule"
	_scheduleServices "github.com/zakariawahyu/go-echo-news/modules/schedule/services"
	"github.com/zakariawahyu/go-echo-news/modules/sub_channel"
	_subChannelServices "github.com/zakariawahyu/go-echo-news/modules/sub_channel/services"
	"github.com/zakariawahyu/go-echo-news/pkg/logger"
	"time"
)

type Services struct {
	contentServices    content.ContentServices
	channelServices    channel.ChannelServices
	subChannelServices sub_channel.SubChannelServices
	regionServices     region.RegionServices
	configServices     config.ConfigServices
	itemServices       item.ItemServices
	scheduleServices   schedule.ScheduleServices
}

func NewServices(repo *Repository, zapLogger logger.Logger, timeoutContext time.Duration) *Services {
	return &Services{
		contentServices:    _contentServices.NewContentServices(repo.contentRepo, repo.contentRedisRepo, repo.recommendedContentRepo, repo.channelRepo, repo.subChannelRepo, repo.regionRepo, zapLogger, timeoutContext),
		channelServices:    _channelServices.NewChannelServices(repo.channelRepo, repo.channelRedisRepo, zapLogger, timeoutContext),
		subChannelServices: _subChannelServices.NewSubChannelServices(repo.subChannelRepo, zapLogger, timeoutContext),
		regionServices:     _regionServices.NewRegionServices(repo.regionRepo, zapLogger, timeoutContext),
		configServices:     _configServices.NewConfigServices(repo.configRepo, repo.configRedisRep, repo.channelRepo, repo.subChannelRepo, repo.regionRepo, zapLogger, timeoutContext),
		itemServices:       _itemServices.NewItemServices(repo.itemRepo, repo.itemRedisRepo, zapLogger, timeoutContext),
		scheduleServices:   _scheduleServices.NewScheduleServices(repo.scheduleRepo, zapLogger, timeoutContext),
	}
}
