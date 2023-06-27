package content

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
)

type ContentRepository interface {
	GetBySlugOrId(ctx context.Context, slug string) (*entity.Content, error)
	GetAllRow(ctx context.Context, types string, key string, limit int, offset int) ([]*entity.ContentRowResponse, error)
	GetAllRowAds(ctx context.Context, types string, key string, limit int, offset int) ([]*entity.ContentRowResponse, error)
	GetAllLatest(ctx context.Context, types string, key string, limit int, offset int) ([]*entity.ContentRowResponse, error)
	GetAllLatestMultimedia(ctx context.Context, types string, featured bool, limit int, offset int) ([]*entity.ContentRowResponse, error)
	GetAllHeadline(ctx context.Context, types string, key string, limit int, offset int) ([]*entity.ContentRowResponse, error)
	GetAllHeadlineAds(ctx context.Context, types string, key string, limit int, offset int) ([]*entity.ContentRowResponse, error)
	GetAllMultimediaRow(ctx context.Context, multimediaType string, types string, key string, limit int, offset int) ([]*entity.ContentRowResponse, error)
}

type ContentRedisRepository interface {
	GetAllContentRow(ctx context.Context, key string) ([]*entity.ContentRowResponse, error)
	GetContent(ctx context.Context, key string) (*entity.Content, error)
	SetALlContentRow(ctx context.Context, key string, ttl int, content []entity.ContentRowResponse) error
	SetContent(ctx context.Context, key string, ttl int, content *entity.Content) error
}
