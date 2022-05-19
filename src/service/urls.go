package service

import (
	"net/http"
	"strings"

	"github.com/ivanauliaa/go-appoinment/src/domain"
	"github.com/ivanauliaa/go-appoinment/src/middleware/auth"
	"github.com/ivanauliaa/go-appoinment/src/model"
)

type urlsService struct {
	urlsRepository         domain.URLsRepository
	appointmentsRepository domain.AppointmentsRepository
}

func NewURLsService(
	ur domain.URLsRepository,
	ar domain.AppointmentsRepository,
) domain.URLsService {
	newService := urlsService{
		urlsRepository:         ur,
		appointmentsRepository: ar,
	}

	return &newService
}

func (s *urlsService) AddURL(
	payload model.URL,
	requestHeader model.RequestHeader,
) (model.PostURLResponse, int, error) {
	if code, err := s.appointmentsRepository.VerifyAppointment(payload.AppointmentID); err != nil {
		return model.PostURLResponse{}, code, err
	}

	if code, err := s.urlsRepository.VerifyNewURL(payload.AppointmentID); err != nil {
		return model.PostURLResponse{}, code, err
	}

	accessToken := strings.Split(requestHeader.Authorization, " ")[1]

	credentialId, err := auth.GetAuthCredential(accessToken)
	if err != nil {
		return model.PostURLResponse{}, http.StatusBadRequest, err
	}

	if code, err := s.appointmentsRepository.VerifyAppointmentOwner(payload.AppointmentID, credentialId); err != nil {
		return model.PostURLResponse{}, code, err
	}

	return s.urlsRepository.AddURL(payload)
}

func (s *urlsService) GetURL(
	payload model.URL,
	requestHeader model.RequestHeader,
) (model.URL, int, error) {
	if code, err := s.appointmentsRepository.VerifyAppointment(payload.AppointmentID); err != nil {
		return model.URL{}, code, err
	}

	accessToken := strings.Split(requestHeader.Authorization, " ")[1]

	credentialId, err := auth.GetAuthCredential(accessToken)
	if err != nil {
		return model.URL{}, http.StatusBadRequest, err
	}

	if code, err := s.appointmentsRepository.VerifyAppointmentOwner(payload.AppointmentID, credentialId); err != nil {
		return model.URL{}, code, err
	}

	return s.urlsRepository.GetURL(payload.AppointmentID)
}
