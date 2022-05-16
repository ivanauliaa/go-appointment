package postgres

import (
	"fmt"
	"net/http"

	"github.com/ivanauliaa/go-appoinment/src/domain"
	"github.com/ivanauliaa/go-appoinment/src/model"

	"gorm.io/gorm"
)

type usersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(d *gorm.DB) domain.UsersRepository {
	newRepository := usersRepository{
		db: d,
	}

	return &newRepository
}

func (r *usersRepository) AddUser(payload model.User) (int, error) {
	result := r.db.Create(&payload)

	if result.RowsAffected == 0 {
		return http.StatusInternalServerError, result.Error
	}

	return http.StatusOK, nil
}

func (r *usersRepository) VerifyNewUserEmail(email string) (int, error) {
	result := r.db.First(&model.User{}, "email = ?", email)

	if result.RowsAffected > 0 {
		return http.StatusBadRequest, fmt.Errorf("email already registered")
	}

	return http.StatusOK, nil
}

func (r *usersRepository) VerifyUserCredential(email string, password string) (uint, int, error) {
	user := model.User{}
	result := r.db.First(&user, "email = ? AND password = ?", email, password)

	if result.RowsAffected == 0 {
		return 0, http.StatusNotFound, fmt.Errorf("invalid email or password")
	}

	return user.ID, http.StatusOK, nil
}

func (r *usersRepository) Hello() string {
	return "hello world"
}

func (r *usersRepository) VerifyUser(userID uint) (int, error) {
	result := r.db.First(&model.User{}, "id = ?", userID)

	if result.RowsAffected == 0 {
		return http.StatusNotFound, fmt.Errorf("user not found")
	}

	return http.StatusOK, nil
}
