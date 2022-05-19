package domain

import (
	"github.com/ivanauliaa/go-appoinment/src/model"
	"github.com/labstack/echo/v4"
)

type EventsHandler interface {
	GetEventHandler(c echo.Context) error
	PostSendEventHandler(c echo.Context) error
}

type EventsService interface {
	GetEvent(payload model.GetEventPayload, requestHeader model.RequestHeader) (model.EventWithRelation, int, error)
	SendEvent(payload model.SendEventPayload) (int, error)
}

type EventsRepository interface {
	AddEvent(payload model.Event) (int, error)
	VerifyNewEvent(appointmentID uint) (int, error)
	GetEvent(appointmentID uint) (model.Event, int, error)
	VerifyEventParticipant(appointmentID uint, userID uint) (int, error)
}

type EventsAPI interface {
	SendEvent(payload model.SendEvent) (int, error)
}
