package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/zakariawahyu/go-echo-news/modules/region"
	"github.com/zakariawahyu/go-echo-news/pkg/response"
	"net/http"
)

type RegionController struct {
	regionServices region.RegionServices
}

func NewRegionController(regionServices region.RegionServices) RegionController {
	return RegionController{
		regionServices: regionServices,
	}
}

func (ctrl *RegionController) AllRegion(ctx echo.Context) error {
	c := ctx.Request().Context()

	regions := ctrl.regionServices.GetAllRegion(c)

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse(http.StatusOK, regions))
}

func (ctrl *RegionController) RegionBySlugOrId(ctx echo.Context) error {
	slug := ctx.Param("slug")
	c := ctx.Request().Context()

	region := ctrl.regionServices.GetRegionBySlugOrId(c, slug)

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse(http.StatusOK, region))
}
