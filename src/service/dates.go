package service

import (
	"net/http"
	"strings"

	"github.com/ivanauliaa/go-appoinment/src/domain"
	"github.com/ivanauliaa/go-appoinment/src/middleware/auth"
	"github.com/ivanauliaa/go-appoinment/src/model"
)

type datesService struct {
	datesRepository        domain.DatesRepository
	appointmentsRepository domain.AppointmentsRepository
}

func NewDatesService(
	dr domain.DatesRepository,
	ar domain.AppointmentsRepository,
) domain.DatesService {
	newService := datesService{
		datesRepository:        dr,
		appointmentsRepository: ar,
	}

	return &newService
}

func (s *datesService) AddDate(
	payload model.Date,
	requestHeader model.RequestHeader,
) (model.PostDateResponse, int, error) {
	if code, err := s.appointmentsRepository.VerifyAppointment(payload.AppointmentID); err != nil {
		return model.PostDateResponse{}, code, err
	}

	accessToken := strings.Split(requestHeader.Authorization, " ")[1]

	credentialId, err := auth.GetAuthCredential(accessToken)
	if err != nil {
		return model.PostDateResponse{}, http.StatusBadRequest, err
	}

	if code, err := s.appointmentsRepository.VerifyAppointmentOwner(payload.AppointmentID, credentialId); err != nil {
		return model.PostDateResponse{}, code, err
	}

	return s.datesRepository.AddDate(payload)
}
