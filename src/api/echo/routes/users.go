package routes

import (
	"github.com/ivanauliaa/go-appoinment/src/api/echo/handler"
	"github.com/ivanauliaa/go-appoinment/src/database"
	"github.com/ivanauliaa/go-appoinment/src/middleware/auth"
	repository "github.com/ivanauliaa/go-appoinment/src/repository/postgres"
	"github.com/ivanauliaa/go-appoinment/src/service"

	"github.com/labstack/echo/v4"
)

func UsersRoutes(e *echo.Echo) {
	db := database.Connect()
	repository := repository.NewUsersRepository(db)
	service := service.NewUsersService(repository)
	handler := handler.NewUsersHandler(service)

	e.GET("/hello", handler.HelloHandler, auth.JWTMiddleware())
}
