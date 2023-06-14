package repository

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/channel"
	"github.com/zakariawahyu/go-echo-news/pkg/helpers"
)

type channelRedisRepository struct {
	Redis *redis.Client
}

func NewChannelRedisRepository(Redis *redis.Client) channel.ChannelRedisRepository {
	return &channelRedisRepository{
		Redis: Redis,
	}
}

func (repo *channelRedisRepository) GetChannel(ctx context.Context, key string) (*entity.Channel, error) {
	channelBytes, err := repo.Redis.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	channel := &entity.Channel{}
	if err = json.Unmarshal(channelBytes, channel); err != nil {
		return nil, err
	}

	return channel, nil
}

func (repo *channelRedisRepository) SetChannel(ctx context.Context, key string, ttl int, channel *entity.Channel) error {
	channelBytes, err := json.Marshal(channel)
	if err != nil {
		return err
	}

	if err = repo.Redis.Set(ctx, key, channelBytes, helpers.TtlRedis(ttl)).Err(); err != nil {
		return err
	}

	return nil
}
