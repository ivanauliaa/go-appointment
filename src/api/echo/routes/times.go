package routes

import (
	"github.com/ivanauliaa/go-appoinment/src/api/echo/handler"
	"github.com/ivanauliaa/go-appoinment/src/database"
	"github.com/ivanauliaa/go-appoinment/src/middleware/auth"
	repository "github.com/ivanauliaa/go-appoinment/src/repository/postgres"
	"github.com/ivanauliaa/go-appoinment/src/service"
	"github.com/labstack/echo/v4"
)

func TimesRoutes(e *echo.Echo) {
	db := database.Connect()
	timeRepository := repository.NewTimesRepository(db)
	appointmentRepository := repository.NewAppointmentsRepository(db)
	dateRepository := repository.NewDatesRepository(db)

	service := service.NewTimesService(timeRepository, appointmentRepository, dateRepository)
	handler := handler.NewTimesHandler(service)

	e.POST("/appointments/:appointmentID/dates/:dateID/times", handler.PostTimeHandler, auth.JWTMiddleware())
}
