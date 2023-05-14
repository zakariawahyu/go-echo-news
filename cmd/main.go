package main

import (
	"github.com/zakariawahyu/go-echo-news/cmd/server"
	"github.com/zakariawahyu/go-echo-news/config"
	"github.com/zakariawahyu/go-echo-news/utils"
	"log"
)

func main() {
	cfg := config.NewConfig()
	db := utils.NewDbConnection(cfg)

	defer func() {
		err := db.Mysql.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	repo := server.NewRepository(db)
	serv := server.NewServices(repo, cfg.App.ContextTimeout)
	server.NewHandler(cfg, serv)
}
