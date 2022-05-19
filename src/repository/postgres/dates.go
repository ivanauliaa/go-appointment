package postgres

import (
	"fmt"
	"net/http"

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

func (r *datesRepository) AddDate(payload model.Date) (model.PostDateResponse, int, error) {
	result := r.db.Create(&payload)

	if result.RowsAffected == 0 {
		return model.PostDateResponse{}, http.StatusInternalServerError, result.Error
	}

	dateID := model.PostDateResponse{
		DateID: payload.ID,
	}

	return dateID, http.StatusOK, nil
}

func (r *datesRepository) VerifyDate(dateID uint) (int, error) {
	result := r.db.First(&model.Date{}, "id = ?", dateID)

	if result.RowsAffected == 0 {
		return http.StatusNotFound, fmt.Errorf("date not found")
	}

	return http.StatusOK, nil
}

func (r *datesRepository) VerifyDateAppointmentID(dateID uint, appointmentID uint) (int, error) {
	date := model.Date{}
	result := r.db.First(&date, "id = ?", dateID)

	if result.RowsAffected == 0 {
		return http.StatusNotFound, fmt.Errorf("date not found")
	}

	if date.AppointmentID != appointmentID {
		return http.StatusBadRequest, fmt.Errorf("date not belongs to related appointment")
	}

	return http.StatusOK, nil
}

func (r *datesRepository) GetDate(dateID uint) (model.Date, int, error) {
	date := model.Date{}
	result := r.db.First(&date, "id = ?", dateID)

	if result.RowsAffected == 0 {
		return model.Date{}, http.StatusNotFound, fmt.Errorf("date not found")
	}

	return date, http.StatusOK, nil
}

func (r *datesRepository) GetDates(appointmentID uint) ([]model.Date, int, error) {
	dates := []model.Date{}
	r.db.Where("appointment_id = ?", appointmentID).Find(&dates)

	return dates, http.StatusOK, nil
}
