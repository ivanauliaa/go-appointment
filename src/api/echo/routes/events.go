package routes

import (
	"net/http"

	"github.com/ivanauliaa/go-appoinment/src/api/echo/handler"
	"github.com/ivanauliaa/go-appoinment/src/database"
	"github.com/ivanauliaa/go-appoinment/src/middleware/auth"
	api "github.com/ivanauliaa/go-appoinment/src/repository/api"
	repository "github.com/ivanauliaa/go-appoinment/src/repository/postgres"
	"github.com/ivanauliaa/go-appoinment/src/service"
	"github.com/labstack/echo/v4"
)

func EventsRoutes(e *echo.Echo) {
	db := database.Connect()
	appointmentsRepository := repository.NewAppointmentsRepository(db)
	datesRepository := repository.NewDatesRepository(db)
	timesRepository := repository.NewTimesRepository(db)
	eventsRepository := repository.NewEventsRepository(db)
	usersRepository := repository.NewUsersRepository(db)

	client := &http.Client{}
	eventsAPI := api.NewEventsAPI(client)

	service := service.NewEventsService(
		eventsRepository,
		appointmentsRepository,
		datesRepository,
		timesRepository,
		usersRepository,
		eventsAPI,
	)
	handler := handler.NewEventsHandler(service)

	e.GET("/appointments/:appointmentID/event", handler.GetEventHandler, auth.JWTMiddleware())
	e.POST("/appointments/:appointmentID/event", handler.PostSendEventHandler, auth.JWTMiddleware())
}
