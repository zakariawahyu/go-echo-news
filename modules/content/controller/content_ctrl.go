package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/zakariawahyu/go-echo-news/modules/content"
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

	content := ctrl.contentServices.GetContent(c, slug)

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse(http.StatusOK, content))
}
