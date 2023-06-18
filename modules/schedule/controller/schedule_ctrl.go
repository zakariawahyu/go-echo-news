package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/zakariawahyu/go-echo-news/modules/schedule"
	"github.com/zakariawahyu/go-echo-news/pkg/response"
	"net/http"
)

type ScheduleController struct {
	scheduleServices schedule.ScheduleServices
}

func NewScheduleController(scheduleServices schedule.ScheduleServices) ScheduleController {
	return ScheduleController{
		scheduleServices: scheduleServices,
	}
}

func (ctrl *ScheduleController) AllLiveStream(ctx echo.Context) error {
	c := ctx.Request().Context()

	schedules := ctrl.scheduleServices.GetAllSchedule(c)

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse(http.StatusOK, schedules))
}

func (ctrl *ScheduleController) LiveStreamBySpecificKey(ctx echo.Context) error {
	key := ctx.Param("key")
	c := ctx.Request().Context()

	schedule := ctrl.scheduleServices.GetScheduleBySpecificKey(c, key)

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse(http.StatusOK, schedule))
}
