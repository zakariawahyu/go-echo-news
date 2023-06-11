package services

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/content/repository"
	repository2 "github.com/zakariawahyu/go-echo-news/modules/recommended/repository"
	"github.com/zakariawahyu/go-echo-news/pkg/exception"
	"github.com/zakariawahyu/go-echo-news/pkg/helpers"
	"time"
)

type ContentServicesImpl struct {
	contentRepo     repository.ContentRepository
	redisRepo       repository.ContentRedisRepository
	recommendedRepo repository2.RecommendedRepository
	contextTimeout  time.Duration
}

func NewContentServices(contentRepo repository.ContentRepository, redisRepo repository.ContentRedisRepository, recommendedRepo repository2.RecommendedRepository, timeout time.Duration) ContentServices {
	return &ContentServicesImpl{
		contentRepo:     contentRepo,
		redisRepo:       redisRepo,
		recommendedRepo: recommendedRepo,
		contextTimeout:  timeout,
	}
}

func (serv *ContentServicesImpl) GetContent(c context.Context, slug string) entity.ContentResponse {
	ctx, cancel := context.WithTimeout(c, serv.contextTimeout)
	defer cancel()

	newBase, err := serv.redisRepo.GetContent(ctx, "test")

	if newBase != nil {
		return entity.NewContentResponse(*newBase)
	}

	content, err := serv.contentRepo.GetContent(ctx, slug)
	exception.PanicIfNeeded(err)

	if !content.IsEmpty() {
		recommended, err := serv.recommendedRepo.GetByContentID(ctx, content.ID)
		exception.PanicIfNeeded(err)

		tagName := content.TagNameArray()

		if recommended[0].IsEmpty() {

		}

		content.Content = helpers.AutoLinkedTags(tagName, content.Content, content.TypeID)
	}

	err = serv.redisRepo.SetContent(ctx, "test", 10, content)
	exception.PanicIfNeeded(err)

	return entity.NewContentResponse(content)
}
