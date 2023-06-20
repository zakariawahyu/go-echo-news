package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"github.com/zakariawahyu/go-echo-news/config"
	_channelController "github.com/zakariawahyu/go-echo-news/modules/channel/controller"
	_configController "github.com/zakariawahyu/go-echo-news/modules/config/controller"
	_contentController "github.com/zakariawahyu/go-echo-news/modules/content/controller"
	_itemController "github.com/zakariawahyu/go-echo-news/modules/item/controller"
	_regionController "github.com/zakariawahyu/go-echo-news/modules/region/controller"
	_scheduleController "github.com/zakariawahyu/go-echo-news/modules/schedule/controller"
	_subChannelController "github.com/zakariawahyu/go-echo-news/modules/sub_channel/controller"
	"github.com/zakariawahyu/go-echo-news/pkg/exception"
	"log"
	"net/http"
)

func NewAppHandler(e *echo.Echo) {
	cfg := config.NewConfig()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"app":         cfg.App.Name,
			"version":     cfg.App.Version,
			"app_timeout": cfg.App.ContextTimeout,
		})
	})
}

func NewHandler(cfg *config.Config, serv *Services) {
	e := echo.New()

	e.HTTPErrorHandler = exception.NewHttpErrorHandler(exception.NewErrorStatusCodeMaps()).Handler
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10,
	}))
	NewAppHandler(e)

	contentController := _contentController.NewContentController(serv.contentServices)
	channelController := _channelController.NewChannelController(serv.channelServices)
	subChannelController := _subChannelController.NewSubChannelController(serv.subChannelServices)
	regionController := _regionController.NewRegionController(serv.regionServices)
	configController := _configController.NewConfigController(serv.configServices)
	itemController := _itemController.NewItemController(serv.itemServices)
	scheduleController := _scheduleController.NewScheduleController(serv.scheduleServices)

	v1 := e.Group("/v1")
	v2 := e.Group("/v2")

	v1.GET("/read/:slug", contentController.Read)

	v1.GET("/channel", channelController.AllChannel)
	v1.GET("/channel/:slug", channelController.ChannelBySlugOrId)

	v1.GET("/sub-channel", subChannelController.AllSubChannel)
	v1.GET("/sub-channel/:slug", subChannelController.SubChannelBySlugOrId)

	v1.GET("/region", regionController.AllRegion)
	v1.GET("/region/:slug", regionController.RegionBySlugOrId)

	// Replica API iNews.id
	v2.GET("/config", configController.AllConfig)
	v2.GET("/meta/:type/:key", configController.Metas)
	v2.GET("/item/:type", itemController.ItemByType)
	v2.GET("/live_stream", scheduleController.AllLiveStream)
	v2.GET("/live_stream/:key", scheduleController.LiveStreamBySpecificKey)

	v2.GET("/news-row", contentController.NewsRowHome)
	v2.GET("/news-row/channel/:key", contentController.NewsRowChannel)
	v2.GET("/news-row/subchannel/:key", contentController.NewsRowSubChannel)
	v2.GET("/news-row/region/:key", contentController.NewsRowRegion)
	v2.GET("/news-row/ads", contentController.NewsRowAdsAll)
	v2.GET("/news-row/ads/:type/:key", contentController.NewsRowAdsAll)

	log.Fatal(e.Start(viper.GetString("APP_ADDRESS")))
}
