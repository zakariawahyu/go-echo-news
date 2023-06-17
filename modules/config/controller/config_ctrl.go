package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/zakariawahyu/go-echo-news/modules/config"
	"github.com/zakariawahyu/go-echo-news/pkg/response"
	"net/http"
)

type ConfigController struct {
	configServices config.ConfigServices
}

func NewConfigController(configServices config.ConfigServices) ConfigController {
	return ConfigController{
		configServices: configServices,
	}
}

func (ctrl *ConfigController) AllConfig(ctx echo.Context) error {
	c := ctx.Request().Context()

	configs := ctrl.configServices.GetAllConfig(c)

	myMap := make(map[string]interface{})

	for _, config := range configs {
		myMap[config.Key] = config.Value
	}

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse(http.StatusOK, myMap))
}

func (ctrl *ConfigController) Metas(ctx echo.Context) error {
	c := ctx.Request().Context()
	types := ctx.Param("type")
	key := ctx.Param("key")

	data := ctrl.configServices.GetMetas(c, types, key)

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse(http.StatusOK, data))
}
