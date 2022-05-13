package domain

import (
	"github.com/ivanauliaa/go-appoinment/src/model"
	"github.com/labstack/echo/v4"
)

type UsersHandler interface {
	HelloHandler(c echo.Context) error
}

type UsersService interface {
	Hello() string
}

type UsersRepository interface {
	AddUser(payload model.User) (int, error)
	VerifyNewUserEmail(email string) (int, error)
	VerifyUserCredential(email string, password string) (uint, int, error)
	Hello() string
}