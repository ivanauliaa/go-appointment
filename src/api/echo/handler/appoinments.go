package handler

import (
	"log"
	"net/http"

	"github.com/ivanauliaa/go-appoinment/src/domain"
	"github.com/ivanauliaa/go-appoinment/src/model"
	"github.com/ivanauliaa/go-appoinment/src/utils"
	"github.com/labstack/echo/v4"
)

type appointmentsHandler struct {
	service domain.AppointmentsService
}

func NewAppointmentsHandler(s domain.AppointmentsService) domain.AppointmentsHandler {
	newHandler := appointmentsHandler{
		service: s,
	}

	return &newHandler
}

func (h *appointmentsHandler) PostAppointmentHandler(c echo.Context) error {
	appointmentPayload := model.Appointment{}
	if err := c.Bind(&appointmentPayload); err != nil {
		return c.JSON(utils.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	eventPayload := model.Event{}
	if err := c.Bind(&eventPayload); err != nil {
		return c.JSON(utils.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	datePayload := model.Date{}
	if err := c.Bind(&datePayload); err != nil {
		return c.JSON(utils.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	timePayload := model.Time{}
	if err := c.Bind(&timePayload); err != nil {
		return c.JSON(utils.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	requestHeader := model.RequestHeader{
		Authorization: c.Request().Header["Authorization"][0],
	}

	appointment, code, err := h.service.AddAppointment(
		appointmentPayload,
		eventPayload,
		datePayload,
		timePayload,
		requestHeader,
	)
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(utils.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(utils.ServerErrorResponse())
	}

	return c.JSON(utils.SuccessResponseWithData(appointment))
}
