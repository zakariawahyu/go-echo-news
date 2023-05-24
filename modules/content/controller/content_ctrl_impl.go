package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/zakariawahyu/go-echo-news/modules/content/services"
	"github.com/zakariawahyu/go-echo-news/utils/response"
	"net/http"
)

type ContentControllerImpl struct {
	contentServices services.ContentServices
}

func NewContentController(e *echo.Group, serv services.ContentServices) {
	content := &ContentControllerImpl{
		contentServices: serv,
	}
	e.GET("/read/:slug", content.GetBySlug)
}

func (ctrl *ContentControllerImpl) GetBySlug(c echo.Context) error {
	slug := c.Param("slug")

	ctx := c.Request().Context()

	content := ctrl.contentServices.GetBySlug(ctx, slug)

	return c.JSON(http.StatusOK, &response.SuccessResponse{
		Success: true,
		Code:    http.StatusOK,
		Data:    content,
	})
}
