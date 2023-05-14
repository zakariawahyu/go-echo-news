package repository

import (
	"github.com/zakariawahyu/go-echo-news/entity"
	"golang.org/x/net/context"
)

type ContentRepository interface {
	GetBySlug(ctx context.Context, slug string) (res entity.Content, err error)
}
