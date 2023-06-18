package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/zakariawahyu/go-echo-news/modules/item"
	"github.com/zakariawahyu/go-echo-news/pkg/response"
	"net/http"
)

type ItemController struct {
	itemServices item.ItemServices
}

func NewItemController(itemServices item.ItemServices) ItemController {
	return ItemController{
		itemServices: itemServices,
	}
}

func (ctrl *ItemController) ItemByType(ctx echo.Context) error {
	c := ctx.Request().Context()
	types := ctx.Param("type")

	item := ctrl.itemServices.GetItemByTypes(c, types)

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse(http.StatusOK, item))
}
