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
	payload := model.PostAppointmentPayload{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(utils.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(utils.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	requestHeader := model.RequestHeader{
		Authorization: c.Request().Header["Authorization"][0],
	}

	appointmentID, code, err := h.service.AddAppointment(
		payload,
		requestHeader,
	)
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(utils.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(utils.ServerErrorResponse())
	}

	return c.JSON(utils.SuccessResponseWithData(appointmentID))
}

func (h *appointmentsHandler) PostAppointmentConfirmHandler(c echo.Context) error {
	payload := model.PostAppointmentConfirmPayload{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(utils.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(utils.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	requestHeader := model.RequestHeader{
		Authorization: c.Request().Header["Authorization"][0],
	}

	code, err := h.service.ConfirmAppointment(
		payload,
		requestHeader,
	)
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(utils.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(utils.ServerErrorResponse())
	}

	return c.JSON(utils.SuccessResponse())
}

func (h *appointmentsHandler) GetAppointmentsHandler(c echo.Context) error {
	requestHeader := model.RequestHeader{
		Authorization: c.Request().Header["Authorization"][0],
	}

	appointments, code, err := h.service.GetAppointments(
		requestHeader,
	)
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(utils.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(utils.ServerErrorResponse())
	}

	return c.JSON(utils.SuccessResponseWithData(appointments))
}

func (h *appointmentsHandler) GetAppointmentHandler(c echo.Context) error {
	payload := model.GetAppointmentPayload{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(utils.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	requestHeader := model.RequestHeader{
		Authorization: c.Request().Header["Authorization"][0],
	}

	appointmentWithRelation, code, err := h.service.GetAppointment(
		payload,
		requestHeader,
	)

	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(utils.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(utils.ServerErrorResponse())
	}

	return c.JSON(utils.SuccessResponseWithData(appointmentWithRelation))
}
