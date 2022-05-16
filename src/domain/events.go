package domain

import "github.com/ivanauliaa/go-appoinment/src/model"

type EventsHandler interface {
}

type EventsService interface {
}

type EventsRepository interface {
	AddEvent(payload model.Event) (int, error)
	VerifyNewEvent(appointmentID uint) (int, error)
}
