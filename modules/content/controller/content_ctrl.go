package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/zakariawahyu/go-echo-news/modules/content/services"
	"github.com/zakariawahyu/go-echo-news/pkg/response"
	"net/http"
)

type ContentController struct {
	contentServices services.ContentServices
}

func NewContentController(contentServices services.ContentServices) ContentController {
	return ContentController{
		contentServices: contentServices,
	}
}

func (ctrl *ContentController) Read(c echo.Context) error {
	slug := c.Param("slug")

	ctx := c.Request().Context()

	content := ctrl.contentServices.GetContent(ctx, slug)

	return c.JSON(http.StatusOK, response.NewSuccessResponse(http.StatusOK, content))
}