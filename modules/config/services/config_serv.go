package services

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/channel"
	"github.com/zakariawahyu/go-echo-news/modules/config"
	"github.com/zakariawahyu/go-echo-news/modules/region"
	"github.com/zakariawahyu/go-echo-news/modules/sub_channel"
	"github.com/zakariawahyu/go-echo-news/pkg/exception"
	"time"
)

type configServices struct {
	configRepo       config.ConfigRepository
	channelRepo      channel.ChannelRepository
	subChannelRepo   sub_channel.SubChannelRepository
	regionRepository region.RegionRepository
	contextTimeout   time.Duration
}

func NewConfigServices(configRepo config.ConfigRepository, channelRepo channel.ChannelRepository, subChannelRepo sub_channel.SubChannelRepository, regionRepository region.RegionRepository, timeout time.Duration) config.ConfigServices {
	return &configServices{
		configRepo:       configRepo,
		channelRepo:      channelRepo,
		subChannelRepo:   subChannelRepo,
		regionRepository: regionRepository,
		contextTimeout:   timeout,
	}
}

func (serv *configServices) GetAllConfig(ctx context.Context) (configs []entity.ConfigResponse) {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	res, err := serv.configRepo.GetAll(c)
	exception.PanicIfNeeded(err)

	for _, config := range *res {
		configs = append(configs, entity.NewConfigResponse(&config))
	}

	return configs
}

func (serv *configServices) GetMetas(ctx context.Context, types string, key string) interface{} {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	var data interface{}
	var err error

	if types == "channel" {
		data, err = serv.channelRepo.GetBySlugOrId(c, key)
		exception.PanicIfNeeded(err)
	} else if types == "subchannel" {
		data, err = serv.subChannelRepo.GetBySlugOrId(c, key)
	} else if types == "region" {
		data, err = serv.regionRepository.GetBySlugOrId(c, key)
	} else {
		data = OtherMeta(key)
	}

	return data
}

func OtherMeta(key string) interface{} {
	data := echo.Map{
		"homepage": echo.Map{
			"title":       "Berita Terkini dan Informasi Terbaru Hari Ini",
			"description": "iNews ID - Situs portal berita nasional dan daerah yang menyajikan informasi terkini dan terbaru seperti, Berita Politik, Hukum, Keuangan, Teknologi",
		},
	}

	return data[key]
}
