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

func NewServices(repo *Repository, timeoutContext time.Duration) *Services {
	return &Services{
		contentServices:    _contentServices.NewContentServices(repo.contentRepo, repo.contentRedisRepo, repo.recommendedContentRepo, repo.channelRepo, repo.subChannelRepo, timeoutContext),
		channelServices:    _channelServices.NewChannelServices(repo.channelRepo, repo.channelRedisRepo, timeoutContext),
		subChannelServices: _subChannelServices.NewSubChannelServices(repo.subChannelRepo, timeoutContext),
		regionServices:     _regionServices.NewRegionServices(repo.regionRepo, timeoutContext),
		configServices:     _configServices.NewConfigServices(repo.configRepo, repo.channelRepo, repo.subChannelRepo, repo.regionRepo, timeoutContext),
		itemServices:       _itemServices.NewItemServices(repo.itemRepo, timeoutContext),
		scheduleServices:   _scheduleServices.NewScheduleServices(repo.scheduleRepo, timeoutContext),
	}
}
