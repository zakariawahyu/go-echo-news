package suplemen

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
)

type SuplemenRepository interface {
	GetAll(ctx context.Context) ([]*entity.SuplemenResponse, error)
	GetBySlugOrId(ctx context.Context, slug string) (*entity.SuplemenResponse, error)
}
