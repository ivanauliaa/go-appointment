package handler

import (
	"github.com/ivanauliaa/go-appoinment/src/domain"

	"github.com/labstack/echo/v4"
)

type usersHandler struct {
	service domain.UsersService
}

func NewUsersHandler(s domain.UsersService) domain.UsersHandler {
	newHandler := usersHandler{
		service: s,
	}

	return &newHandler
}

func (h *usersHandler) HelloHandler(c echo.Context) error {
	hello := h.service.Hello()
	return c.JSON(200, map[string]interface{}{
		"hello": hello,
	})
}

func (h *usersHandler) HelloWorldHandler(c echo.Context) error {
	return c.JSON(200, map[string]interface{}{
		"today message": "hello world",
	})
}
