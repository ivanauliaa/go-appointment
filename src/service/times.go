package service

import (
	"net/http"
	"strings"

	"github.com/ivanauliaa/go-appoinment/src/domain"
	"github.com/ivanauliaa/go-appoinment/src/middleware/auth"
	"github.com/ivanauliaa/go-appoinment/src/model"
)

type TimesService struct {
	timesRepository        domain.TimesRepository
	appointsmentRepository domain.AppointmentsRepository
	datesRepository        domain.DatesRepository
}

func NewTimesService(
	tr domain.TimesRepository,
	ar domain.AppointmentsRepository,
	dr domain.DatesRepository,
) domain.TimesService {
	newService := TimesService{
		timesRepository:        tr,
		appointsmentRepository: ar,
		datesRepository:        dr,
	}

	return &newService
}

func (s *TimesService) AddTime(
	payload model.PostTimePayload,
	requestHeader model.RequestHeader,
) (model.PostTimeResponse, int, error) {
	if code, err := s.appointsmentRepository.VerifyAppointment(payload.AppointmentID); err != nil {
		return model.PostTimeResponse{}, code, err
	}

	accessToken := strings.Split(requestHeader.Authorization, " ")[1]

	credentialId, err := auth.GetAuthCredential(accessToken)
	if err != nil {
		return model.PostTimeResponse{}, http.StatusBadRequest, err
	}

	if code, err := s.appointsmentRepository.VerifyAppointmentOwner(payload.AppointmentID, credentialId); err != nil {
		return model.PostTimeResponse{}, code, err
	}

	if code, err := s.datesRepository.VerifyDate(payload.DateID); err != nil {
		return model.PostTimeResponse{}, code, err
	}

	time := model.Time{
		DateID: payload.DateID,
		Start:  payload.Start,
		End:    payload.End,
	}

	return s.timesRepository.AddTime(time)
}
