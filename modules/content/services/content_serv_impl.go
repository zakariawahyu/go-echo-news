package services

import (
	"context"
	"github.com/GRbit/go-pcre"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/content/repository"
	repository2 "github.com/zakariawahyu/go-echo-news/modules/recommended/repository"
	"github.com/zakariawahyu/go-echo-news/pkg/exception"
	"github.com/zakariawahyu/go-echo-news/pkg/helpers"
	"time"
)

type ContentServicesImpl struct {
	contentRepo     repository.ContentRepository
	recommendedRepo repository2.RecommendedRepository
	contextTimeout  time.Duration
}

func NewContentServices(repoContent repository.ContentRepository, repoRecommended repository2.RecommendedRepository, timeout time.Duration) ContentServices {
	return &ContentServicesImpl{
		contentRepo:     repoContent,
		recommendedRepo: repoRecommended,
		contextTimeout:  timeout,
	}
}

func (serv *ContentServicesImpl) GetContent(c context.Context, slug string) entity.ContentResponse {
	content := entity.Content{}
	var err error
	ctx, cancel := context.WithTimeout(c, serv.contextTimeout)
	defer cancel()

	regex := pcre.MustCompile(`^[0-9]+$`, 0)
	numericSlug := regex.MatchString(slug, 0)

	if numericSlug {
		content, err = serv.contentRepo.GetByID(ctx, slug)
	} else {
		content, err = serv.contentRepo.GetBySlug(ctx, slug)
	}

	exception.PanicIfNeeded(err)

	if !content.IsEmpty() {
		recommended, err := serv.recommendedRepo.GetByContentID(ctx, content.ID)
		exception.PanicIfNeeded(err)

		tagName := content.TagNameArray()

		if recommended[0].IsEmpty() {

		}

		content.Content = helpers.AutoLinkedTags(tagName, content.Content, content.TypeID)
	}

	return entity.NewContentResponse(content)
}
