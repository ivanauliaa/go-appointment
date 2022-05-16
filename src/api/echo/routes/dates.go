package routes

import (
	"github.com/ivanauliaa/go-appoinment/src/api/echo/handler"
	"github.com/ivanauliaa/go-appoinment/src/database"
	"github.com/ivanauliaa/go-appoinment/src/middleware/auth"
	repository "github.com/ivanauliaa/go-appoinment/src/repository/postgres"
	"github.com/ivanauliaa/go-appoinment/src/service"
	"github.com/labstack/echo/v4"
)

func DatesRoutes(e *echo.Echo) {
	db := database.Connect()
	dateRepository := repository.NewDatesRepository(db)
	appointmentRepository := repository.NewAppointmentsRepository(db)

	service := service.NewDatesService(dateRepository, appointmentRepository)
	handler := handler.NewDatesHandler(service)

	e.POST("/appointments/:appointmentID/dates", handler.PostDateHandler, auth.JWTMiddleware())
}
