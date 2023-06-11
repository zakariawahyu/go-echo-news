package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type CacheConfig struct {
	RedisHost     string
	RedisPassword string
	RedisPort     string
	DSN           string
}

func LoadCacheConfig() CacheConfig {
	return CacheConfig{
		RedisHost:     viper.GetString("REDIS_HOST"),
		RedisPassword: viper.GetString("REDIS_PASSWORD"),
		RedisPort:     viper.GetString("REDIS_PORT"),
		DSN:           fmt.Sprintf("%s:%s", viper.GetString("REDIS_HOST"), viper.GetString("REDIS_PORT")),
	}
}
