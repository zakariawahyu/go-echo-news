package config

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type DBConfig struct {
	DSN             string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxIdleTime time.Duration
	ConnMaxLifetime time.Duration
}

func LoadDBConfig() DBConfig {
	return DBConfig{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			viper.GetString("DB_USERNAME"),
			viper.GetString("DB_PASSWORD"),
			viper.GetString("DB_HOST"),
			viper.GetString("DB_PORT"),
			viper.GetString("DB_NAME"),
		),
		MaxIdleConns:    viper.GetInt("DB_MAX_IDLE_CONNS"),
		MaxOpenConns:    viper.GetInt("DB_MAX_OPEN_CONNS"),
		ConnMaxIdleTime: viper.GetDuration("DB_CONN_MAX_IDLE_TIME") * time.Second,
		ConnMaxLifetime: viper.GetDuration("DB_CONN_MAX_LIFETIME") * time.Second,
	}
}
