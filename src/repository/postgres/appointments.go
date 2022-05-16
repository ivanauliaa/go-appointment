package postgres

import (
	"fmt"
	"net/http"

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

func (r *appointmentsRepository) AddAppointment(payload model.Appointment) (model.PostAppointmentResponse, int, error) {
	result := r.db.Create(&payload)

	if result.RowsAffected == 0 {
		return model.PostAppointmentResponse{}, http.StatusInternalServerError, result.Error
	}

	appointmentID := model.PostAppointmentResponse{
		AppointmentID: payload.ID,
	}

	return appointmentID, http.StatusOK, nil
}

func (r *appointmentsRepository) VerifyAppointment(appointmentID uint) (int, error) {
	result := r.db.First(&model.Appointment{}, "id = ?", appointmentID)

	if result.RowsAffected == 0 {
		return http.StatusNotFound, fmt.Errorf("appointment not found")
	}

	return http.StatusOK, nil
}

func (r *appointmentsRepository) VerifyAppointmentOwner(appointmentID uint, userID uint) (int, error) {
	appointment := model.Appointment{}
	result := r.db.First(&appointment, "id = ?", appointmentID)

	if result.RowsAffected == 0 {
		return http.StatusNotFound, fmt.Errorf("appointment not found")
	}

	if appointment.UserID != userID {
		return http.StatusForbidden, fmt.Errorf("restricted resource")
	}

	return http.StatusOK, nil
}

func (r *appointmentsRepository) VerifyAppointmentGuest(appointmentID uint, userID uint) (int, error) {
	appointment := model.Appointment{}
	result := r.db.First(&appointment, "id = ?", appointmentID)

	if result.RowsAffected == 0 {
		return http.StatusNotFound, fmt.Errorf("appointment not found")
	}

	if appointment.UserID == userID {
		return http.StatusBadRequest, fmt.Errorf("you are owner of this appointment, you can't be the guest")
	}

	return http.StatusOK, nil
}

func (r *appointmentsRepository) GetAppointment(appointmentID uint) (model.Appointment, int, error) {
	appointment := model.Appointment{}
	result := r.db.First(&appointment, "id = ?", appointmentID)

	if result.RowsAffected == 0 {
		return model.Appointment{}, http.StatusNotFound, fmt.Errorf("appointment not found")
	}

	return appointment, http.StatusOK, nil
}
