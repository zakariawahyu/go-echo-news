package config

import (
	"github.com/spf13/viper"
	"time"
)

type AppConfig struct {
	Name           string
	Version        string
	ContextTimeout time.Duration
}

func LoadAppConfig() AppConfig {
	return AppConfig{
		Name:           viper.GetString("APP_NAME"),
		Version:        viper.GetString("APP_VERSION"),
		ContextTimeout: viper.GetDuration("APP_TIMEOUT") * time.Second,
	}
}
