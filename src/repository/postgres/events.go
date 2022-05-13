package postgres

import (
	"github.com/ivanauliaa/go-appoinment/src/domain"
	"github.com/ivanauliaa/go-appoinment/src/model"
	"gorm.io/gorm"
)

type eventsRepository struct {
	db *gorm.DB
}

func NewEventsRepository(d *gorm.DB) domain.EventsRepository {
	newRepository := eventsRepository{
		db: d,
	}

	return &newRepository
}

func (r *eventsRepository) AddEvent(payload model.Event) (int, error) {
	return 0, nil
}
