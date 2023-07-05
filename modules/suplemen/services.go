package suplemen

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
)

type SuplemenServices interface {
	GetAllSuplemen(ctx context.Context) (suplemens []entity.SuplemenResponse)
	GetSuplemenBySlugOrId(ctx context.Context, slug string) entity.SuplemenResponse
}
