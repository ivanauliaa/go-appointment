package domain

import "github.com/ivanauliaa/go-appoinment/src/model"

type TimesHandler interface {
}

type TimesService interface {
}

type TimesRepository interface {
	AddTime(payload model.Time) (int, error)
}
