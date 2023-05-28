package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-echo-news/pkg/exception"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}
