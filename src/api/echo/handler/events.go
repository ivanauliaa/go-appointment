package handler

import (
	"log"
	"net/http"

	"github.com/ivanauliaa/go-appoinment/src/domain"
	"github.com/ivanauliaa/go-appoinment/src/model"
	"github.com/ivanauliaa/go-appoinment/src/utils"
	"github.com/labstack/echo/v4"
)

type eventsHandler struct {
	service domain.EventsService
}

func NewEventsHandler(s domain.EventsService) domain.EventsHandler {
	newHandler := eventsHandler{
		service: s,
	}

	return &newHandler
}

func (h *eventsHandler) GetEventHandler(c echo.Context) error {
	payload := model.GetEventPayload{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(utils.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(utils.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	requestHeader := model.RequestHeader{
		Authorization: c.Request().Header["Authorization"][0],
	}

	appointmentEvent, code, err := h.service.GetEvent(
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

	return c.JSON(utils.SuccessResponseWithData(appointmentEvent))
}

func (h *eventsHandler) PostSendEventHandler(c echo.Context) error {
	payload := model.SendEventPayload{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(utils.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if code, err := h.service.SendEvent(payload); err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(utils.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(utils.ServerErrorResponse())
	}

	return c.JSON(utils.SuccessResponse())
}
