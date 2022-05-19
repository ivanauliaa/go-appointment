package domain

import (
	"github.com/ivanauliaa/go-appoinment/src/model"
	"github.com/labstack/echo/v4"
)

type DatesHandler interface {
	PostDateHandler(c echo.Context) error
}

type DatesService interface {
	AddDate(
		payload model.Date,
		requestHeader model.RequestHeader,
	) (model.PostDateResponse, int, error)
}

type DatesRepository interface {
	AddDate(payload model.Date) (model.PostDateResponse, int, error)
	VerifyDate(dateID uint) (int, error)
	VerifyDateAppointmentID(dateID uint, appointmentID uint) (int, error)
	GetDate(dateID uint) (model.Date, int, error)
	GetDates(appointmentID uint) ([]model.Date, int, error)
}
