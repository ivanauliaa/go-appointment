package service

import (
	"net/http"
	"strings"

	"github.com/ivanauliaa/go-appoinment/src/domain"
	"github.com/ivanauliaa/go-appoinment/src/middleware/auth"
	"github.com/ivanauliaa/go-appoinment/src/model"
)

type eventService struct {
	eventsRepository       domain.EventsRepository
	appointmentsRepository domain.AppointmentsRepository
	datesRepository        domain.DatesRepository
	timesRepository        domain.TimesRepository
	usersRepository        domain.UsersRepository
	eventsAPI              domain.EventsAPI
}

func NewEventsService(
	er domain.EventsRepository,
	ar domain.AppointmentsRepository,
	dr domain.DatesRepository,
	tr domain.TimesRepository,
	ur domain.UsersRepository,
	ea domain.EventsAPI,
) domain.EventsService {
	newService := eventService{
		eventsRepository:       er,
		appointmentsRepository: ar,
		datesRepository:        dr,
		timesRepository:        tr,
		usersRepository:        ur,
		eventsAPI:              ea,
	}

	return &newService
}

func (s *eventService) GetEvent(
	payload model.GetEventPayload,
	requestHeader model.RequestHeader,
) (model.EventWithRelation, int, error) {
	if code, err := s.appointmentsRepository.VerifyAppointment(payload.AppointmentID); err != nil {
		return model.EventWithRelation{}, code, err
	}

	accessToken := strings.Split(requestHeader.Authorization, " ")[1]

	credentialId, err := auth.GetAuthCredential(accessToken)
	if err != nil {
		return model.EventWithRelation{}, http.StatusBadRequest, err
	}

	if code, err := s.eventsRepository.VerifyEventParticipant(payload.AppointmentID, credentialId); err != nil {
		return model.EventWithRelation{}, code, err
	}

	event, code, err := s.eventsRepository.GetEvent(payload.AppointmentID)
	if err != nil {
		return model.EventWithRelation{}, code, err
	}

	appointment, code, err := s.appointmentsRepository.GetAppointment(payload.AppointmentID)
	if err != nil {
		return model.EventWithRelation{}, code, err
	}

	date, code, err := s.datesRepository.GetDate(event.DateID)
	if err != nil {
		return model.EventWithRelation{}, code, err
	}

	time, code, err := s.timesRepository.GetTime(event.TimeID)
	if err != nil {
		return model.EventWithRelation{}, code, err
	}

	guest, code, err := s.usersRepository.GetUser(event.GuestID)
	if err != nil {
		return model.EventWithRelation{}, code, err
	}

	host, code, err := s.usersRepository.GetUser(event.HostID)
	if err != nil {
		return model.EventWithRelation{}, code, err
	}

	eventWithRelation := model.EventWithRelation{
		Event: event,
		SendEvent: model.SendEvent{
			Name:       appointment.Name,
			Room:       appointment.Room,
			Date:       date.Date,
			Start:      time.Start,
			End:        time.End,
			GuestEmail: guest.Email,
			HostEmail:  host.Email,
		},
	}

	return eventWithRelation, http.StatusOK, nil
}

func (s *eventService) SendEvent(payload model.SendEventPayload) (int, error) {
	if code, err := s.appointmentsRepository.VerifyAppointment(payload.AppointmentID); err != nil {
		return code, err
	}

	event, code, err := s.eventsRepository.GetEvent(payload.AppointmentID)
	if err != nil {
		return code, err
	}

	appointment, code, err := s.appointmentsRepository.GetAppointment(payload.AppointmentID)
	if err != nil {
		return code, err
	}

	host, code, err := s.usersRepository.GetUser(event.HostID)
	if err != nil {
		return code, err
	}

	guest, code, err := s.usersRepository.GetUser(event.GuestID)
	if err != nil {
		return code, err
	}

	date, code, err := s.datesRepository.GetDate(event.DateID)
	if err != nil {
		return code, err
	}

	time, code, err := s.timesRepository.GetTime(event.TimeID)
	if err != nil {
		return code, err
	}

	sendEvent := model.SendEvent{
		Name:       appointment.Name,
		GuestEmail: guest.Email,
		HostEmail:  host.Email,
		Date:       date.Date,
		Start:      time.Start,
		End:        time.End,
		Room:       appointment.Room,
	}

	return s.eventsAPI.SendEvent(sendEvent)
}
