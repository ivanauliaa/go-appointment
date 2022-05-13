package routes

import (
	"github.com/ivanauliaa/go-appoinment/src/api/echo/handler"
	"github.com/ivanauliaa/go-appoinment/src/database"
	repository "github.com/ivanauliaa/go-appoinment/src/repository/postgres"
	"github.com/ivanauliaa/go-appoinment/src/service"
	"github.com/labstack/echo/v4"
)

func AuthenticationsRoutes(e *echo.Echo) {
	db := database.Connect()
	authenticationsRepository := repository.NewAuthenticationsRepository(db)
	usersRepository := repository.NewUsersRepository(db)
	service := service.NewAuthenticationsService(authenticationsRepository, usersRepository)
	handler := handler.NewAuthenticationsHandler(service)

	e.POST("/login", handler.LoginHandler)
	e.POST("/register", handler.RegisterHandler)
	e.PUT("/authentications", handler.PutAuthenticationHandler)
	e.DELETE("/authentications", handler.DeleteAuthenticationHandler)
}
