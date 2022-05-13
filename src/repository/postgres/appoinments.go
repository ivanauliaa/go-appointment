package postgres

import (
	"github.com/ivanauliaa/go-appoinment/src/domain"
	"github.com/ivanauliaa/go-appoinment/src/model"
	"gorm.io/gorm"
)

type appointmentsRepository struct {
	db *gorm.DB
}

func NewAppointmentsRepository(d *gorm.DB) domain.AppointmentsRepository {
	newRepository := appointmentsRepository{
		db: d,
	}

	return &newRepository
}

func (r *appointmentsRepository) AddAppointment(payload model.Appointment) (model.Appointment, int, error) {
	return model.Appointment{}, 0, nil
}
