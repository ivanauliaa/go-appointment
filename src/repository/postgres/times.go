package postgres

import (
	"github.com/ivanauliaa/go-appoinment/src/domain"
	"github.com/ivanauliaa/go-appoinment/src/model"
	"gorm.io/gorm"
)

type timesRepository struct {
	db *gorm.DB
}

func NewTimesRepository(d *gorm.DB) domain.TimesRepository {
	newRepository := timesRepository{
		db: d,
	}

	return &newRepository
}

func (r *timesRepository) AddTime(payload model.Time) (int, error) {
	return 0, nil
}
