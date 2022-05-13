package domain

import (
	"github.com/ivanauliaa/go-appoinment/src/model"
	"github.com/labstack/echo/v4"
)

type AuthenticationsHandler interface {
	LoginHandler(c echo.Context) error
	RegisterHandler(c echo.Context) error
	PutAuthenticationHandler(c echo.Context) error
	DeleteAuthenticationHandler(c echo.Context) error
}

type AuthenticationsService interface {
	Login(user model.User) (model.LoginResponse, int, error)
	Register(user model.User) (int, error)
	UpdateAccessToken(token model.RefreshTokenPayload) (model.AccessTokenResponse, int, error)
	DeleteRefreshToken(token model.RefreshTokenPayload) (int, error)
}

type AuthenticationsRepository interface {
	AddRefreshToken(token model.Authentication) (int, error)
	VerifyRefreshToken(token model.RefreshTokenPayload) (int, error)
	DeleteRefreshToken(token model.RefreshTokenPayload) (int, error)
}
