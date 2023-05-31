package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/zakariawahyu/go-echo-news/modules/channel/services"
	"github.com/zakariawahyu/go-echo-news/pkg/response"
	"net/http"
)

type ChannelController struct {
	channelServices services.ChannelServices
}

func NewChannelController(channelServices services.ChannelServices) ChannelController {
	return ChannelController{
		channelServices: channelServices,
	}
}

func (ctrl *ChannelController) AllChannel(c echo.Context) error {
	ctx := c.Request().Context()

	channel := ctrl.channelServices.GetAllChannel(ctx)

	return c.JSON(http.StatusOK, response.NewSuccessResponse(http.StatusOK, channel))
}

func (ctrl *ChannelController) GetChannel(c echo.Context) error {
	slug := c.Param("slug")

	ctx := c.Request().Context()

	channel := ctrl.channelServices.GetChannel(ctx, slug)

	return c.JSON(http.StatusOK, response.NewSuccessResponse(http.StatusOK, channel))
}
