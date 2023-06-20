package content

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
)

type ContentServices interface {
	GetContentBySlugOrId(ctx context.Context, slug string) entity.ContentResponse
	GetContentAllHome(ctx context.Context, limit int, offset int) (contents []entity.ContentRowResponse)
	GetContentAllChannel(ctx context.Context, key string, limit int, offset int) (contents []entity.ContentRowResponse)
	GetContentAllSubChannel(ctx context.Context, key string, limit int, offset int) (contents []entity.ContentRowResponse)
	GetContentAllRegion(ctx context.Context, key string, limit int, offset int) (contents []entity.ContentRowResponse)
	GetContentAllAds(ctx context.Context, types string, key string, limit int, offset int) (contents []entity.ContentRowResponse)
}
