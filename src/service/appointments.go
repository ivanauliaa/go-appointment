package service

import (
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
	usersRepository        domain.UsersRepository
}

func NewAppointmentsService(
	ar domain.AppointmentsRepository,
	dr domain.DatesRepository,
	tr domain.TimesRepository,
	er domain.EventsRepository,
	ur domain.UsersRepository,
) domain.AppointmentsService {
	newService := appointmentsService{
		appointmentsRepository: ar,
		datesRepository:        dr,
		timesRepository:        tr,
		eventsRepository:       er,
		usersRepository:        ur,
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

func (s *appointmentsService) GetAppointments(requestHeader model.RequestHeader) ([]model.Appointment, int, error) {
	accessToken := strings.Split(requestHeader.Authorization, " ")[1]

	credentialId, err := auth.GetAuthCredential(accessToken)
	if err != nil {
		return []model.Appointment{}, http.StatusBadRequest, err
	}

	return s.appointmentsRepository.GetAppointments(credentialId)
}

func (s *appointmentsService) GetAppointment(
	payload model.GetAppointmentPayload,
	requestHeader model.RequestHeader,
) (model.AppointmentWithRelation, int, error) {
	if code, err := s.appointmentsRepository.VerifyAppointment(payload.AppointmentID); err != nil {
		return model.AppointmentWithRelation{}, code, err
	}

	accessToken := strings.Split(requestHeader.Authorization, " ")[1]

	credentialId, err := auth.GetAuthCredential(accessToken)
	if err != nil {
		return model.AppointmentWithRelation{}, http.StatusBadRequest, err
	}

	if code, err := s.appointmentsRepository.VerifyAppointmentOwner(payload.AppointmentID, credentialId); err != nil {
		return model.AppointmentWithRelation{}, code, err
	}

	appointment, code, err := s.appointmentsRepository.GetAppointment(payload.AppointmentID)
	if err != nil {
		return model.AppointmentWithRelation{}, code, err
	}

	dates, code, err := s.datesRepository.GetDates(appointment.ID)
	if err != nil {
		return model.AppointmentWithRelation{}, code, err
	}

	appointmentDates := []model.AppointmentDate{}

	for j := 0; j < len(dates); j++ {
		times, code, err := s.timesRepository.GetTimes(dates[j].ID)
		if err != nil {
			return model.AppointmentWithRelation{}, code, err
		}

		appointmentDate := model.AppointmentDate{
			Date:  dates[j],
			Times: times,
		}

		appointmentDates = append(appointmentDates, appointmentDate)
	}

	appointmentWithRelation := model.AppointmentWithRelation{
		Appointment: appointment,
		Dates:       appointmentDates,
	}

	return appointmentWithRelation, http.StatusOK, nil
}

func (s *appointmentsService) UpdateAppointment(
	payload model.PutAppointmentPayload,
	requestHeader model.RequestHeader,
) (int, error) {
	if code, err := s.appointmentsRepository.VerifyAppointment(payload.AppointmentID); err != nil {
		return code, err
	}

	accessToken := strings.Split(requestHeader.Authorization, " ")[1]

	credentialId, err := auth.GetAuthCredential(accessToken)
	if err != nil {
		return http.StatusBadRequest, err
	}

	if code, err := s.appointmentsRepository.VerifyAppointmentOwner(payload.AppointmentID, credentialId); err != nil {
		return code, err
	}

	updatedAppointment := model.Appointment{
		UserID: credentialId,
		Name:   payload.Name,
		Room:   payload.Room,
	}
	return s.appointmentsRepository.UpdateAppointment(updatedAppointment, payload.AppointmentID)
}

func (s *appointmentsService) DeleteAppointment(
	payload model.DeleteAppointmentPayload,
	requestHeader model.RequestHeader,
) (int, error) {
	if code, err := s.appointmentsRepository.VerifyAppointment(payload.AppointmentID); err != nil {
		return code, err
	}

	accessToken := strings.Split(requestHeader.Authorization, " ")[1]

	credentialId, err := auth.GetAuthCredential(accessToken)
	if err != nil {
		return http.StatusBadRequest, err
	}

	if code, err := s.appointmentsRepository.VerifyAppointmentOwner(payload.AppointmentID, credentialId); err != nil {
		return code, err
	}

	return s.appointmentsRepository.DeleteAppointment(payload.AppointmentID)
}
