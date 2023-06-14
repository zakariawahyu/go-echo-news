package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/zakariawahyu/go-echo-news/modules/channel"
	"github.com/zakariawahyu/go-echo-news/pkg/response"
	"net/http"
)

type ChannelController struct {
	channelServices channel.ChannelServices
}

func NewChannelController(channelServices channel.ChannelServices) ChannelController {
	return ChannelController{
		channelServices: channelServices,
	}
}

func (ctrl *ChannelController) AllChannel(ctx echo.Context) error {
	c := ctx.Request().Context()

	channel := ctrl.channelServices.GetAllChannel(c)

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse(http.StatusOK, channel))
}

func (ctrl *ChannelController) GetChannel(ctx echo.Context) error {
	slug := ctx.Param("slug")

	c := ctx.Request().Context()

	channel := ctrl.channelServices.GetChannel(c, slug)

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse(http.StatusOK, channel))
}
