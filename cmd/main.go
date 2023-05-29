package main

import (
	"github.com/zakariawahyu/go-echo-news/cmd/server"
	"github.com/zakariawahyu/go-echo-news/config"
	"github.com/zakariawahyu/go-echo-news/pkg"
	"github.com/zakariawahyu/go-echo-news/pkg/exception"
)

func main() {
	cfg := config.NewConfig()
	db := pkg.NewDbConnection(cfg)

	defer func() {
		err := db.Mysql.Close()
		exception.PanicIfNeeded(err)
	}()

	repo := server.NewRepository(db)
	serv := server.NewServices(repo, cfg.App.ContextTimeout)
	server.NewHandler(cfg, serv)
}
