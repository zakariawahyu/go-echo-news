package content

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
)

type ContentRepository interface {
	GetBySlugOrId(ctx context.Context, slug string) (*entity.Content, error)
}

type ContentRedisRepository interface {
	GetContent(ctx context.Context, key string) (*entity.Content, error)
	SetContent(ctx context.Context, key string, ttl int, content *entity.Content) error
}
