package domain

import "github.com/ivanauliaa/go-appoinment/src/model"

type DatesHandler interface {
}

type DatesService interface {
}

type DatesRepository interface {
	AddDate(payload model.Date) (int, error)
}
