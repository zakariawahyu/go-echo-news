package recommended_content

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
)

type RecommendedContentRepository interface {
	GetByContentID(ctx context.Context, contentID int64) ([]entity.RecommendedContent, error)
}
