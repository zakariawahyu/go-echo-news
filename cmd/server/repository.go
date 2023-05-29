package server

import (
	_contentRepository "github.com/zakariawahyu/go-echo-news/modules/content/repository"
	_recommendedRepository "github.com/zakariawahyu/go-echo-news/modules/recommended/repository"
	"github.com/zakariawahyu/go-echo-news/pkg"
)

type Repository struct {
	contentRepo     _contentRepository.ContentRepository
	recommendedRepo _recommendedRepository.RecommendedRepository
}

func NewRepository(conn *pkg.Conn) *Repository {
	return &Repository{
		contentRepo:     _contentRepository.NewContentRepository(conn.Mysql),
		recommendedRepo: _recommendedRepository.NewRecommendedRepository(conn.Mysql),
	}
}
