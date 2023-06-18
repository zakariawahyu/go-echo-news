package services

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/content"
	"github.com/zakariawahyu/go-echo-news/modules/recommended_content"
	"github.com/zakariawahyu/go-echo-news/pkg/exception"
	"github.com/zakariawahyu/go-echo-news/pkg/helpers"
	"time"
)

type contentServices struct {
	contentRepo            content.ContentRepository
	redisRepo              content.ContentRedisRepository
	recommendedContentRepo recommended_content.RecommendedContentRepository
	contextTimeout         time.Duration
}

func NewContentServices(contentRepo content.ContentRepository, redisRepo content.ContentRedisRepository, recommendedContentRepo recommended_content.RecommendedContentRepository, timeout time.Duration) content.ContentServices {
	return &contentServices{
		contentRepo:            contentRepo,
		redisRepo:              redisRepo,
		recommendedContentRepo: recommendedContentRepo,
		contextTimeout:         timeout,
	}
}

func (serv *contentServices) GetContentBySlugOrId(ctx context.Context, slug string) entity.ContentResponse {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	newBase, err := serv.redisRepo.GetContent(c, helpers.KeyRedis("read", slug))
	if newBase != nil {
		return entity.NewContentResponse(newBase)
	}

	content, err := serv.contentRepo.GetBySlugOrId(c, slug)
	exception.PanicIfNeeded(err)

	if !content.IsEmpty() {
		recommended, err := serv.recommendedContentRepo.GetByContentID(c, content.ID)
		exception.PanicIfNeeded(err)

		tagName := content.TagNameArray()

		if recommended[0].IsEmpty() {

		}

		content.Content = helpers.AutoLinkedTags(tagName, content.Content, content.TypeID)
	}

	err = serv.redisRepo.SetContent(c, helpers.KeyRedis("read", slug), helpers.Faster, content)
	exception.PanicIfNeeded(err)

	return entity.NewContentResponse(content)
}

func (serv *contentServices) GetContentAllHome(ctx context.Context, limit int, offset int) (contents []entity.ContentRowResponse) {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	res, err := serv.contentRepo.GetAllHome(c, limit, offset)
	exception.PanicIfNeeded(err)

	for _, content := range *res {
		contents = append(contents, content)
	}

	return contents
}
