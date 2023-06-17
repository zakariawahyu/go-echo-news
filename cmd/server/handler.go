package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"github.com/zakariawahyu/go-echo-news/config"
	_channelController "github.com/zakariawahyu/go-echo-news/modules/channel/controller"
	_contentController "github.com/zakariawahyu/go-echo-news/modules/content/controller"
	_regionController "github.com/zakariawahyu/go-echo-news/modules/region/controller"
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

	contentCtrl := _contentController.NewContentController(serv.contentServices)
	channelCtrl := _channelController.NewChannelController(serv.channelServices)
	subChannelCtrl := _subChannelController.NewSubChannelController(serv.subChannelServices)
	regionCtrl := _regionController.NewRegionController(serv.regionServices)

	v1 := e.Group("/v1")

	v1.GET("/read/:slug", contentCtrl.Read)

	v1.GET("/channel", channelCtrl.AllChannel)
	v1.GET("/channel/:slug", channelCtrl.GetChannel)

	v1.GET("/sub-channel", subChannelCtrl.AllSubChannel)
	v1.GET("/sub-channel/:slug", subChannelCtrl.SubChannelBySlugOrId)

	v1.GET("/region", regionCtrl.AllRegion)
	v1.GET("/region/:slug", regionCtrl.RegionBySlugOrId)

	log.Fatal(e.Start(viper.GetString("APP_ADDRESS")))
}
