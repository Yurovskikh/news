package repository

import (
	"context"
	"github.com/Yurovskikh/news/storage/pkg/model"
	"github.com/jinzhu/gorm"
)

type NewsRepository interface {
	Save(ctx context.Context, news *model.News) (*model.News, error)
	FindById(ctx context.Context, id uint) (*model.News, error)
}
type newsRepository struct {
	db *gorm.DB
}

func NewNewsRepository(db *gorm.DB) NewsRepository {
	return &newsRepository{
		db: db,
	}
}

func (r *newsRepository) Save(ctx context.Context, news *model.News) (*model.News, error) {
	tx := r.db.Begin()
	err := tx.Save(news).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return news, tx.Commit().Error
}

func (r *newsRepository) FindById(ctx context.Context, id uint) (*model.News, error) {
	var news model.News
	err := r.db.First(&news, id).Error
	if err != nil {
		return nil, err
	}
	return &news, nil
}
