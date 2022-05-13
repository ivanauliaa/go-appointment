package domain

import (
	"github.com/ivanauliaa/go-appoinment/src/model"
	"github.com/labstack/echo/v4"
)

type AppointmentsHandler interface {
	PostAppointmentHandler(c echo.Context) error
}

type AppointmentsService interface {
	AddAppointment(
		appointmentPayload model.Appointment,
		eventPayload model.Event,
		datePayload model.Date,
		timePayload model.Time,
		requestHeader model.RequestHeader,
	) (model.Appointment, int, error)
}

type AppointmentsRepository interface {
	AddAppointment(payload model.Appointment) (model.Appointment, int, error)
}
