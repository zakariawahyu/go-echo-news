package main

import (
	"github.com/zakariawahyu/go-echo-news/cmd/server"
	"github.com/zakariawahyu/go-echo-news/config"
	"github.com/zakariawahyu/go-echo-news/pkg/db"
	"github.com/zakariawahyu/go-echo-news/pkg/exception"
	"github.com/zakariawahyu/go-echo-news/pkg/logger"
)

func main() {
	cfg := config.NewConfig()
	db := db.NewDbConnection(cfg)

	defer func() {
		err := db.Mysql.Close()
		exception.PanicIfNeeded(err)
	}()

	repo := server.NewRepository(db)
	serv := server.NewServices(repo, cfg.App.ContextTimeout)
	appLogger := logger.NewApiLogger(cfg)
	appLogger.InitLogger()

	server.NewHandler(cfg, serv, appLogger)
}
