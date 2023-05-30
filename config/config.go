package config

import (
	"github.com/spf13/viper"
	"github.com/zakariawahyu/go-echo-news/pkg/exception"
)

type Config struct {
	DB  DBConfig
	App AppConfig
}

func NewConfig() *Config {
	viper.AddConfigPath("../")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	exception.PanicIfNeeded(err)

	return &Config{
		DB:  LoadDBConfig(),
		App: LoadAppConfig(),
	}
}