package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/zakariawahyu/go-echo-news/modules/content"
	"github.com/zakariawahyu/go-echo-news/pkg/exception"
	"github.com/zakariawahyu/go-echo-news/pkg/helpers/payload"
	"github.com/zakariawahyu/go-echo-news/pkg/response"
	"net/http"
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

func (ctrl *ContentController) NewsRowAll(ctx echo.Context) error {
	c := ctx.Request().Context()
	types := ctx.Param("type")
	key := ctx.Param("key")

	payloads := payload.NewPayload()

	err := echo.QueryParamsBinder(ctx).
		Int("limit", &payloads.Limit).
		Int("offset", &payloads.Offset).
		BindError()
	exception.PanicIfNeeded(err)

	contents := ctrl.contentServices.GetContentAllRow(c, types, key, payloads.Limit, payloads.Offset)

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse(http.StatusOK, contents))
}

func (ctrl *ContentController) NewsRowAdsAll(ctx echo.Context) error {
	c := ctx.Request().Context()
	types := ctx.Param("type")
	key := ctx.Param("key")

	payloads := payload.NewPayload()

	err := echo.QueryParamsBinder(ctx).
		Int("limit", &payloads.Limit).
		Int("offset", &payloads.Offset).
		BindError()
	exception.PanicIfNeeded(err)

	contents := ctrl.contentServices.GetContentAllRowAds(c, types, key, payloads.Limit, payloads.Offset)

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse(http.StatusOK, contents))
}

func (ctrl *ContentController) NewsRowLatestAll(ctx echo.Context) error {
	c := ctx.Request().Context()
	types := ctx.Param("type")
	key := ctx.Param("key")

	payloads := payload.NewPayload()

	err := echo.QueryParamsBinder(ctx).
		Int("limit", &payloads.Limit).
		Int("offset", &payloads.Offset).
		BindError()
	exception.PanicIfNeeded(err)

	contents := ctrl.contentServices.GetContentAllLatest(c, types, key, payloads.Limit, payloads.Offset)

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse(http.StatusOK, contents))
}

func (ctrl *ContentController) NewsRowLatestVideoAll(ctx echo.Context) error {
	c := ctx.Request().Context()
	payloads := payload.NewPayload()

	err := echo.QueryParamsBinder(ctx).
		Int("limit", &payloads.Limit).
		Int("offset", &payloads.Offset).
		BindError()
	exception.PanicIfNeeded(err)

	contents := ctrl.contentServices.GetContentAllLatestMultimedia(c, "video", false, payloads.Limit, payloads.Offset)

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse(http.StatusOK, contents))
}

func (ctrl *ContentController) NewsRowLatestPhotoAll(ctx echo.Context) error {
	c := ctx.Request().Context()
	payloads := payload.NewPayload()

	err := echo.QueryParamsBinder(ctx).
		Int("limit", &payloads.Limit).
		Int("offset", &payloads.Offset).
		BindError()
	exception.PanicIfNeeded(err)

	contents := ctrl.contentServices.GetContentAllLatestMultimedia(c, "photo", false, payloads.Limit, payloads.Offset)

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse(http.StatusOK, contents))
}

func (ctrl *ContentController) NewsRowLatestInfografisAll(ctx echo.Context) error {
	c := ctx.Request().Context()
	payloads := payload.NewPayload()

	err := echo.QueryParamsBinder(ctx).
		Int("limit", &payloads.Limit).
		Int("offset", &payloads.Offset).
		BindError()
	exception.PanicIfNeeded(err)

	contents := ctrl.contentServices.GetContentAllLatestMultimedia(c, "infografis", false, payloads.Limit, payloads.Offset)

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse(http.StatusOK, contents))
}

func (ctrl *ContentController) NewsRowLatestVideoFeaturedAll(ctx echo.Context) error {
	c := ctx.Request().Context()
	payloads := payload.NewPayload()

	err := echo.QueryParamsBinder(ctx).
		Int("limit", &payloads.Limit).
		Int("offset", &payloads.Offset).
		BindError()
	exception.PanicIfNeeded(err)

	contents := ctrl.contentServices.GetContentAllLatestMultimedia(c, "video", true, payloads.Limit, payloads.Offset)

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse(http.StatusOK, contents))
}

func (ctrl *ContentController) NewsRowLatestPhotoFeaturedAll(ctx echo.Context) error {
	c := ctx.Request().Context()
	payloads := payload.NewPayload()

	err := echo.QueryParamsBinder(ctx).
		Int("limit", &payloads.Limit).
		Int("offset", &payloads.Offset).
		BindError()
	exception.PanicIfNeeded(err)

	contents := ctrl.contentServices.GetContentAllLatestMultimedia(c, "photo", true, payloads.Limit, payloads.Offset)

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse(http.StatusOK, contents))
}

func (ctrl *ContentController) NewsRowLatestInfografisFeaturedAll(ctx echo.Context) error {
	c := ctx.Request().Context()
	payloads := payload.NewPayload()

	err := echo.QueryParamsBinder(ctx).
		Int("limit", &payloads.Limit).
		Int("offset", &payloads.Offset).
		BindError()
	exception.PanicIfNeeded(err)

	contents := ctrl.contentServices.GetContentAllLatestMultimedia(c, "infografis", true, payloads.Limit, payloads.Offset)

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse(http.StatusOK, contents))
}
