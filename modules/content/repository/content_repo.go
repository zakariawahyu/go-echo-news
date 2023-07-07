package repository

import (
	"context"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/content"
	"strings"
	"time"
)

var (
	currentTime = time.Now()
)

type contentRepository struct {
	DB *bun.DB
}

func NewContentRepository(DB *bun.DB) content.ContentRepository {
	return &contentRepository{DB}
}

func (repo *contentRepository) GetBySlugOrId(ctx context.Context, slug string) (*entity.Content, error) {
	content := &entity.Content{}

	if err := repo.DB.NewSelect().Model(content).
		Relation("User").
		Relation("Region").
		Relation("Channel").
		Relation("SubChannel").
		Relation("Tags").
		Relation("Topics").
		Relation("Reporters").
		Relation("SubPhotos").
		Where("content.slug = ?", slug).
		WhereOr("content.id = ?", slug).
		Where("content.is_active", true).
		Scan(ctx); err != nil {
		return nil, err
	}

	return content, nil
}

func (repo *contentRepository) GetAllRow(ctx context.Context, types string, key string, limit int, offset int) ([]*entity.ContentRowResponse, error) {
	content := []*entity.ContentRowResponse{}

	if err := repo.DB.NewSelect().Model(&content).
		Where("content_row_response.is_active = ?", true).
		Apply(func(q *bun.SelectQuery) *bun.SelectQuery { //Relation Function
			if types == "channel" || types == "subchannel" {
				q = q.Relation("Channel").Relation("SubChannel")
			} else if types == "region" {
				q = q.Relation("Region").Relation("SubPhotos")
			} else { //home
				q = q.Relation("Channel").Relation("SubChannel").Relation("Region").Relation("SubPhotos")
			}
			return q
		}).
		Apply(func(q *bun.SelectQuery) *bun.SelectQuery { //Where Function
			if types == "channel" {
				q = q.Where("headline_type NOT IN (?)", bun.In([]int64{1, 2})).Where("type = ?", "channel").Where("type_id = ?", key)
			} else if types == "subchannel" {
				q = q.Where("headline_type = 0").Where("type = ?", "channel").Where("type_child_id = ?", key)
			} else if types == "region" {
				q = q.Where("headline_type NOT IN (?)", bun.In([]int{1, 2}))
				if key == "1189" {
					q = q.Where("type = ?", "region")
				} else {
					q = q.WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
						return q.Where("type = ?", "region").Where("type_id = ?", key)
					}).WhereGroup(" OR ", func(q *bun.SelectQuery) *bun.SelectQuery {
						return q.Where("type = ?", "photo").Where("type_id = ?", key).Where("type_child_id is null")
					}).WhereGroup(" OR ", func(q *bun.SelectQuery) *bun.SelectQuery {
						return q.Where("type = ?", "video").Where("type_id = ?", key).Where("type_child_id is null")
					})
				}
			} else { //home
				q = q.Where("headline_type != 1").Where("is_national = 1")
			}
			return q
		}).
		WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Where("ads_position is null").WhereOr("ads_position = 0")
		}).
		Order("published_date desc").
		Limit(limit).
		Offset(offset).
		Scan(ctx); err != nil {
		return nil, err
	}

	return content, nil
}

func (repo *contentRepository) GetAllRowAds(ctx context.Context, types string, key string, limit int, offset int) ([]*entity.ContentRowResponse, error) {
	content := []*entity.ContentRowResponse{}

	if err := repo.DB.NewSelect().Model(&content).
		Relation("Region").
		Relation("Channel").
		Relation("SubChannel").
		Relation("SubPhotos").
		Where("content_row_response.is_active = ?", true).
		Where("content_row_response.headline_type != 1").
		Where("ads_expired_date >= ?", currentTime.Format("2006-01-02 15:01:05")).
		Where("ads_position between ? and ?", 1, 10).
		Order("ads_position asc").
		Apply(func(q *bun.SelectQuery) *bun.SelectQuery {
			if types != "" {
				if types == "channel" {
					q = q.WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
						return q.Where("type_id = ?", key).Where("type_child_id is not null")
					})
				} else if types == "region" {
					q = q.WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
						return q.Where("type_id = ?", key).Where("type_child_id is null")
					})
				} else if types == "subchannel" {
					q = q.WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
						return q.Where("type_child_id = ?", key)
					})
				}
			}
			return q
		}).
		Order("published_date desc").
		Limit(limit).
		Offset(offset).
		Scan(ctx); err != nil {
		return nil, err
	}

	return content, nil
}

func (repo *contentRepository) GetAllLatest(ctx context.Context, types string, key string, limit int, offset int) ([]*entity.ContentRowResponse, error) {
	content := []*entity.ContentRowResponse{}

	if err := repo.DB.NewSelect().Model(&content).
		Where("content_row_response.is_active = ?", true).
		Apply(func(q *bun.SelectQuery) *bun.SelectQuery { //Relation Function
			if types == "channel" || types == "subchannel" {
				q = q.Relation("Channel").Relation("SubChannel")
			} else if types == "region" {
				q = q.Relation("Region").Relation("SubPhotos")
			} else { //home
				q = q.Relation("Channel").Relation("SubChannel").Relation("Region").Relation("SubPhotos")
			}
			return q
		}).
		Where("content_row_response.headline_type != 1").
		Apply(func(q *bun.SelectQuery) *bun.SelectQuery { //Where Function
			if types == "channel" {
				q = q.Where("type = ?", "channel").Where("type_id = ?", key)
			} else if types == "subchannel" {
				q = q.Where("type = ?", "channel").Where("type_child_id = ?", key)
			} else if types == "region" {
				if key == "1189" {
					q = q.Where("type = ?", "region")
				} else {
					q = q.WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
						return q.Where("type = ?", "region").Where("type_id = ?", key)
					}).WhereGroup(" OR ", func(q *bun.SelectQuery) *bun.SelectQuery {
						return q.Where("type = ?", "photo").Where("type_id = ?", key).Where("type_child_id is null")
					}).WhereGroup(" OR ", func(q *bun.SelectQuery) *bun.SelectQuery {
						return q.Where("type = ?", "video").Where("type_id = ?", key).Where("type_child_id is null")
					})
				}
			}
			return q
		}).
		Apply(func(q *bun.SelectQuery) *bun.SelectQuery {
			if types != "region" {
				q = q.Where("is_national = 1")
			}
			return q
		}).
		WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Where("ads_position is null").WhereOr("ads_position = 0")
		}).
		Order("published_date desc").
		Limit(limit).
		Offset(offset).
		Scan(ctx); err != nil {
		return nil, err
	}

	return content, nil
}

func (repo *contentRepository) GetAllLatestMultimedia(ctx context.Context, types string, featured bool, limit int, offset int) ([]*entity.ContentRowResponse, error) {
	content := []*entity.ContentRowResponse{}

	if err := repo.DB.NewSelect().Model(&content).
		Relation("Channel").Relation("SubChannel").Relation("Region").Relation("SubPhotos").
		Where("content_row_response.is_active = ?", true).
		Where("content_row_response.type = ?", types).
		Apply(func(q *bun.SelectQuery) *bun.SelectQuery {
			if featured {
				q = q.Where("is_featured = ?", true)
			}
			return q
		}).
		Order("published_date desc").
		Limit(limit).
		Offset(offset).
		Scan(ctx); err != nil {
		return nil, err
	}

	return content, nil
}

func (repo *contentRepository) GetAllHeadline(ctx context.Context, types string, key string, limit int, offset int) ([]*entity.ContentRowResponse, error) {
	content := []*entity.ContentRowResponse{}

	if err := repo.DB.NewSelect().Model(&content).
		Where("content_row_response.is_active = ?", true).
		Apply(func(q *bun.SelectQuery) *bun.SelectQuery { //Relation Function
			if types == "channel" || types == "subchannel" {
				q = q.Relation("Channel").Relation("SubChannel")
			} else if types == "region" {
				q = q.Relation("Region").Relation("SubPhotos")
			} else { //home
				q = q.Relation("Channel").Relation("SubChannel").Relation("Region").Relation("SubPhotos")
			}
			return q
		}).
		Apply(func(q *bun.SelectQuery) *bun.SelectQuery { //Where Function
			if types == "channel" {
				q = q.Where("headline_type IN (?)", bun.In([]int64{1, 2})).Where("type = ?", "channel").Where("type_id = ?", key)
			} else if types == "subchannel" {
				q = q.Where("headline_type != 0").Where("type = ?", "channel").Where("type_child_id = ?", key)
			} else if types == "region" {
				q = q.Where("headline_type IN (?)", bun.In([]int{1, 2}))
				if key == "1189" {
					q = q.Where("type = ?", "region")
				} else {
					q = q.WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
						return q.Where("type = ?", "region").Where("type_id = ?", key)
					}).WhereGroup(" OR ", func(q *bun.SelectQuery) *bun.SelectQuery {
						return q.Where("type = ?", "photo").Where("type_id = ?", key).Where("type_child_id is null")
					}).WhereGroup(" OR ", func(q *bun.SelectQuery) *bun.SelectQuery {
						return q.Where("type = ?", "video").Where("type_id = ?", key).Where("type_child_id is null")
					})
				}
			} else { //home
				q = q.Where("headline_type = 1").Where("is_national = true")
			}
			return q
		}).
		WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Where("ads_position is null").WhereOr("ads_position = 0")
		}).
		Order("published_date desc").
		Limit(limit).
		Offset(offset).
		Scan(ctx); err != nil {
		return nil, err
	}

	return content, nil
}

func (repo *contentRepository) GetAllHeadlineAds(ctx context.Context, types string, key string, limit int, offset int) ([]*entity.ContentRowResponse, error) {
	content := []*entity.ContentRowResponse{}

	if err := repo.DB.NewSelect().Model(&content).
		Relation("Region").
		Relation("Channel").
		Relation("SubChannel").
		Relation("SubPhotos").
		Where("content_row_response.is_active = ?", true).
		Where("content_row_response.headline_type = 1").
		Where("ads_expired_date >= ?", currentTime.Format("2006-01-02 15:01:05")).
		Where("ads_position between ? and ?", 1, 5).
		Order("ads_position asc").
		Apply(func(q *bun.SelectQuery) *bun.SelectQuery {
			if types != "" {
				if types == "channel" {
					q = q.WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
						return q.Where("type_id = ?", key).Where("type_child_id is not null")
					})
				} else if types == "region" {
					q = q.WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
						return q.Where("type_id = ?", key).Where("type_child_id is null")
					})
				} else if types == "subchannel" {
					q = q.WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
						return q.Where("type_child_id = ?", key)
					})
				}
			}
			return q
		}).
		Order("published_date desc").
		Limit(limit).
		Offset(offset).
		Scan(ctx); err != nil {
		return nil, err
	}

	return content, nil
}

func (repo *contentRepository) GetAllMultimediaRow(ctx context.Context, multimediaType string, types string, key string, limit int, offset int) ([]*entity.ContentRowResponse, error) {
	content := []*entity.ContentRowResponse{}

	if err := repo.DB.NewSelect().Model(&content).
		Apply(func(q *bun.SelectQuery) *bun.SelectQuery { //Relation Function
			q = q.Relation("Channel").Relation("SubChannel").Relation("Region")
			if multimediaType == "photo" {
				q = q.Relation("SubPhotos")
			}
			return q
		}).
		Where("content_row_response.is_active = ?", true).
		Where("content_row_response.type = ?", multimediaType).
		WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Where("ads_position is null").WhereOr("ads_position = 0")
		}).
		Apply(func(q *bun.SelectQuery) *bun.SelectQuery {
			if types != "" {
				if types == "channel" {
					q = q.WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
						return q.Where("type_id = ?", key).Where("type_child_id is not null")
					})
				} else if types == "region" {
					q = q.WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
						return q.Where("type_id = ?", key).Where("type_child_id is null")
					})
				} else if types == "subchannel" {
					q = q.WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
						return q.Where("type_child_id = ?", key)
					})
				}
			}
			return q
		}).
		Order("published_date desc").
		Limit(limit).
		Offset(offset).
		Scan(ctx); err != nil {
		return nil, err
	}

	return content, nil
}

func (repo *contentRepository) GetAllArticleRow(ctx context.Context, limit int, offset int) ([]*entity.ContentRowResponse, error) {
	content := []*entity.ContentRowResponse{}

	if err := repo.DB.NewSelect().Model(&content).
		Relation("Channel").Relation("SubChannel").
		Where("content_row_response.is_active = ?", true).
		Where("content_row_response.type = ?", "channel").
		Where("content_row_response.suplemen_id = 1").
		WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Where("ads_position is null").WhereOr("ads_position = 0")
		}).
		Order("published_date desc").
		Limit(limit).
		Offset(offset).
		Scan(ctx); err != nil {
		return nil, err
	}

	return content, nil
}
func (repo *contentRepository) GetAllEditorChoiceRow(ctx context.Context, limit int, offset int) ([]*entity.ContentRowResponse, error) {
	content := []*entity.ContentRowResponse{}

	if err := repo.DB.NewSelect().Model(&content).
		Relation("Channel").Relation("SubChannel").Relation("Region").
		Where("content_row_response.is_active = ?", true).
		Where("content_row_response.type IN (?)", bun.In([]string{"channel", "region"})).
		WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Where("ads_position is null").WhereOr("ads_position = 0")
		}).
		Where("content_row_response.is_editor_choice = 1").
		Where("content_row_response.is_national = 1").
		Order("published_date desc").
		Limit(limit).
		Offset(offset).
		Scan(ctx); err != nil {
		return nil, err
	}

	return content, nil
}

func (repo *contentRepository) GetAllIndeks(ctx context.Context, types string, key string, date string, limit int, offset int) ([]*entity.ContentRowResponse, error) {
	content := []*entity.ContentRowResponse{}

	if err := repo.DB.NewSelect().Model(&content).
		Relation("Region").
		Relation("Channel").
		Relation("SubChannel").
		Relation("SubPhotos").
		Where("content_row_response.is_active = ?", true).
		Apply(func(q *bun.SelectQuery) *bun.SelectQuery {
			if date != "" && date != "all" {
				startDate, _ := time.Parse("2006-01-02", date)
				endDate := startDate.AddDate(0, 0, 1)
				q = q.Where("published_date between ? and ?", startDate, endDate)
			}
			if types != "" && types != "all" {
				if types == "channel" {
					if key == "video" || key == "photo" || key == "infografis" {
						q = q.Where("type = ?", key)
					} else if key == "multimedia" {
						q = q.Where("type IN (?)", bun.In([]string{"video", "photo", "infografis"}))
					} else {
						q = q.WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
							return q.Where("type_id = ?", key).Where("type_child_id is not null")
						})
					}
				} else if types == "region" {
					q = q.WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
						return q.Where("type_id = ?", key).Where("type_child_id is null")
					})
				} else if types == "suplemen" {
					q = q.WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
						return q.Where("suplemen_id = ?", key).Where("type_id is not null")
					})
				}
			}
			return q
		}).
		Order("published_date desc").
		Limit(limit).
		Offset(offset).
		Scan(ctx); err != nil {
		return nil, err
	}

	return content, nil
}

func (repo *contentRepository) GetAllSearch(ctx context.Context, types string, key interface{}, limit int, offset int) ([]*entity.ContentRowResponse, error) {
	content := []*entity.ContentRowResponse{}

	if err := repo.DB.NewSelect().Model(&content).
		Relation("Region").
		Relation("Channel").
		Relation("SubChannel").
		Relation("SubPhotos").
		Where("content_row_response.is_active = ?", true).
		Apply(func(q *bun.SelectQuery) *bun.SelectQuery {
			if types == "keyword" {
				sliceKey := []string{"%", fmt.Sprintf("%s", key), "%"}
				q = q.Where("? like ?", bun.Ident("content_row_response.title"), strings.Join(sliceKey, ""))
			} else if types == "tag" || types == "tag-headline" {
				q = q.Where("content_row_response.id IN (?)", bun.In(key))
				if types == "tag-headline" {
					q = q.Where("headline_type IN (?)", bun.In([]int{1, 2}))
				} else {
					q = q.Where("headline_type != 2")
				}
			} else if types == "topic" || types == "topic-headline" || types == "topic-headline-subkanal" {
				q = q.Where("content_row_response.id IN (?)", bun.In(key))
				if types == "topic-headline" {
					q = q.Where("headline_type IN (?)", bun.In([]int{1, 2}))
				} else if types == "topic-headline-subkanal" {
					q = q.Where("headline_type IN (?)", bun.In([]int{1, 2, 3}))
				} else {
					q = q.Where("headline_type != 2")
				}
			}
			return q
		}).
		Order("published_date desc").
		Limit(limit).
		Offset(offset).
		Scan(ctx); err != nil {
		return nil, err
	}

	return content, nil
}
