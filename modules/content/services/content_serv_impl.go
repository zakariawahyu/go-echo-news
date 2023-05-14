package services

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/content/repository"
	"time"
)

type ContentServicesImpl struct {
	contentRepo    repository.ContentRepository
	contextTimeout time.Duration
}

func NewContentServices(repo repository.ContentRepository, timeout time.Duration) ContentServices {
	return &ContentServicesImpl{
		contentRepo:    repo,
		contextTimeout: timeout,
	}
}

func (serv *ContentServicesImpl) GetBySlug(c context.Context, slug string) (res entity.Content, err error) {
	ctx, cancel := context.WithTimeout(c, serv.contextTimeout)
	defer cancel()

	res, err = serv.contentRepo.GetBySlug(ctx, slug)
	if err != nil {
		return entity.Content{}, err
	}

	return
}
