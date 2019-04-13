package service

import (
	"context"
	"github.com/Yurovskikh/news/storage/pkg/model"
	"github.com/Yurovskikh/news/storage/pkg/repository"
)

type NewsService interface {
	Create(ctx context.Context, news *model.News) (*model.News, error)
	GetById(ctx context.Context, id uint) (*model.News, error)
}
type newsService struct {
	repository repository.NewsRepository
}

func NewNewsService(repository repository.NewsRepository) NewsService {
	return &newsService{
		repository: repository,
	}
}

func (s *newsService) Create(ctx context.Context, news *model.News) (*model.News, error) {
	return s.repository.Save(ctx, news)
}

func (s *newsService) GetById(ctx context.Context, id uint) (*model.News, error) {
	return s.repository.FindById(ctx, id)
}
