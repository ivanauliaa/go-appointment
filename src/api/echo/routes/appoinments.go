package routes

import (
	"github.com/ivanauliaa/go-appoinment/src/api/echo/handler"
	"github.com/ivanauliaa/go-appoinment/src/database"
	"github.com/ivanauliaa/go-appoinment/src/middleware/auth"
	repository "github.com/ivanauliaa/go-appoinment/src/repository/postgres"
	"github.com/ivanauliaa/go-appoinment/src/service"

	"github.com/labstack/echo/v4"
)

func AppointmentsRoutes(e *echo.Echo) {
	db := database.Connect()
	appointmentsRepository := repository.NewAppointmentsRepository(db)
	datesRepository := repository.NewDatesRepository(db)
	timesRepository := repository.NewTimesRepository(db)
	eventsRepository := repository.NewEventsRepository(db)

	service := service.NewAppointmentsService(
		appointmentsRepository,
		datesRepository,
		timesRepository,
		eventsRepository,
	)
	handler := handler.NewAppointmentsHandler(service)

	e.POST("/appointments", handler.PostAppointmentHandler, auth.JWTMiddleware())
	e.POST("/appointments/:appointmentID/confirm", handler.PostAppointmentConfirmHandler, auth.JWTMiddleware())
}
