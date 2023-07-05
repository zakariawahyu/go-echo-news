package suplemen

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
)

type SuplemenRepository interface {
	GetAll(ctx context.Context) ([]*entity.SuplemenResponse, error)
	GetBySlugOrId(ctx context.Context, slug string) (*entity.SuplemenResponse, error)
}

type SuplemenRedisRepository interface {
	GetAllSuplemen(ctx context.Context, key string) ([]*entity.SuplemenResponse, error)
	GetSuplemen(ctx context.Context, key string) (*entity.SuplemenResponse, error)
	SetALlSuplemen(ctx context.Context, key string, ttl int, suplemen []entity.SuplemenResponse) error
	SetSuplemen(ctx context.Context, key string, ttl int, suplemen *entity.SuplemenResponse) error
}
