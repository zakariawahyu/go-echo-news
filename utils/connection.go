package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zakariawahyu/go-echo-news/config"
	"log"
	"time"
)

type Conn struct {
	Mysql *sql.DB
}

func NewDbConnection(cfg *config.Config) *Conn {
	return &Conn{
		Mysql: InitMysql(cfg),
	}
}

func InitMysql(cfg *config.Config) *sql.DB {
	db, err := sql.Open("mysql", cfg.DB.DSN)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxIdleConns(cfg.DB.MaxIdleConns)
	db.SetMaxOpenConns(cfg.DB.MaxOpenConns)
	db.SetConnMaxIdleTime(cfg.DB.ConnMaxIdleTime * time.Minute)
	db.SetConnMaxLifetime(cfg.DB.ConnMaxLifetime * time.Minute)

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
