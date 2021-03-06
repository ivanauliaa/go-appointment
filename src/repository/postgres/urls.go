package postgres

import (
	"fmt"
	"net/http"

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

func (r *urlsRepository) AddURL(payload model.URL) (model.PostURLResponse, int, error) {
	result := r.db.Create(&payload)

	if result.RowsAffected == 0 {
		return model.PostURLResponse{}, http.StatusInternalServerError, result.Error
	}

	appointmentURL := model.PostURLResponse{
		AppointmentURL: payload.URL,
	}

	return appointmentURL, http.StatusOK, nil
}

func (r *urlsRepository) VerifyNewURL(appointmentID uint) (int, error) {
	result := r.db.First(&model.URL{}, "appointment_id = ?", appointmentID)

	if result.RowsAffected > 0 {
		return http.StatusBadRequest, fmt.Errorf("url for related appointment already exist, try get request instead to get it")
	}

	return http.StatusOK, nil
}

func (r *urlsRepository) GetURL(appointmentID uint) (model.URL, int, error) {
	url := model.URL{}
	result := r.db.First(&url, "appointment_id = ?", appointmentID)

	if result.RowsAffected == 0 {
		return model.URL{}, http.StatusBadRequest, fmt.Errorf("url for related appointment not found, try post request instead to create it")
	}

	return url, http.StatusOK, nil
}
