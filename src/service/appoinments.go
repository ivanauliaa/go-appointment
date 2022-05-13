package service

import (
	"github.com/ivanauliaa/go-appoinment/src/domain"
	"github.com/ivanauliaa/go-appoinment/src/model"
)

type appointmentsService struct {
	appointmentsRepository domain.AppointmentsRepository
	eventsRepository       domain.EventsRepository
	datesRepository        domain.DatesRepository
	timesRepository        domain.TimesRepository
	urlsRepository         domain.URLsRepository
}

func NewAppointmentsService(
	ar domain.AppointmentsRepository,
	er domain.EventsRepository,
	dr domain.DatesRepository,
	tr domain.TimesRepository,
	ur domain.URLsRepository,
) domain.AppointmentsService {
	newService := appointmentsService{
		appointmentsRepository: ar,
		eventsRepository:       er,
		datesRepository:        dr,
		timesRepository:        tr,
		urlsRepository:         ur,
	}

	return &newService
}

func (s *appointmentsService) AddAppointment(
	appointmentPayload model.Appointment,
	eventPayload model.Event,
	datePayload model.Date,
	timePayload model.Time,
) (model.Appointment, int, error) {
	// TODO: bikin appointment
	s.appointmentsRepository.AddAppointment(appointmentPayload)

	// TODO: bikin event
	s.eventsRepository.AddEvent(eventPayload)

	// TODO: bikin date
	s.datesRepository.AddDate(datePayload)

	// TODO: bikin time
	s.timesRepository.AddTime(timePayload)

	// TODO: bikin url
	// urlPayload := model.URL{

	// }
	// TODO: return url

	return model.Appointment{}, 0, nil
}
