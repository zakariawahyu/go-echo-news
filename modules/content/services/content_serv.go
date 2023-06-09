package services

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/channel"
	"github.com/zakariawahyu/go-echo-news/modules/content"
	"github.com/zakariawahyu/go-echo-news/modules/content_has"
	"github.com/zakariawahyu/go-echo-news/modules/recommended_content"
	"github.com/zakariawahyu/go-echo-news/modules/region"
	"github.com/zakariawahyu/go-echo-news/modules/sub_channel"
	"github.com/zakariawahyu/go-echo-news/modules/suplemen"
	"github.com/zakariawahyu/go-echo-news/modules/tag"
	"github.com/zakariawahyu/go-echo-news/modules/topic"
	"github.com/zakariawahyu/go-echo-news/pkg/helpers"
	"github.com/zakariawahyu/go-echo-news/pkg/logger"
	"strconv"
	"time"
)

type contentServices struct {
	contentRepo            content.ContentRepository
	contentRedisRepo       content.ContentRedisRepository
	recommendedContentRepo recommended_content.RecommendedContentRepository
	channelRepo            channel.ChannelRepository
	subChannelRepo         sub_channel.SubChannelRepository
	regionRepo             region.RegionRepository
	suplemenRepo           suplemen.SuplemenRepository
	tagRepo                tag.TagRepository
	topicRepo              topic.TopicRepository
	contentHasTagRepo      content_has.ContentHasTagRepository
	contentHasTopicRepo    content_has.ContentHasTopicRepository
	zapLogger              logger.Logger
	contextTimeout         time.Duration
}

func NewContentServices(contentRepo content.ContentRepository, contentRedisRepo content.ContentRedisRepository, recommendedContentRepo recommended_content.RecommendedContentRepository, channelRepo channel.ChannelRepository, subChannelRepo sub_channel.SubChannelRepository, regionRepo region.RegionRepository, suplemenRepo suplemen.SuplemenRepository, tagRepo tag.TagRepository, topicRepo topic.TopicRepository, contentHasTagRepo content_has.ContentHasTagRepository, contentHasTopicRepo content_has.ContentHasTopicRepository, zapLogger logger.Logger, timeout time.Duration) content.ContentServices {
	return &contentServices{
		contentRepo:            contentRepo,
		contentRedisRepo:       contentRedisRepo,
		recommendedContentRepo: recommendedContentRepo,
		channelRepo:            channelRepo,
		subChannelRepo:         subChannelRepo,
		regionRepo:             regionRepo,
		suplemenRepo:           suplemenRepo,
		tagRepo:                tagRepo,
		topicRepo:              topicRepo,
		contentHasTagRepo:      contentHasTagRepo,
		contentHasTopicRepo:    contentHasTopicRepo,
		zapLogger:              zapLogger,
		contextTimeout:         timeout,
	}
}

func (serv *contentServices) GetContentBySlugOrId(ctx context.Context, slug string) entity.ContentResponse {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	redisData, err := serv.contentRedisRepo.GetContent(c, helpers.KeyRedis("read", slug))
	if redisData != nil {
		return entity.NewContentResponse(redisData)
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

	if err = serv.contentRedisRepo.SetContent(c, helpers.KeyRedis("read", slug), helpers.Faster, content); err != nil {
		serv.zapLogger.Errorf("contentServ.GetContentBySlugOrId.contentRedisRepo.SetContent, err = %s", err)
		panic(err)
	}

	return entity.NewContentResponse(content)
}

func (serv *contentServices) GetContentAllRow(ctx context.Context, types string, key string, limit int, offset int) (contents []entity.ContentRowResponse) {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	var id string

	redisData, err := serv.contentRedisRepo.GetAllContentRow(c, helpers.KeyRedisTypeKey("news-row", types, key, limit, offset))
	if redisData != nil {
		return entity.NewContentRowArrayResponse(redisData)
	}

	if types != "" {
		if types == "channel" {
			channel, err := serv.channelRepo.GetBySlugOrId(c, key)
			if err != nil {
				serv.zapLogger.Errorf("contentServ.GetContentAllRow.channelRepo.GetBySlugOrId, err = %s", err)
				panic(err)
			}

			id = strconv.FormatInt(channel.ID, 10)
		} else if types == "region" {
			region, err := serv.regionRepo.GetBySlugOrId(c, key)
			if err != nil {
				serv.zapLogger.Errorf("contentServ.GetContentAllRow.regionRepo.GetBySlugOrId, err = %s", err)
				panic(err)
			}

			id = strconv.FormatInt(region.ID, 10)
		} else if types == "subchannel" {
			subChannel, err := serv.subChannelRepo.GetBySlugOrId(c, key)
			if err != nil {
				serv.zapLogger.Errorf("contentServ.GetContentAllRow.subChannelRepo.GetBySlugOrId, err = %s", err)
				panic(err)
			}

			id = strconv.FormatInt(subChannel.ID, 10)
		} else if types != "channel" || types != "subchannel" || types != "region" {
			serv.zapLogger.Errorf("contentServ.GetContentAllRow.NotFound, err = %s", helpers.ErrNotFound)
			panic(helpers.ErrNotFound)
		}
	}

	res, err := serv.contentRepo.GetAllRow(c, types, id, limit, offset)
	if err != nil {
		serv.zapLogger.Errorf("contentServ.GetContentAllRow.contentRepo.GetAllRow, err = %s", err)
		panic(err)
	}

	contents = entity.NewContentRowArrayResponse(res)

	if err = serv.contentRedisRepo.SetALlContentRow(c, helpers.KeyRedisTypeKey("news-row", types, key, limit, offset), helpers.Faster, contents); err != nil {
		serv.zapLogger.Errorf("contentServ.GetContentAllRow.contentRedisRepo.SetALlContentRow, err = %s", err)
		panic(err)
	}

	return contents
}

func (serv *contentServices) GetContentAllRowAds(ctx context.Context, types string, key string, limit int, offset int) (contents []entity.ContentRowResponse) {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	var id string

	redisData, err := serv.contentRedisRepo.GetAllContentRow(c, helpers.KeyRedisTypeKey("news-row-ads", types, key, limit, offset))
	if redisData != nil {
		return entity.NewContentRowArrayResponse(redisData)
	}

	if types != "" {
		if types == "channel" {
			channel, err := serv.channelRepo.GetBySlugOrId(c, key)
			if err != nil {
				serv.zapLogger.Errorf("contentServ.GetContentAllRowAds.channelRepo.GetBySlugOrId, err = %s", err)
				panic(err)
			}

			id = strconv.FormatInt(channel.ID, 10)
		} else if types == "region" {
			region, err := serv.regionRepo.GetBySlugOrId(c, key)
			if err != nil {
				serv.zapLogger.Errorf("contentServ.GetContentAllRowAds.regionRepo.GetBySlugOrId, err = %s", err)
				panic(err)
			}

			id = strconv.FormatInt(region.ID, 10)
		} else if types == "subchannel" {
			subChannel, err := serv.subChannelRepo.GetBySlugOrId(c, key)
			if err != nil {
				serv.zapLogger.Errorf("contentServ.GetContentAllRowAds.subChannelRepo.GetBySlugOrId, err = %s", err)
				panic(err)
			}

			id = strconv.FormatInt(subChannel.ID, 10)
		} else if types != "channel" || types != "subchannel" || types != "region" {
			serv.zapLogger.Errorf("contentServ.GetContentAllRowAds.NotFound, err = %s", helpers.ErrNotFound)
			panic(helpers.ErrNotFound)
		}
	}

	res, err := serv.contentRepo.GetAllRowAds(c, types, id, limit, offset)
	if err != nil {
		serv.zapLogger.Errorf("contentServ.GetContentAllRowAds.contentRepo.GetAllRowAds, err = %s", err)
		panic(err)
	}

	contents = entity.NewContentRowArrayResponse(res)

	if err = serv.contentRedisRepo.SetALlContentRow(c, helpers.KeyRedisTypeKey("news-row-ads", types, key, limit, offset), helpers.Faster, contents); err != nil {
		serv.zapLogger.Errorf("contentServ.GetContentAllRowAds.contentRedisRepo.SetALlContentRow, err = %s", err)
		panic(err)
	}

	return contents
}

func (serv *contentServices) GetContentAllLatest(ctx context.Context, types string, key string, limit int, offset int) (contents []entity.ContentRowResponse) {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	var id string

	redisData, err := serv.contentRedisRepo.GetAllContentRow(c, helpers.KeyRedisTypeKey("latest", types, key, limit, offset))
	if redisData != nil {
		return entity.NewContentRowArrayResponse(redisData)
	}

	if types != "" {
		if types == "channel" {
			channel, err := serv.channelRepo.GetBySlugOrId(c, key)
			if err != nil {
				serv.zapLogger.Errorf("contentServ.GetContentAllLatest.channelRepo.GetBySlugOrId, err = %s", err)
				panic(err)
			}

			id = strconv.FormatInt(channel.ID, 10)
		} else if types == "region" {
			region, err := serv.regionRepo.GetBySlugOrId(c, key)
			if err != nil {
				serv.zapLogger.Errorf("contentServ.GetContentAllLatest.regionRepo.GetBySlugOrId, err = %s", err)
				panic(err)
			}

			id = strconv.FormatInt(region.ID, 10)
		} else if types == "subchannel" {
			subChannel, err := serv.subChannelRepo.GetBySlugOrId(c, key)
			if err != nil {
				serv.zapLogger.Errorf("contentServ.GetContentAllLatest.subChannelRepo.GetBySlugOrId, err = %s", err)
				panic(err)
			}

			id = strconv.FormatInt(subChannel.ID, 10)
		} else if types != "channel" || types != "subchannel" || types != "region" {
			serv.zapLogger.Errorf("contentServ.GetContentAllLatest.NotFound, err = %s", helpers.ErrNotFound)
			panic(helpers.ErrNotFound)
		}
	}

	res, err := serv.contentRepo.GetAllLatest(c, types, id, limit, offset)
	if err != nil {
		serv.zapLogger.Errorf("contentServ.GetContentAllLatest.contentRepo.GetAllLatest, err = %s", err)
		panic(err)
	}

	contents = entity.NewContentRowArrayResponse(res)

	if err = serv.contentRedisRepo.SetALlContentRow(c, helpers.KeyRedisTypeKey("latest", types, key, limit, offset), helpers.Faster, contents); err != nil {
		serv.zapLogger.Errorf("contentServ.GetContentAllLatest.contentRedisRepo.SetALlContentRow, err = %s", err)
		panic(err)
	}

	return contents
}

func (serv *contentServices) GetContentAllLatestMultimedia(ctx context.Context, types string, featured bool, limit int, offset int) (contents []entity.ContentRowResponse) {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	redisData, err := serv.contentRedisRepo.GetAllContentRow(c, helpers.KeyRedisTypeFeatured("latest", types, featured, limit, offset))
	if redisData != nil {
		return entity.NewContentRowArrayResponse(redisData)
	}

	res, err := serv.contentRepo.GetAllLatestMultimedia(c, types, featured, limit, offset)
	if err != nil {
		serv.zapLogger.Errorf("contentServ.GetContentAllLatestMultimedia.contentRepo.GetAllLatestMultimedia, err = %s", err)
		panic(err)
	}

	contents = entity.NewContentRowArrayResponse(res)

	if err = serv.contentRedisRepo.SetALlContentRow(c, helpers.KeyRedisTypeFeatured("latest", types, featured, limit, offset), helpers.Faster, contents); err != nil {
		serv.zapLogger.Errorf("contentServ.GetContentAllLatestMultimedia.contentRedisRepo.SetALlContentRow, err = %s", err)
		panic(err)
	}

	return contents
}

func (serv *contentServices) GetContentAllHeadline(ctx context.Context, types string, key string, limit int, offset int) (contents []entity.ContentRowResponse) {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	var id string

	redisData, err := serv.contentRedisRepo.GetAllContentRow(c, helpers.KeyRedisTypeKey("headline", types, key, limit, offset))
	if redisData != nil {
		return entity.NewContentRowArrayResponse(redisData)
	}

	if types != "" {
		if types == "channel" {
			channel, err := serv.channelRepo.GetBySlugOrId(c, key)
			if err != nil {
				serv.zapLogger.Errorf("contentServ.GetContentAllHeadline.channelRepo.GetBySlugOrId, err = %s", err)
				panic(err)
			}

			id = strconv.FormatInt(channel.ID, 10)
		} else if types == "region" {
			region, err := serv.regionRepo.GetBySlugOrId(c, key)
			if err != nil {
				serv.zapLogger.Errorf("contentServ.GetContentAllHeadline.regionRepo.GetBySlugOrId, err = %s", err)
				panic(err)
			}

			id = strconv.FormatInt(region.ID, 10)
		} else if types == "subchannel" {
			subChannel, err := serv.subChannelRepo.GetBySlugOrId(c, key)
			if err != nil {
				serv.zapLogger.Errorf("contentServ.GetContentAllHeadline.subChannelRepo.GetBySlugOrId, err = %s", err)
				panic(err)
			}

			id = strconv.FormatInt(subChannel.ID, 10)
		} else if types != "channel" || types != "subchannel" || types != "region" {
			serv.zapLogger.Errorf("contentServ.GetContentAllHeadline.NotFound, err = %s", helpers.ErrNotFound)
			panic(helpers.ErrNotFound)
		}
	}

	res, err := serv.contentRepo.GetAllHeadline(c, types, id, limit, offset)
	if err != nil {
		serv.zapLogger.Errorf("contentServ.GetContentAllHeadline.contentRepo.GetAllHeadline, err = %s", err)
		panic(err)
	}

	contents = entity.NewContentRowArrayResponse(res)

	if err = serv.contentRedisRepo.SetALlContentRow(c, helpers.KeyRedisTypeKey("headline", types, key, limit, offset), helpers.Faster, contents); err != nil {
		serv.zapLogger.Errorf("contentServ.GetContentAllHeadline.contentRedisRepo.SetALlContentRow, err = %s", err)
		panic(err)
	}

	return contents
}

func (serv *contentServices) GetContentAllHeadlineAds(ctx context.Context, types string, key string, limit int, offset int) (contents []entity.ContentRowResponse) {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	var id string

	redisData, err := serv.contentRedisRepo.GetAllContentRow(c, helpers.KeyRedisTypeKey("headline-ads", types, key, limit, offset))
	if redisData != nil {
		return entity.NewContentRowArrayResponse(redisData)
	}

	if types != "" {
		if types == "channel" {
			channel, err := serv.channelRepo.GetBySlugOrId(c, key)
			if err != nil {
				serv.zapLogger.Errorf("contentServ.GetContentAllHeadlineAds.channelRepo.GetBySlugOrId, err = %s", err)
				panic(err)
			}

			id = strconv.FormatInt(channel.ID, 10)
		} else if types == "region" {
			region, err := serv.regionRepo.GetBySlugOrId(c, key)
			if err != nil {
				serv.zapLogger.Errorf("contentServ.GetContentAllHeadlineAds.regionRepo.GetBySlugOrId, err = %s", err)
				panic(err)
			}

			id = strconv.FormatInt(region.ID, 10)
		} else if types == "subchannel" {
			subChannel, err := serv.subChannelRepo.GetBySlugOrId(c, key)
			if err != nil {
				serv.zapLogger.Errorf("contentServ.GetContentAllHeadlineAds.subChannelRepo.GetBySlugOrId, err = %s", err)
				panic(err)
			}

			id = strconv.FormatInt(subChannel.ID, 10)
		} else if types != "channel" || types != "subchannel" || types != "region" {
			serv.zapLogger.Errorf("contentServ.GetContentAllHeadlineAds.NotFound, err = %s", helpers.ErrNotFound)
			panic(helpers.ErrNotFound)
		}
	}

	res, err := serv.contentRepo.GetAllHeadlineAds(c, types, id, limit, offset)
	if err != nil {
		serv.zapLogger.Errorf("contentServ.GetContentAllHeadlineAds.contentRepo.GetAllHeadlineAds, err = %s", err)
		panic(err)
	}

	contents = entity.NewContentRowArrayResponse(res)

	if err = serv.contentRedisRepo.SetALlContentRow(c, helpers.KeyRedisTypeKey("headline-ads", types, key, limit, offset), helpers.Faster, contents); err != nil {
		serv.zapLogger.Errorf("contentServ.GetContentAllHeadlineAds.contentRedisRepo.SetALlContentRow, err = %s", err)
		panic(err)
	}

	return contents
}

func (serv *contentServices) GetContentAllMultimediaRow(ctx context.Context, multimediaType string, types string, key string, limit int, offset int) (contents []entity.ContentRowResponse) {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	var id string

	redisData, err := serv.contentRedisRepo.GetAllContentRow(c, helpers.KeyRedisMultimediaTypeKey(multimediaType, "row", types, key, limit, offset))
	if redisData != nil {
		return entity.NewContentRowArrayResponse(redisData)
	}

	if types != "" {
		if types == "channel" {
			channel, err := serv.channelRepo.GetBySlugOrId(c, key)
			if err != nil {
				serv.zapLogger.Errorf("contentServ.GetContentAllMultimediaRow.channelRepo.GetBySlugOrId, err = %s", err)
				panic(err)
			}

			id = strconv.FormatInt(channel.ID, 10)
		} else if types == "region" {
			region, err := serv.regionRepo.GetBySlugOrId(c, key)
			if err != nil {
				serv.zapLogger.Errorf("contentServ.GetContentAllMultimediaRow.regionRepo.GetBySlugOrId, err = %s", err)
				panic(err)
			}

			id = strconv.FormatInt(region.ID, 10)
		} else if types == "subchannel" {
			subChannel, err := serv.subChannelRepo.GetBySlugOrId(c, key)
			if err != nil {
				serv.zapLogger.Errorf("contentServ.GetContentAllMultimediaRow.subChannelRepo.GetBySlugOrId, err = %s", err)
				panic(err)
			}

			id = strconv.FormatInt(subChannel.ID, 10)
		} else if types != "channel" || types != "subchannel" || types != "region" {
			serv.zapLogger.Errorf("contentServ.GetContentAllMultimediaRow.NotFound, err = %s", helpers.ErrNotFound)
			panic(helpers.ErrNotFound)
		}
	}

	res, err := serv.contentRepo.GetAllMultimediaRow(c, multimediaType, types, id, limit, offset)
	if err != nil {
		serv.zapLogger.Errorf("contentServ.GetContentAllMultimediaRow.contentRepo.GetAllMultimediaRow, err = %s", err)
		panic(err)
	}

	contents = entity.NewContentRowArrayResponse(res)

	if err = serv.contentRedisRepo.SetALlContentRow(c, helpers.KeyRedisMultimediaTypeKey(multimediaType, "row", types, key, limit, offset), helpers.Faster, contents); err != nil {
		serv.zapLogger.Errorf("contentServ.GetContentAllMultimediaRow.contentRedisRepo.SetALlContentRow, err = %s", err)
		panic(err)
	}

	return contents
}

func (serv *contentServices) GetContentAllArticleRow(ctx context.Context, limit int, offset int) (contents []entity.ContentRowResponse) {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	redisData, err := serv.contentRedisRepo.GetAllContentRow(c, helpers.KeyRedisTypeKey("article-row", "", "", limit, offset))
	if redisData != nil {
		return entity.NewContentRowArrayResponse(redisData)
	}

	res, err := serv.contentRepo.GetAllArticleRow(c, limit, offset)
	if err != nil {
		serv.zapLogger.Errorf("contentServ.GetContentAllArticleRow.contentRepo.GetAllArticleRow, err = %s", err)
		panic(err)
	}

	contents = entity.NewContentRowArrayResponse(res)

	if err = serv.contentRedisRepo.SetALlContentRow(c, helpers.KeyRedisTypeKey("article-row", "", "", limit, offset), helpers.Faster, contents); err != nil {
		serv.zapLogger.Errorf("contentServ.GetContentAllArticleRow.contentRedisRepo.SetALlContentRow, err = %s", err)
		panic(err)
	}

	return contents
}

func (serv *contentServices) GetContentAllEditorChoiceRow(ctx context.Context, limit int, offset int) (contents []entity.ContentRowResponse) {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	redisData, err := serv.contentRedisRepo.GetAllContentRow(c, helpers.KeyRedisTypeKey("editor-choice", "", "", limit, offset))
	if redisData != nil {
		return entity.NewContentRowArrayResponse(redisData)
	}

	res, err := serv.contentRepo.GetAllEditorChoiceRow(c, limit, offset)
	if err != nil {
		serv.zapLogger.Errorf("contentServ.GetContentAllEditorChoiceRow.contentRepo.GetAllEditorChoiceRow, err = %s", err)
		panic(err)
	}

	contents = entity.NewContentRowArrayResponse(res)

	if err = serv.contentRedisRepo.SetALlContentRow(c, helpers.KeyRedisTypeKey("editor-choice", "", "", limit, offset), helpers.Faster, contents); err != nil {
		serv.zapLogger.Errorf("contentServ.GetContentAllEditorChoiceRow.contentRedisRepo.SetALlContentRow, err = %s", err)
		panic(err)
	}

	return contents
}

func (serv *contentServices) GetContentAllIndeksRow(ctx context.Context, types string, key string, date string, limit int, offset int) (contents []entity.ContentRowResponse) {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	var id string

	redisData, err := serv.contentRedisRepo.GetAllContentRow(c, helpers.KeyRedisTypeKeyDate("indeks", types, key, date, limit, offset))
	if redisData != nil {
		return entity.NewContentRowArrayResponse(redisData)
	}

	if types == "channel" {
		if key == "video" || key == "photo" || key == "infografis" || key == "multimedia" {
			id = key
		} else {
			channel, err := serv.channelRepo.GetBySlugOrId(c, key)
			if err != nil {
				serv.zapLogger.Errorf("contentServ.GetContentAllIndeksRow.channelRepo.GetBySlugOrId, err = %s", err)
				panic(err)
			}

			id = strconv.FormatInt(channel.ID, 10)
		}
	} else if types == "region" {
		region, err := serv.regionRepo.GetBySlugOrId(c, key)
		if err != nil {
			serv.zapLogger.Errorf("contentServ.GetContentAllIndeksRow.regionRepo.GetBySlugOrId, err = %s", err)
			panic(err)
		}

		id = strconv.FormatInt(region.ID, 10)
	} else if types == "suplemen" {
		suplemen, err := serv.suplemenRepo.GetBySlugOrId(c, key)
		if err != nil {
			serv.zapLogger.Errorf("contentServ.GetContentAllIndeksRow.regionRepo.GetBySlugOrId, err = %s", err)
			panic(err)
		}

		id = strconv.FormatInt(suplemen.ID, 10)
	}

	res, err := serv.contentRepo.GetAllIndeks(c, types, id, date, limit, offset)
	if err != nil {
		serv.zapLogger.Errorf("contentServ.GetContentAllIndeksRow.contentRepo.GetAllIndeks, err = %s", err)
		panic(err)
	}

	contents = entity.NewContentRowArrayResponse(res)

	if err = serv.contentRedisRepo.SetALlContentRow(c, helpers.KeyRedisTypeKeyDate("indeks", types, key, date, limit, offset), helpers.Faster, contents); err != nil {
		serv.zapLogger.Errorf("contentServ.GetContentAllIndeksRow.contentRedisRepo.SetALlContentRow, err = %s", err)
		panic(err)
	}

	return contents
}

func (serv *contentServices) GetContentAllSearchRow(ctx context.Context, types string, key string, limit int, offset int) (contents []entity.ContentRowResponse) {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	var data interface{}
	var contentIds []int64

	redisData, err := serv.contentRedisRepo.GetAllContentRow(c, helpers.KeyRedisTypeKey("search", types, key, limit, offset))
	if redisData != nil {
		return entity.NewContentRowArrayResponse(redisData)
	}

	if types == "tag" || types == "tag-headline" {
		tag, err := serv.tagRepo.GetBySlugOrID(c, key)
		if err != nil {
			serv.zapLogger.Errorf("contentServ.GetContentAllSearchRow.tagRepo.GetBySlugOrID, err = %s", err)
			panic(err)
		}

		contentId, err := serv.contentHasTagRepo.GetByTagIDLimited(c, tag.ID, limit)
		for _, value := range contentId {
			contentIds = append(contentIds, value.ContentID)
		}
		if err != nil {
			serv.zapLogger.Errorf("contentServ.GetContentAllSearchRow.contentHasTagRepo.GetByTagID, err = %s", err)
			panic(err)
		}

		data = contentIds
	} else if types == "topic" || types == "topic-headline" || types == "topic-headline-subkanal" {
		topic, err := serv.topicRepo.GetBySlugOrID(c, key)
		if err != nil {
			serv.zapLogger.Errorf("contentServ.GetContentAllSearchRow.topicRepo.GetBySlugOrID, err = %s", err)
			panic(err)
		}

		contentId, err := serv.contentHasTopicRepo.GetByTopicIDLimited(c, topic.ID, limit)
		for _, value := range contentId {
			contentIds = append(contentIds, value.ContentID)
		}
		if err != nil {
			serv.zapLogger.Errorf("contentServ.GetContentAllSearchRow.contentHasTopicRepo.GetByTopicID, err = %s", err)
			panic(err)
		}
		data = contentIds
	} else {
		data = key
	}

	res, err := serv.contentRepo.GetAllSearch(c, types, data, limit, offset)
	if err != nil {
		serv.zapLogger.Errorf("contentServ.GetContentAllSearchRow.contentRepo.GetAllSearch, err = %s", err)
		panic(err)
	}

	contents = entity.NewContentRowArrayResponse(res)

	if err = serv.contentRedisRepo.SetALlContentRow(c, helpers.KeyRedisTypeKey("search", types, key, limit, offset), helpers.Faster, contents); err != nil {
		serv.zapLogger.Errorf("contentServ.GetContentAllSearchRow.contentRedisRepo.SetALlContentRow, err = %s", err)
		panic(err)
	}

	return contents
}
