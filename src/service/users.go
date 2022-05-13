package service

import (
	"github.com/ivanauliaa/go-appoinment/src/domain"
)

type userService struct {
	repository domain.UsersRepository
}

func NewUsersService(r domain.UsersRepository) domain.UsersService {
	newService := userService{
		repository: r,
	}

	return &newService
}

func (s *userService) Hello() string {
	str := s.repository.Hello()

	return str
}
