package content

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
)

type ContentServices interface {
	GetContentBySlugOrId(ctx context.Context, slug string) entity.ContentResponse
	GetContentAllHome(ctx context.Context, limit int, offset int) (contents []entity.ContentRowResponse)
}
