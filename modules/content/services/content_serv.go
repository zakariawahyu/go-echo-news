package services

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/channel"
	"github.com/zakariawahyu/go-echo-news/modules/content"
	"github.com/zakariawahyu/go-echo-news/modules/recommended_content"
	"github.com/zakariawahyu/go-echo-news/modules/region"
	"github.com/zakariawahyu/go-echo-news/modules/sub_channel"
	"github.com/zakariawahyu/go-echo-news/pkg/exception"
	"github.com/zakariawahyu/go-echo-news/pkg/helpers"
	"strconv"
	"time"
)

type contentServices struct {
	contentRepo            content.ContentRepository
	redisRepo              content.ContentRedisRepository
	recommendedContentRepo recommended_content.RecommendedContentRepository
	channelRepo            channel.ChannelRepository
	subChannelRepo         sub_channel.SubChannelRepository
	regionRepo             region.RegionRepository
	contextTimeout         time.Duration
}

func NewContentServices(contentRepo content.ContentRepository, redisRepo content.ContentRedisRepository, recommendedContentRepo recommended_content.RecommendedContentRepository, channelRepo channel.ChannelRepository, subChannelRepo sub_channel.SubChannelRepository, regionRepo region.RegionRepository, timeout time.Duration) content.ContentServices {
	return &contentServices{
		contentRepo:            contentRepo,
		redisRepo:              redisRepo,
		recommendedContentRepo: recommendedContentRepo,
		channelRepo:            channelRepo,
		subChannelRepo:         subChannelRepo,
		regionRepo:             regionRepo,
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
		contents = append(contents, entity.NewContentRowResponse(&content))
	}

	return contents
}

func (serv *contentServices) GetContentAllChannel(ctx context.Context, key string, limit int, offset int) (contents []entity.ContentRowResponse) {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	channel, err := serv.channelRepo.GetBySlugOrId(c, key)
	exception.PanicIfNeeded(err)

	res, err := serv.contentRepo.GetAllChannel(c, channel.ID, limit, offset)
	exception.PanicIfNeeded(err)

	for _, content := range *res {
		contents = append(contents, entity.NewContentRowResponse(&content))
	}

	return contents
}

func (serv *contentServices) GetContentAllSubChannel(ctx context.Context, key string, limit int, offset int) (contents []entity.ContentRowResponse) {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	subChannel, err := serv.subChannelRepo.GetBySlugOrId(c, key)
	exception.PanicIfNeeded(err)

	res, err := serv.contentRepo.GetAllSubChannel(c, subChannel.ID, limit, offset)
	exception.PanicIfNeeded(err)

	for _, content := range *res {
		contents = append(contents, entity.NewContentRowResponse(&content))
	}

	return contents
}

func (serv *contentServices) GetContentAllRegion(ctx context.Context, key string, limit int, offset int) (contents []entity.ContentRowResponse) {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	region, err := serv.regionRepo.GetBySlugOrId(c, key)
	exception.PanicIfNeeded(err)

	res, err := serv.contentRepo.GetAllRegion(c, region.ID, limit, offset)
	exception.PanicIfNeeded(err)

	for _, content := range *res {
		contents = append(contents, entity.NewContentRowResponse(&content))
	}

	return contents
}

func (serv *contentServices) GetContentAllAds(ctx context.Context, types string, key string, limit int, offset int) (contents []entity.ContentRowResponse) {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	if types != "" {
		if types == "channel" {
			channel, err := serv.channelRepo.GetBySlugOrId(c, key)
			exception.PanicIfNeeded(err)

			key = strconv.FormatInt(channel.ID, 10)
		} else if types == "region" {
			region, err := serv.regionRepo.GetBySlugOrId(c, key)
			exception.PanicIfNeeded(err)

			key = strconv.FormatInt(region.ID, 10)
		} else if types == "subchannel" {
			subChannel, err := serv.subChannelRepo.GetBySlugOrId(c, key)
			exception.PanicIfNeeded(err)

			key = strconv.FormatInt(subChannel.ID, 10)
		}
	}

	res, err := serv.contentRepo.GetAllAds(c, types, key, limit, offset)
	exception.PanicIfNeeded(err)

	for _, content := range *res {
		contents = append(contents, entity.NewContentRowResponse(&content))
	}

	return contents
}
