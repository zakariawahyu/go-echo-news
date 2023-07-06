package content_has

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
)

type ContentHasTagRepository interface {
	GetByTagID(ctx context.Context, id string) (*entity.ContentHasTag, error)
}

type ContentHasTopicRepository interface {
	GetByTopicID(ctx context.Context, id string) (*entity.ContentHasTopic, error)
}

type ContentHasReporterRepository interface {
	GetByReporterID(ctx context.Context, id string) (*entity.ContentHasReporter, error)
}
