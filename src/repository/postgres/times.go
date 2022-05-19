package postgres

import (
	"fmt"
	"net/http"

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

func (r *timesRepository) AddTime(payload model.Time) (model.PostTimeResponse, int, error) {
	result := r.db.Create(&payload)

	if result.RowsAffected == 0 {
		return model.PostTimeResponse{}, http.StatusInternalServerError, result.Error
	}

	timeID := model.PostTimeResponse{
		TimeID: payload.ID,
	}

	return timeID, http.StatusOK, nil
}

func (r *timesRepository) VerifyTime(timeID uint) (int, error) {
	result := r.db.First(&model.Time{}, "id = ?", timeID)

	if result.RowsAffected == 0 {
		return http.StatusNotFound, fmt.Errorf("time not found")
	}

	return http.StatusOK, nil
}

func (r *timesRepository) VerifyTimeDateID(timeID uint, dateID uint) (int, error) {
	time := model.Time{}
	result := r.db.First(&time, "id = ?", timeID)

	if result.RowsAffected == 0 {
		return http.StatusNotFound, fmt.Errorf("time not found")
	}

	if time.DateID != dateID {
		return http.StatusBadRequest, fmt.Errorf("time not belongs to related date")
	}

	return http.StatusOK, nil
}

func (r *timesRepository) GetTime(timeID uint) (model.Time, int, error) {
	time := model.Time{}
	result := r.db.First(&time, "id = ?", timeID)

	if result.RowsAffected == 0 {
		return model.Time{}, http.StatusNotFound, fmt.Errorf("time not found")
	}

	return time, http.StatusOK, nil
}

func (r *timesRepository) GetTimes(dateID uint) ([]model.Time, int, error) {
	times := []model.Time{}
	r.db.Where("date_id = ?", dateID).Find(&times)

	return times, http.StatusOK, nil
}
