package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	DB  DBConfig
	App AppConfig
}

func NewConfig() *Config {
	// Tell viper the type of your file
	viper.AddConfigPath("../")

	viper.SetConfigName("app")
	// Tell viper the type of your file
	viper.SetConfigType("env")

	// Viper reads all the variables from app.env file and log error if any found
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading app.env file", err)
	}

	return &Config{
		DB:  LoadDBConfig(),
		App: LoadAppConfig(),
	}
}
