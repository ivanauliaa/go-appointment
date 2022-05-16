package service

import (
	"net/http"

	"github.com/ivanauliaa/go-appoinment/src/domain"
	"github.com/ivanauliaa/go-appoinment/src/middleware/auth"
	"github.com/ivanauliaa/go-appoinment/src/model"
)

type authenticationsService struct {
	authenticationsRepository domain.AuthenticationsRepository
	usersRepository           domain.UsersRepository
}

func NewAuthenticationsService(ar domain.AuthenticationsRepository, ur domain.UsersRepository) domain.AuthenticationsService {
	newService := authenticationsService{
		authenticationsRepository: ar,
		usersRepository:           ur,
	}

	return &newService
}

func (s *authenticationsService) Login(user model.User) (model.LoginResponse, int, error) {
	id, code, err := s.usersRepository.VerifyUserCredential(user.Email, user.Password)
	if err != nil {
		return model.LoginResponse{}, code, err
	}

	accessToken := auth.GenerateAccessToken(id)
	refreshToken := auth.GenerateRefreshToken(id)

	if code, err := s.authenticationsRepository.AddRefreshToken(model.Authentication{
		Token: refreshToken,
	}); err != nil {
		return model.LoginResponse{}, code, err
	}

	return model.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, code, err
}

func (s *authenticationsService) Register(user model.User) (int, error) {
	if code, err := s.usersRepository.VerifyNewUserEmail(user.Email); err != nil {
		return code, err
	}

	return s.usersRepository.AddUser(user)
}

func (s *authenticationsService) UpdateAccessToken(token model.RefreshTokenPayload) (model.AccessTokenResponse, int, error) {
	code, err := s.authenticationsRepository.VerifyRefreshToken(token)
	if err != nil {
		return model.AccessTokenResponse{}, code, err
	}

	id, err := auth.VerifyRefreshToken(token.RefreshToken)
	if err != nil {
		return model.AccessTokenResponse{}, http.StatusBadRequest, err
	}

	accessToken := auth.GenerateAccessToken(id)
	return model.AccessTokenResponse{
		AccessToken: accessToken,
	}, http.StatusOK, nil
}

func (s *authenticationsService) DeleteRefreshToken(token model.RefreshTokenPayload) (int, error) {
	if code, err := s.authenticationsRepository.VerifyRefreshToken(token); err != nil {
		return code, err
	}

	return s.authenticationsRepository.DeleteRefreshToken(token)
}
