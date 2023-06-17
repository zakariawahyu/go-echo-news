package content

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
)

type ContentServices interface {
	GetContentBySlugOrId(ctx context.Context, slug string) entity.ContentResponse
}
