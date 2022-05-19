package domain

import (
	"github.com/ivanauliaa/go-appoinment/src/model"
	"github.com/labstack/echo/v4"
)

type URLsHandler interface {
	PostURLHandler(c echo.Context) error
	GetURLHandler(c echo.Context) error
}

type URLsService interface {
	AddURL(payload model.URL, requestHeader model.RequestHeader) (model.PostURLResponse, int, error)
	GetURL(payload model.URL, requestHeader model.RequestHeader) (model.URL, int, error)
}

type URLsRepository interface {
	AddURL(payload model.URL) (model.PostURLResponse, int, error)
	VerifyNewURL(appointmentID uint) (int, error)
	GetURL(appointmentID uint) (model.URL, int, error)
}
