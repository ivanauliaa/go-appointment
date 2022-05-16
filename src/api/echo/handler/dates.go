package handler

import (
	"log"
	"net/http"

	"github.com/ivanauliaa/go-appoinment/src/domain"
	"github.com/ivanauliaa/go-appoinment/src/model"
	"github.com/ivanauliaa/go-appoinment/src/utils"
	"github.com/labstack/echo/v4"
)

type datesHandler struct {
	service domain.DatesService
}

func NewDatesHandler(s domain.DatesService) domain.DatesHandler {
	newHandler := datesHandler{
		service: s,
	}

	return &newHandler
}

func (h *datesHandler) PostDateHandler(c echo.Context) error {
	payload := model.Date{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(utils.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	requestHeader := model.RequestHeader{
		Authorization: c.Request().Header["Authorization"][0],
	}

	dateID, code, err := h.service.AddDate(payload, requestHeader)
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(utils.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(utils.ServerErrorResponse())
	}

	return c.JSON(utils.SuccessResponseWithData(dateID))
}
