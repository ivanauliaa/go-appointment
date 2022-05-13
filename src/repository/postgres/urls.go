package postgres

import (
	"github.com/ivanauliaa/go-appoinment/src/domain"
	"github.com/ivanauliaa/go-appoinment/src/model"
	"gorm.io/gorm"
)

type urlsRepository struct {
	db *gorm.DB
}

func NewURLsRepository(d *gorm.DB) domain.URLsRepository {
	newRepository := urlsRepository{
		db: d,
	}

	return &newRepository
}

func (r *urlsRepository) AddURL(payload model.URL) (string, int, error) {
	return "", 0, nil
}
