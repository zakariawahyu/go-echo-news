package tag

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
)

type TagRepository interface {
	GetBySlugOrID(ctx context.Context, slug string) (*entity.Tag, error)
}
