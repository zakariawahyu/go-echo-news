package topic

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
)

type TopicRepository interface {
	GetBySlugOrID(ctx context.Context, slug string) (*entity.Topic, error)
}
