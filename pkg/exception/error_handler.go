package exception

import (
	"github.com/gofiber/fiber/v2"
	"github.com/labstack/echo/v4"
	"github.com/zakariawahyu/go-echo-news/pkg/helpers"
	"github.com/zakariawahyu/go-echo-news/pkg/response"
	"net/http"
)

type httpErrorHandler struct {
	statusCode map[error]int
}

func NewHttpErrorHandler(errorStatusCode map[error]int) *httpErrorHandler {
	return &httpErrorHandler{
		statusCode: errorStatusCode,
	}
}

func (self *httpErrorHandler) Handler(err error, ctx echo.Context) {
	he, ok := err.(*echo.HTTPError)
	if ok {
		if he.Internal != nil {
			if herr, ok := he.Internal.(*echo.HTTPError); ok {
				he = herr
			}
		}
	} else {
		he = &echo.HTTPError{
			Code:    helpers.GetStatusCode(err),
			Message: err.Error(),
		}
	}

	code := he.Code
	message := he.Message
	if _, ok := he.Message.(string); ok {
		message = response.ErrorResponse{
			Success: false,
			Code:    code,
			Errors:  err.Error(),
		}
	}

	if !ctx.Response().Committed {
		if ctx.Request().Method == http.MethodHead {
			err = ctx.NoContent(he.Code)
		} else {
			err = ctx.JSON(code, message)
		}
		if err != nil {
			ctx.Echo().Logger.Error(err)
		}
	}
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {

	code := helpers.GetStatusCode(err)
	return ctx.Status(code).JSON(response.ErrorResponse{
		Success: false,
		Code:    code,
		Errors:  err.Error(),
	})
}
