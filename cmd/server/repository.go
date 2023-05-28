package server

import (
	"github.com/zakariawahyu/go-echo-news/config"
	_contentRepository "github.com/zakariawahyu/go-echo-news/modules/content/repository"
)

type Repository struct {
	ContentRepo _contentRepository.ContentRepository
}

func NewRepository(conn *config.Conn) *Repository {
	return &Repository{
		ContentRepo: _contentRepository.NewContentRepository(conn.Mysql),
	}
}
