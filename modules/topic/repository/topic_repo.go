package repository

import (
	"context"
	"github.com/uptrace/bun"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/topic"
)

type topicRepository struct {
	DB *bun.DB
}

func NewTopicRepository(DB *bun.DB) topic.TopicRepository {
	return &topicRepository{
		DB: DB,
	}
}

func (repo *topicRepository) GetBySlugOrID(ctx context.Context, slug string) (*entity.Topic, error) {
	topic := &entity.Topic{}

	if err := repo.DB.NewSelect().Model(topic).Where("slug = ?", slug).WhereOr("id = ?", slug).Scan(ctx); err != nil {
		return nil, err
	}

	return topic, nil
}
