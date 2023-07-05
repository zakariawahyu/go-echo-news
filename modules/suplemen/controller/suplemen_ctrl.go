package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/zakariawahyu/go-echo-news/modules/suplemen"
	"github.com/zakariawahyu/go-echo-news/pkg/response"
	"net/http"
)

type SuplemenController struct {
	suplemenServices suplemen.SuplemenServices
}

func NewSuplemenController(suplemenServices suplemen.SuplemenServices) SuplemenController {
	return SuplemenController{
		suplemenServices: suplemenServices,
	}
}

func (ctrl *SuplemenController) AllSuplemen(ctx echo.Context) error {
	c := ctx.Request().Context()

	suplemens := ctrl.suplemenServices.GetAllSuplemen(c)

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse(http.StatusOK, suplemens))
}

func (ctrl *SuplemenController) SuplemenBySlugOrId(ctx echo.Context) error {
	slug := ctx.Param("slug")
	c := ctx.Request().Context()

	suplemen := ctrl.suplemenServices.GetSuplemenBySlugOrId(c, slug)

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse(http.StatusOK, suplemen))
}
