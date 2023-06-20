package content

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
)

type ContentRepository interface {
	GetBySlugOrId(ctx context.Context, slug string) (*entity.Content, error)
	GetAllHome(ctx context.Context, limit int, offset int) (*[]entity.ContentRowResponse, error)
	GetAllChannel(ctx context.Context, type_id int64, limit int, offset int) (*[]entity.ContentRowResponse, error)
	GetAllSubChannel(ctx context.Context, type_child_id int64, limit int, offset int) (*[]entity.ContentRowResponse, error)
	GetAllRegion(ctx context.Context, type_id int64, limit int, offset int) (*[]entity.ContentRowResponse, error)
	GetAllAds(ctx context.Context, types string, key string, limit int, offset int) (*[]entity.ContentRowResponse, error)
}

type ContentRedisRepository interface {
	GetContent(ctx context.Context, key string) (*entity.Content, error)
	SetContent(ctx context.Context, key string, ttl int, content *entity.Content) error
}
