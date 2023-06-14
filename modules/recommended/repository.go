package recommended

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
)

type RecommendedRepository interface {
	GetByContentID(ctx context.Context, contentID int64) ([]entity.RecommendedContent, error)
}
