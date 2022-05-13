package postgres

import (
	"fmt"
	"net/http"

	"github.com/ivanauliaa/go-appoinment/src/domain"
	"github.com/ivanauliaa/go-appoinment/src/model"
	"gorm.io/gorm"
)

type authenticationsRepository struct {
	db *gorm.DB
}

func NewAuthenticationsRepository(d *gorm.DB) domain.AuthenticationsRepository {
	newRepository := authenticationsRepository{
		db: d,
	}

	return &newRepository
}

func (r *authenticationsRepository) AddRefreshToken(token model.Authentication) (int, error) {
	result := r.db.Create(&token)

	if result.RowsAffected == 0 {
		return http.StatusInternalServerError, result.Error
	}

	return http.StatusOK, nil
}

func (r *authenticationsRepository) VerifyRefreshToken(token model.RefreshTokenPayload) (int, error) {
	auth := model.Authentication{}
	result := r.db.First(&auth, "token = ?", token.RefreshToken)

	if result.RowsAffected == 0 {
		return http.StatusNotFound, fmt.Errorf("refresh token not found")
	}

	return http.StatusOK, nil
}

func (r *authenticationsRepository) DeleteRefreshToken(token model.RefreshTokenPayload) (int, error) {
	result := r.db.Delete(&model.Authentication{}, "token = ?", token.RefreshToken)

	if result.RowsAffected == 0 {
		return http.StatusNotFound, fmt.Errorf("refresh token not found")
	}
	return http.StatusOK, nil
}
