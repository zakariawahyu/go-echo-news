package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/zakariawahyu/go-echo-news/modules/content"
	"github.com/zakariawahyu/go-echo-news/pkg/helpers/payload"
	"github.com/zakariawahyu/go-echo-news/pkg/logger"
	"github.com/zakariawahyu/go-echo-news/pkg/response"
	"net/http"
)

type ContentController struct {
	contentServices content.ContentServices
	zapLogger       logger.Logger
}

func NewContentController(contentServices content.ContentServices, zapLogger logger.Logger) ContentController {
	return ContentController{
		contentServices: contentServices,
		zapLogger:       zapLogger,
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
	if err != nil {
		ctrl.zapLogger.Errorf("contentCtrl.NewsRowAll.QueryParamsBinder, err = %s", err)
		panic(err)
	}

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
	if err != nil {
		ctrl.zapLogger.Errorf("contentCtrl.NewsRowAll.QueryParamsBinder, err = %s", err)
		panic(err)
	}

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
	if err != nil {
		ctrl.zapLogger.Errorf("contentCtrl.NewsRowLatestAll.QueryParamsBinder, err = %s", err)
		panic(err)
	}

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
	if err != nil {
		ctrl.zapLogger.Errorf("contentCtrl.NewsRowLatestVideoAll.QueryParamsBinder, err = %s", err)
		panic(err)
	}

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
	if err != nil {
		ctrl.zapLogger.Errorf("contentCtrl.NewsRowLatestPhotoAll.QueryParamsBinder, err = %s", err)
		panic(err)
	}

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
	if err != nil {
		ctrl.zapLogger.Errorf("contentCtrl.NewsRowLatestInfografisAll.QueryParamsBinder, err = %s", err)
		panic(err)
	}

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
	if err != nil {
		ctrl.zapLogger.Errorf("contentCtrl.NewsRowLatestVideoFeaturedAll.QueryParamsBinder, err = %s", err)
		panic(err)
	}

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
	if err != nil {
		ctrl.zapLogger.Errorf("contentCtrl.NewsRowLatestPhotoFeaturedAll.QueryParamsBinder, err = %s", err)
		panic(err)
	}

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
	if err != nil {
		ctrl.zapLogger.Errorf("contentCtrl.NewsRowLatestInfografisFeaturedAll.QueryParamsBinder, err = %s", err)
		panic(err)
	}

	contents := ctrl.contentServices.GetContentAllLatestMultimedia(c, "infografis", true, payloads.Limit, payloads.Offset)

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse(http.StatusOK, contents))
}

func (ctrl *ContentController) HeadlineAll(ctx echo.Context) error {
	c := ctx.Request().Context()
	types := ctx.Param("type")
	key := ctx.Param("key")

	payloads := payload.NewPayload()

	err := echo.QueryParamsBinder(ctx).
		Int("limit", &payloads.Limit).
		Int("offset", &payloads.Offset).
		BindError()
	if err != nil {
		ctrl.zapLogger.Errorf("contentCtrl.HeadlineAll.QueryParamsBinder, err = %s", err)
		panic(err)
	}

	contents := ctrl.contentServices.GetContentAllHeadline(c, types, key, payloads.Limit, payloads.Offset)

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse(http.StatusOK, contents))
}

func (ctrl *ContentController) HeadlineAdsAll(ctx echo.Context) error {
	c := ctx.Request().Context()
	types := ctx.Param("type")
	key := ctx.Param("key")

	payloads := payload.NewPayload()

	err := echo.QueryParamsBinder(ctx).
		Int("limit", &payloads.Limit).
		Int("offset", &payloads.Offset).
		BindError()
	if err != nil {
		ctrl.zapLogger.Errorf("contentCtrl.HeadlineAdsAll.QueryParamsBinder, err = %s", err)
		panic(err)
	}

	contents := ctrl.contentServices.GetContentAllHeadlineAds(c, types, key, payloads.Limit, payloads.Offset)

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse(http.StatusOK, contents))
}

func (ctrl *ContentController) MultimediaRowVideoAll(ctx echo.Context) error {
	c := ctx.Request().Context()
	types := ctx.Param("type")
	key := ctx.Param("key")

	payloads := payload.NewPayload()

	err := echo.QueryParamsBinder(ctx).
		Int("limit", &payloads.Limit).
		Int("offset", &payloads.Offset).
		BindError()
	if err != nil {
		ctrl.zapLogger.Errorf("contentCtrl.MultimediaRowAll.QueryParamsBinder, err = %s", err)
		panic(err)
	}

	contents := ctrl.contentServices.GetContentAllMultimediaRow(c, "video", types, key, payloads.Limit, payloads.Offset)

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse(http.StatusOK, contents))
}

func (ctrl *ContentController) MultimediaRowPhotoAll(ctx echo.Context) error {
	c := ctx.Request().Context()
	types := ctx.Param("type")
	key := ctx.Param("key")

	payloads := payload.NewPayload()

	err := echo.QueryParamsBinder(ctx).
		Int("limit", &payloads.Limit).
		Int("offset", &payloads.Offset).
		BindError()
	if err != nil {
		ctrl.zapLogger.Errorf("contentCtrl.MultimediaRowPhotoAll.QueryParamsBinder, err = %s", err)
		panic(err)
	}

	contents := ctrl.contentServices.GetContentAllMultimediaRow(c, "photo", types, key, payloads.Limit, payloads.Offset)

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse(http.StatusOK, contents))
}

func (ctrl *ContentController) MultimediaRowInfografisAll(ctx echo.Context) error {
	c := ctx.Request().Context()
	types := ctx.Param("type")
	key := ctx.Param("key")

	payloads := payload.NewPayload()

	err := echo.QueryParamsBinder(ctx).
		Int("limit", &payloads.Limit).
		Int("offset", &payloads.Offset).
		BindError()
	if err != nil {
		ctrl.zapLogger.Errorf("contentCtrl.MultimediaRowInfografisAll.QueryParamsBinder, err = %s", err)
		panic(err)
	}

	contents := ctrl.contentServices.GetContentAllMultimediaRow(c, "infografis", types, key, payloads.Limit, payloads.Offset)

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse(http.StatusOK, contents))
}

func (ctrl *ContentController) ArticleRowAll(ctx echo.Context) error {
	c := ctx.Request().Context()

	payloads := payload.NewPayload()

	err := echo.QueryParamsBinder(ctx).
		Int("limit", &payloads.Limit).
		Int("offset", &payloads.Offset).
		BindError()
	if err != nil {
		ctrl.zapLogger.Errorf("contentCtrl.ArticleRowAll.QueryParamsBinder, err = %s", err)
		panic(err)
	}

	contents := ctrl.contentServices.GetContentAllArticleRow(c, payloads.Limit, payloads.Offset)

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse(http.StatusOK, contents))
}

func (ctrl *ContentController) EditorChoiceRowAll(ctx echo.Context) error {
	c := ctx.Request().Context()

	payloads := payload.NewPayload()

	err := echo.QueryParamsBinder(ctx).
		Int("limit", &payloads.Limit).
		Int("offset", &payloads.Offset).
		BindError()
	if err != nil {
		ctrl.zapLogger.Errorf("contentCtrl.EditorChoiceRowAll.QueryParamsBinder, err = %s", err)
		panic(err)
	}

	contents := ctrl.contentServices.GetContentAllEditorChoiceRow(c, payloads.Limit, payloads.Offset)

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse(http.StatusOK, contents))
}

func (ctrl *ContentController) IndeksRowAll(ctx echo.Context) error {
	c := ctx.Request().Context()
	types := ctx.Param("type")
	key := ctx.Param("key")
	date := ctx.Param("date")

	payloads := payload.NewPayload()

	err := echo.QueryParamsBinder(ctx).
		Int("limit", &payloads.Limit).
		Int("offset", &payloads.Offset).
		BindError()
	if err != nil {
		ctrl.zapLogger.Errorf("contentCtrl.IndeksRowAll.QueryParamsBinder, err = %s", err)
		panic(err)
	}

	contents := ctrl.contentServices.GetContentAllIndeksRow(c, types, key, date, payloads.Limit, payloads.Offset)

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse(http.StatusOK, contents))
}

func (ctrl *ContentController) SearchRowAll(ctx echo.Context) error {
	c := ctx.Request().Context()
	types := ctx.Param("type")
	key := ctx.Param("key")

	payloads := payload.NewPayload()

	err := echo.QueryParamsBinder(ctx).
		Int("limit", &payloads.Limit).
		Int("offset", &payloads.Offset).
		BindError()
	if err != nil {
		ctrl.zapLogger.Errorf("contentCtrl.SearchAll.QueryParamsBinder, err = %s", err)
		panic(err)
	}

	contents := ctrl.contentServices.GetContentAllSearchRow(c, types, key, payloads.Limit, payloads.Offset)

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse(http.StatusOK, contents))
}
