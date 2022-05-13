package domain

import "github.com/ivanauliaa/go-appoinment/src/model"

type URLsHandler interface {
}

type URLsService interface {
}

type URLsRepository interface {
	AddURL(payload model.URL) (string, int, error)
}
