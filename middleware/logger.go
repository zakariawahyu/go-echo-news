package middleware

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/zakariawahyu/go-echo-news/config"
	"github.com/zakariawahyu/go-echo-news/pkg/helpers"
	"github.com/zakariawahyu/go-echo-news/pkg/logger"
	"time"
)

type MiddlewareLogger struct {
	cfg    *config.Config
	logger logger.Logger
}

func NewMiddlewareLogger(cfg *config.Config, logger logger.Logger) *MiddlewareLogger {
	return &MiddlewareLogger{cfg: cfg, logger: logger}
}

func (mw *MiddlewareLogger) LoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		start := time.Now()
		err := next(ctx)
		if err != nil {
			ctx.Error(err)
		}

		req := ctx.Request()
		res := ctx.Response()
		status := res.Status
		size := res.Size
		s := time.Since(start).String()
		requestID := helpers.GetRequestID(ctx)

		errors := fmt.Sprintf("Err: %s, RequestID: %s, Method: %s, URI: %s, Status: %v, Size: %v, Time: %s",
			err, requestID, req.Method, req.URL, status, size, s)

		switch {
		case status >= 500:
			mw.logger.Error("SERVER ERROR = ", errors)
		case status >= 400:
			mw.logger.Error("CLIENT ERROR = ", errors)
		case status >= 300:
			mw.logger.Info("REDIRECTION = ", errors)
		default:
			mw.logger.Infof("RequestID: %s, Method: %s, URI: %s, Status: %v, Size: %v, Time: %s",
				requestID, req.Method, req.URL, status, size, s)
		}

		return err
	}
}
