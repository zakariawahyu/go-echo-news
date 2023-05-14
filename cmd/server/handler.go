package server

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"github.com/zakariawahyu/go-echo-news/config"
	"github.com/zakariawahyu/go-echo-news/modules/content/controller"
	"github.com/zakariawahyu/go-echo-news/utils/exception"
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

	v1 := e.Group("/v1")

	NewAppHandler(e)
	controller.NewContentController(v1, serv.contentServices)

	log.Fatal(e.Start(viper.GetString("APP_ADDRESS")))
}
