package db

import (
	"github.com/redis/go-redis/v9"
	"github.com/zakariawahyu/go-echo-news/config"
)

func InitRedis(cfg *config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Cache.DSN,
		Password: cfg.Cache.RedisPassword,
	})

	return client
}
