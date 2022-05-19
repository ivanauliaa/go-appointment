package domain

import (
	"github.com/ivanauliaa/go-appoinment/src/model"
	"github.com/labstack/echo/v4"
)

type TimesHandler interface {
	PostTimeHandler(c echo.Context) error
}

type TimesService interface {
	AddTime(payload model.PostTimePayload, requestHeader model.RequestHeader) (model.PostTimeResponse, int, error)
}

type TimesRepository interface {
	AddTime(payload model.Time) (model.PostTimeResponse, int, error)
	VerifyTime(timeID uint) (int, error)
	VerifyTimeDateID(timeID uint, dateID uint) (int, error)
	GetTime(timeID uint) (model.Time, int, error)
	GetTimes(dateID uint) ([]model.Time, int, error)
}
