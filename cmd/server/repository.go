package server

import (
	_contentRepository "github.com/zakariawahyu/go-echo-news/modules/content/repository"
	"github.com/zakariawahyu/go-echo-news/utils"
)

type Repository struct {
	ContentRepo _contentRepository.ContentRepository
}

func NewRepository(conn *utils.Conn) *Repository {
	return &Repository{
		ContentRepo: _contentRepository.NewContentRepository(conn.Mysql),
	}
}
