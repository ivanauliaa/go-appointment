package routes

import (
	"github.com/ivanauliaa/go-appoinment/src/api/echo/handler"
	"github.com/ivanauliaa/go-appoinment/src/database"
	"github.com/ivanauliaa/go-appoinment/src/middleware/auth"
	repository "github.com/ivanauliaa/go-appoinment/src/repository/postgres"
	"github.com/ivanauliaa/go-appoinment/src/service"
	"github.com/labstack/echo/v4"
)

func URLsRoutes(e *echo.Echo) {
	db := database.Connect()
	urlRepository := repository.NewURLsRepository(db)
	appointmentRepository := repository.NewAppointmentsRepository(db)

	service := service.NewURLsService(urlRepository, appointmentRepository)
	handler := handler.NewURLsHandler(service)

	e.POST("/appointments/:appointmentID/url", handler.PostURLHandler, auth.JWTMiddleware())
	e.GET("/appointments/:appointmentID/url", handler.GetURLHandler, auth.JWTMiddleware())
}
