package postgres

import (
	"fmt"
	"net/http"

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
	result := r.db.Create(&payload)

	if result.RowsAffected == 0 {
		return http.StatusInternalServerError, result.Error
	}

	return http.StatusOK, nil
}

func (r *eventsRepository) VerifyNewEvent(appointmentID uint) (int, error) {
	event := model.Event{}
	result := r.db.First(&event, "appointment_id = ?", appointmentID)

	if result.RowsAffected > 0 {
		return http.StatusBadRequest, fmt.Errorf("appointment has fulfilled")
	}

	return http.StatusOK, nil
}

func (r *eventsRepository) GetEvent(appointmentID uint) (model.Event, int, error) {
	event := model.Event{}
	result := r.db.First(&event, "appointment_id = ?", appointmentID)

	if result.RowsAffected == 0 {
		return model.Event{}, http.StatusBadRequest, fmt.Errorf("appointment not found")
	}

	return event, http.StatusOK, nil
}

func (r *eventsRepository) VerifyEventParticipant(appointmentID uint, userID uint) (int, error) {
	event := model.Event{}
	result := r.db.First(&event, "appointment_id = ?", appointmentID)

	if result.RowsAffected == 0 {
		return http.StatusBadRequest, fmt.Errorf("appointment not found")
	}

	if event.HostID != userID && event.GuestID != userID {
		return http.StatusForbidden, fmt.Errorf("you aren't participants in this event")
	}

	return http.StatusOK, nil
}
