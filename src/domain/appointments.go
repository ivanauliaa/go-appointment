package domain

import (
	"github.com/ivanauliaa/go-appoinment/src/model"
	"github.com/labstack/echo/v4"
)

type AppointmentsHandler interface {
	PostAppointmentHandler(c echo.Context) error
	PostAppointmentConfirmHandler(c echo.Context) error
	GetAppointmentsHandler(c echo.Context) error
	GetAppointmentHandler(c echo.Context) error
}

type AppointmentsService interface {
	AddAppointment(
		payload model.PostAppointmentPayload,
		requestHeader model.RequestHeader,
	) (model.PostAppointmentResponse, int, error)
	ConfirmAppointment(
		payload model.PostAppointmentConfirmPayload,
		requestHeader model.RequestHeader,
	) (int, error)
	GetAppointments(requestHeader model.RequestHeader) ([]model.Appointment, int, error)
	GetAppointment(
		payload model.GetAppointmentPayload,
		requestHeader model.RequestHeader,
	) (model.AppointmentWithRelation, int, error)
}

type AppointmentsRepository interface {
	AddAppointment(payload model.Appointment) (model.PostAppointmentResponse, int, error)
	VerifyAppointment(appointmentID uint) (int, error)
	VerifyAppointmentOwner(appointmentID uint, userID uint) (int, error)
	VerifyAppointmentGuest(appointmentID uint, userID uint) (int, error)
	GetAppointment(appointmentID uint) (model.Appointment, int, error)
	GetAppointments(credentialID uint) ([]model.Appointment, int, error)
}
