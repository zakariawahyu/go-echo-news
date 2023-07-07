package content_has

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
)

type ContentHasTagRepository interface {
	GetByTagIDLimited(ctx context.Context, id int64, limit int) ([]*entity.ContentHasTag, error)
}

type ContentHasTopicRepository interface {
	GetByTopicIDLimited(ctx context.Context, id int64, limit int) ([]*entity.ContentHasTopic, error)
}

type ContentHasReporterRepository interface {
	GetByReporterIDLimited(ctx context.Context, id int64, limit int) ([]*entity.ContentHasReporter, error)
}
