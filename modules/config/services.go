package config

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
)

type ConfigServices interface {
	GetAllConfig(ctx context.Context) (configs []entity.ConfigResponse)
	GetMetas(ctx context.Context, types string, key string) interface{}
}
