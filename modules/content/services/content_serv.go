package services

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/channel"
	"github.com/zakariawahyu/go-echo-news/modules/content"
	"github.com/zakariawahyu/go-echo-news/modules/recommended_content"
	"github.com/zakariawahyu/go-echo-news/modules/region"
	"github.com/zakariawahyu/go-echo-news/modules/sub_channel"
	"github.com/zakariawahyu/go-echo-news/pkg/helpers"
	"github.com/zakariawahyu/go-echo-news/pkg/logger"
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
	zapLogger              logger.Logger
	contextTimeout         time.Duration
}

func NewContentServices(contentRepo content.ContentRepository, redisRepo content.ContentRedisRepository, recommendedContentRepo recommended_content.RecommendedContentRepository, channelRepo channel.ChannelRepository, subChannelRepo sub_channel.SubChannelRepository, regionRepo region.RegionRepository, zapLogger logger.Logger, timeout time.Duration) content.ContentServices {
	return &contentServices{
		contentRepo:            contentRepo,
		redisRepo:              redisRepo,
		recommendedContentRepo: recommendedContentRepo,
		channelRepo:            channelRepo,
		subChannelRepo:         subChannelRepo,
		regionRepo:             regionRepo,
		zapLogger:              zapLogger,
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
	if err != nil {
		serv.zapLogger.Errorf("contentServ.GetContentBySlugOrId.contentRepo.GetBySlugOrId, err = %s", err)
		panic(err)
	}

	if !content.IsEmpty() {
		recommended, err := serv.recommendedContentRepo.GetByContentID(c, content.ID)
		if err != nil {
			serv.zapLogger.Errorf("contentServ.GetContentBySlugOrId.recommendedContentRepo.GetByContentID, err = %s", err)
			panic(err)
		}

		tagName := content.TagNameArray()

		if recommended[0].IsEmpty() {

		}

		content.Content = helpers.AutoLinkedTags(tagName, content.Content, content.TypeID)
	}

	err = serv.redisRepo.SetContent(c, helpers.KeyRedis("read", slug), helpers.Faster, content)
	if err != nil {
		serv.zapLogger.Errorf("contentServ.GetContentBySlugOrId.redisRepo.SetContent, err = %s", err)
		panic(err)
	}

	return entity.NewContentResponse(content)
}

func (serv *contentServices) GetContentAllRow(ctx context.Context, types string, key string, limit int, offset int) (contents []entity.ContentRowResponse) {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	if types != "" {
		if types == "channel" {
			channel, err := serv.channelRepo.GetBySlugOrId(c, key)
			if err != nil {
				serv.zapLogger.Errorf("contentServ.GetContentAllRow.channelRepo.GetBySlugOrId, err = %s", err)
				panic(err)
			}

			key = strconv.FormatInt(channel.ID, 10)
		} else if types == "region" {
			region, err := serv.regionRepo.GetBySlugOrId(c, key)
			if err != nil {
				serv.zapLogger.Errorf("contentServ.GetContentAllRow.regionRepo.GetBySlugOrId, err = %s", err)
				panic(err)
			}

			key = strconv.FormatInt(region.ID, 10)
		} else if types == "subchannel" {
			subChannel, err := serv.subChannelRepo.GetBySlugOrId(c, key)
			if err != nil {
				serv.zapLogger.Errorf("contentServ.GetContentAllRow.subChannelRepo.GetBySlugOrId, err = %s", err)
				panic(err)
			}

			key = strconv.FormatInt(subChannel.ID, 10)
		} else if types != "channel" || types != "subchannel" || types != "region" {
			serv.zapLogger.Errorf("contentServ.GetContentAllRow.NotFound, err = %s", helpers.ErrNotFound)
			panic(helpers.ErrNotFound)
		}
	}

	res, err := serv.contentRepo.GetAllRow(c, types, key, limit, offset)
	if err != nil {
		serv.zapLogger.Errorf("contentServ.GetContentAllRow.contentRepo.GetAllRow, err = %s", err)
		panic(err)
	}

	for _, content := range res {
		contents = append(contents, entity.NewContentRowResponse(content))
	}

	return contents
}

func (serv *contentServices) GetContentAllRowAds(ctx context.Context, types string, key string, limit int, offset int) (contents []entity.ContentRowResponse) {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	if types != "" {
		if types == "channel" {
			channel, err := serv.channelRepo.GetBySlugOrId(c, key)
			if err != nil {
				serv.zapLogger.Errorf("contentServ.GetContentAllRowAds.channelRepo.GetBySlugOrId, err = %s", err)
				panic(err)
			}

			key = strconv.FormatInt(channel.ID, 10)
		} else if types == "region" {
			region, err := serv.regionRepo.GetBySlugOrId(c, key)
			if err != nil {
				serv.zapLogger.Errorf("contentServ.GetContentAllRowAds.regionRepo.GetBySlugOrId, err = %s", err)
				panic(err)
			}

			key = strconv.FormatInt(region.ID, 10)
		} else if types == "subchannel" {
			subChannel, err := serv.subChannelRepo.GetBySlugOrId(c, key)
			if err != nil {
				serv.zapLogger.Errorf("contentServ.GetContentAllRowAds.subChannelRepo.GetBySlugOrId, err = %s", err)
				panic(err)
			}

			key = strconv.FormatInt(subChannel.ID, 10)
		} else if types != "channel" || types != "subchannel" || types != "region" {
			serv.zapLogger.Errorf("contentServ.GetContentAllRowAds.NotFound, err = %s", helpers.ErrNotFound)
			panic(helpers.ErrNotFound)
		}
	}

	res, err := serv.contentRepo.GetAllRowAds(c, types, key, limit, offset)
	if err != nil {
		serv.zapLogger.Errorf("contentServ.GetContentAllRowAds.contentRepo.GetAllRowAds, err = %s", err)
		panic(err)
	}

	for _, content := range res {
		contents = append(contents, entity.NewContentRowResponse(content))
	}

	return contents
}

func (serv *contentServices) GetContentAllLatest(ctx context.Context, types string, key string, limit int, offset int) (contents []entity.ContentRowResponse) {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	if types != "" {
		if types == "channel" {
			channel, err := serv.channelRepo.GetBySlugOrId(c, key)
			if err != nil {
				serv.zapLogger.Errorf("contentServ.GetContentAllLatest.channelRepo.GetBySlugOrId, err = %s", err)
				panic(err)
			}

			key = strconv.FormatInt(channel.ID, 10)
		} else if types == "region" {
			region, err := serv.regionRepo.GetBySlugOrId(c, key)
			if err != nil {
				serv.zapLogger.Errorf("contentServ.GetContentAllLatest.regionRepo.GetBySlugOrId, err = %s", err)
				panic(err)
			}

			key = strconv.FormatInt(region.ID, 10)
		} else if types == "subchannel" {
			subChannel, err := serv.subChannelRepo.GetBySlugOrId(c, key)
			if err != nil {
				serv.zapLogger.Errorf("contentServ.GetContentAllLatest.subChannelRepo.GetBySlugOrId, err = %s", err)
				panic(err)
			}

			key = strconv.FormatInt(subChannel.ID, 10)
		} else if types != "channel" || types != "subchannel" || types != "region" {
			serv.zapLogger.Errorf("contentServ.GetContentAllLatest.NotFound, err = %s", helpers.ErrNotFound)
			panic(helpers.ErrNotFound)
		}
	}

	res, err := serv.contentRepo.GetAllLatest(c, types, key, limit, offset)
	if err != nil {
		serv.zapLogger.Errorf("contentServ.GetContentAllLatest.contentRepo.GetAllLatest, err = %s", err)
		panic(err)
	}

	for _, content := range res {
		contents = append(contents, entity.NewContentRowResponse(content))
	}

	return contents
}

func (serv *contentServices) GetContentAllLatestMultimedia(ctx context.Context, types string, featured bool, limit int, offset int) (contents []entity.ContentRowResponse) {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	res, err := serv.contentRepo.GetAllLatestMultimedia(c, types, featured, limit, offset)
	if err != nil {
		serv.zapLogger.Errorf("contentServ.GetContentAllLatestMultimedia.contentRepo.GetAllLatestMultimedia, err = %s", err)
		panic(err)
	}

	for _, content := range res {
		contents = append(contents, entity.NewContentRowResponse(content))
	}

	return contents
}
