package content

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
)

type ContentServices interface {
	GetContentBySlugOrId(ctx context.Context, slug string) entity.ContentResponse
	GetContentAllRow(ctx context.Context, types string, key string, limit int, offset int) (contents []entity.ContentRowResponse)
	GetContentAllRowAds(ctx context.Context, types string, key string, limit int, offset int) (contents []entity.ContentRowResponse)
	GetContentAllLatest(ctx context.Context, types string, key string, limit int, offset int) (contents []entity.ContentRowResponse)
	GetContentAllLatestMultimedia(ctx context.Context, types string, featured bool, limit int, offset int) (contents []entity.ContentRowResponse)
	GetContentAllHeadline(ctx context.Context, types string, key string, limit int, offset int) (contents []entity.ContentRowResponse)
	GetContentAllHeadlineAds(ctx context.Context, types string, key string, limit int, offset int) (contents []entity.ContentRowResponse)
	GetContentAllMultimediaRow(ctx context.Context, multimediaType string, types string, key string, limit int, offset int) (contents []entity.ContentRowResponse)
	GetContentAllArticleRow(ctx context.Context, limit int, offset int) (contents []entity.ContentRowResponse)
	GetContentAllEditorChoiceRow(ctx context.Context, limit int, offset int) (contents []entity.ContentRowResponse)
	GetContentAllIndeksRow(ctx context.Context, types string, key string, date string, limit int, offset int) (contents []entity.ContentRowResponse)
	GetContentAllSearchRow(ctx context.Context, types string, key string, limit int, offset int) (contents []entity.ContentRowResponse)
}
