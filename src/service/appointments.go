package service

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ivanauliaa/go-appoinment/src/domain"
	"github.com/ivanauliaa/go-appoinment/src/middleware/auth"
	"github.com/ivanauliaa/go-appoinment/src/model"
)

type appointmentsService struct {
	appointmentsRepository domain.AppointmentsRepository
	datesRepository        domain.DatesRepository
	timesRepository        domain.TimesRepository
	eventsRepository       domain.EventsRepository
}

func NewAppointmentsService(
	ar domain.AppointmentsRepository,
	dr domain.DatesRepository,
	tr domain.TimesRepository,
	er domain.EventsRepository,
) domain.AppointmentsService {
	newService := appointmentsService{
		appointmentsRepository: ar,
		datesRepository:        dr,
		timesRepository:        tr,
		eventsRepository:       er,
	}

	return &newService
}

func (s *appointmentsService) AddAppointment(
	payload model.PostAppointmentPayload,
	requestHeader model.RequestHeader,
) (model.PostAppointmentResponse, int, error) {
	accessToken := strings.Split(requestHeader.Authorization, " ")[1]

	credentialId, err := auth.GetAuthCredential(accessToken)
	if err != nil {
		return model.PostAppointmentResponse{}, http.StatusBadRequest, err
	}

	appointment := model.Appointment{
		UserID: credentialId,
		Name:   payload.Name,
		Room:   payload.Room,
	}
	result, code, err := s.appointmentsRepository.AddAppointment(appointment)
	if err != nil {
		return model.PostAppointmentResponse{}, code, err
	}
	fmt.Println(result)

	return result, code, err
}

func (s *appointmentsService) ConfirmAppointment(
	payload model.PostAppointmentConfirmPayload,
	requestHeader model.RequestHeader,
) (int, error) {
	if code, err := s.eventsRepository.VerifyNewEvent(payload.AppointmentID); err != nil {
		return code, err
	}

	if code, err := s.appointmentsRepository.VerifyAppointment(payload.AppointmentID); err != nil {
		return code, err
	}

	accessToken := strings.Split(requestHeader.Authorization, " ")[1]

	credentialId, err := auth.GetAuthCredential(accessToken)
	if err != nil {
		return http.StatusBadRequest, err
	}

	if code, err := s.appointmentsRepository.VerifyAppointmentGuest(payload.AppointmentID, credentialId); err != nil {
		return code, err
	}

	if code, err := s.datesRepository.VerifyDate(payload.DateID); err != nil {
		return code, err
	}

	if code, err := s.datesRepository.VerifyDateAppointmentID(payload.DateID, payload.AppointmentID); err != nil {
		return code, err
	}

	if code, err := s.timesRepository.VerifyTime(payload.TimeID); err != nil {
		return code, err
	}

	if code, err := s.timesRepository.VerifyTimeDateID(payload.TimeID, payload.DateID); err != nil {
		return code, err
	}

	appointment, code, err := s.appointmentsRepository.GetAppointment(payload.AppointmentID)
	if err != nil {
		return code, err
	}

	event := model.Event{
		AppointmentID: payload.AppointmentID,
		HostID:        appointment.UserID,
		GuestID:       credentialId,
		DateID:        payload.DateID,
		TimeID:        payload.TimeID,
		Status:        "CONFIRMED",
	}

	return s.eventsRepository.AddEvent(event)
}
