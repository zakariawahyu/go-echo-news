package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"github.com/zakariawahyu/go-echo-news/config"
	_ctrChannel "github.com/zakariawahyu/go-echo-news/modules/channel/controller"
	_ctrContent "github.com/zakariawahyu/go-echo-news/modules/content/controller"
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

	v1 := e.Group("/v1")
	contentCtrl := _ctrContent.NewContentController(serv.contentServices)
	channelCtrl := _ctrChannel.NewChannelController(serv.channelServices)

	v1.GET("/read/:slug", contentCtrl.Read)

	v1.GET("/channel", channelCtrl.AllChannel)
	v1.GET("/channel/:slug", channelCtrl.GetChannel)

	log.Fatal(e.Start(viper.GetString("APP_ADDRESS")))
}
