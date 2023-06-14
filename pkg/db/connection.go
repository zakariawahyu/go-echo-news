package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"github.com/uptrace/bun"
	"github.com/zakariawahyu/go-echo-news/config"
)

type Conn struct {
	Mysql *bun.DB
	Redis *redis.Client
}

func NewDbConnection(cfg *config.Config) *Conn {
	return &Conn{
		Mysql: InitMysql(cfg),
		Redis: InitRedis(cfg),
	}
}
