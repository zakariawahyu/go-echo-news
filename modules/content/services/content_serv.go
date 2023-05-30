package services

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
)

type ContentServices interface {
	GetContent(c context.Context, slug string) entity.ContentResponse
}
