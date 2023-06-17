package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/zakariawahyu/go-echo-news/modules/sub_channel"
	"github.com/zakariawahyu/go-echo-news/pkg/response"
	"net/http"
)

type SubChannelController struct {
	subChannelServices sub_channel.SubChannelServices
}

func NewSubChannelController(subChannelServices sub_channel.SubChannelServices) SubChannelController {
	return SubChannelController{
		subChannelServices: subChannelServices,
	}
}

func (ctrl *SubChannelController) AllSubChannel(ctx echo.Context) error {
	c := ctx.Request().Context()

	subChannels := ctrl.subChannelServices.GetAllSubChannel(c)

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse(http.StatusOK, subChannels))
}

func (ctrl *SubChannelController) SubChannelBySlugOrId(ctx echo.Context) error {
	slug := ctx.Param("slug")
	c := ctx.Request().Context()

	subChannel := ctrl.subChannelServices.GetSubChannelBySlugOrId(c, slug)

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse(http.StatusOK, subChannel))
}
