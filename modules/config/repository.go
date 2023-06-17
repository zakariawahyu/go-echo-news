package config

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
)

type ConfigRepository interface {
	GetAll(ctx context.Context) (*[]entity.Config, error)
}
