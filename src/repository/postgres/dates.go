package postgres

import (
	"github.com/ivanauliaa/go-appoinment/src/domain"
	"github.com/ivanauliaa/go-appoinment/src/model"
	"gorm.io/gorm"
)

type datesRepository struct {
	db *gorm.DB
}

func NewDatesRepository(d *gorm.DB) domain.DatesRepository {
	newRepository := datesRepository{
		db: d,
	}

	return &newRepository
}

func (r *datesRepository) AddDate(payload model.Date) (int, error) {
	return 0, nil
}
