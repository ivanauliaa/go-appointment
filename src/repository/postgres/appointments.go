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

func (r *appointmentsRepository) GetAppointments(credentialID uint) ([]model.Appointment, int, error) {
	appointments := []model.Appointment{}
	r.db.Where("user_id = ?", credentialID).Find(&appointments)

	return appointments, http.StatusOK, nil
}

func (r *appointmentsRepository) UpdateAppointment(payload model.Appointment, appointmentID uint) (int, error) {
	appointment := model.Appointment{}
	result := r.db.First(&appointment, "id = ?", appointmentID)

	if result.RowsAffected == 0 {
		return http.StatusNotFound, fmt.Errorf("appointment not found")
	}

	appointment.Name = payload.Name
	appointment.Room = payload.Room

	result = r.db.Save(&appointment)

	if result.RowsAffected == 0 {
		return http.StatusInternalServerError, result.Error
	}

	return http.StatusOK, nil
}

func (r *appointmentsRepository) DeleteAppointment(appointmentID uint) (int, error) {
	result := r.db.Where("id = ?", appointmentID).Delete(&model.Appointment{})

	if result.RowsAffected == 0 {
		return http.StatusInternalServerError, result.Error
	}

	return http.StatusOK, nil
}
