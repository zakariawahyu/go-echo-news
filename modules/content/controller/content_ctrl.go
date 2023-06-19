package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/zakariawahyu/go-echo-news/modules/content"
	"github.com/zakariawahyu/go-echo-news/pkg/exception"
	"github.com/zakariawahyu/go-echo-news/pkg/response"
	"net/http"
)

var (
	limit  = 20
	offset = 0
)

type ContentController struct {
	contentServices content.ContentServices
}

func NewContentController(contentServices content.ContentServices) ContentController {
	return ContentController{
		contentServices: contentServices,
	}
}

func (ctrl *ContentController) Read(ctx echo.Context) error {
	slug := ctx.Param("slug")

	c := ctx.Request().Context()

	content := ctrl.contentServices.GetContentBySlugOrId(c, slug)

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse(http.StatusOK, content))
}

func (ctrl *ContentController) NewsRowHome(ctx echo.Context) error {
	c := ctx.Request().Context()
	err := echo.QueryParamsBinder(ctx).
		Int("limit", &limit).
		Int("offset", &offset).
		BindError()
	exception.PanicIfNeeded(err)

	contents := ctrl.contentServices.GetContentAllHome(c, limit, offset)

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse(http.StatusOK, contents))
}

func (ctrl *ContentController) NewsRowChannel(ctx echo.Context) error {
	c := ctx.Request().Context()
	key := ctx.Param("key")

	err := echo.QueryParamsBinder(ctx).
		Int("limit", &limit).
		Int("offset", &offset).
		BindError()
	exception.PanicIfNeeded(err)

	contents := ctrl.contentServices.GetContentAllChannel(c, key, limit, offset)

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse(http.StatusOK, contents))
}

func (ctrl *ContentController) NewsRowSubChannel(ctx echo.Context) error {
	c := ctx.Request().Context()
	key := ctx.Param("key")

	err := echo.QueryParamsBinder(ctx).
		Int("limit", &limit).
		Int("offset", &offset).
		BindError()
	exception.PanicIfNeeded(err)

	contents := ctrl.contentServices.GetContentAllSubChannel(c, key, limit, offset)

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse(http.StatusOK, contents))
}
