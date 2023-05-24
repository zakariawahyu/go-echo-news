package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"
	"github.com/zakariawahyu/go-echo-news/config"
	"github.com/zakariawahyu/go-echo-news/entity"
	"log"
	"time"
)

type Conn struct {
	Mysql *bun.DB
}

func NewDbConnection(cfg *config.Config) *Conn {
	return &Conn{
		Mysql: InitMysql(cfg),
	}
}

func InitMysql(cfg *config.Config) *bun.DB {
	sqldb, err := sql.Open("mysql", cfg.DB.DSN)
	if err != nil {
		log.Fatal(err)
	}

	db := bun.NewDB(sqldb, mysqldialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	db.SetMaxIdleConns(cfg.DB.MaxIdleConns)
	db.SetMaxOpenConns(cfg.DB.MaxOpenConns)
	db.SetConnMaxIdleTime(cfg.DB.ConnMaxIdleTime * time.Minute)
	db.SetConnMaxLifetime(cfg.DB.ConnMaxLifetime * time.Minute)

	db.RegisterModel((*entity.ContentHasTag)(nil))
	db.RegisterModel((*entity.ContentHasTopic)(nil))
	db.RegisterModel((*entity.ContentHasReporter)(nil))

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
