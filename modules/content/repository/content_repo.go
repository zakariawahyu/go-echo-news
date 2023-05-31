package repository

import (
	"github.com/zakariawahyu/go-echo-news/entity"
	"golang.org/x/net/context"
)

type ContentRepository interface {
	GetContent(ctx context.Context, slug string) (entity.Content, error)
}
