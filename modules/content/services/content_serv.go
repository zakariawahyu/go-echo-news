package services

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
)

type ContentServices interface {
	GetBySlug(c context.Context, slug string) (res entity.Content, err error)
}
