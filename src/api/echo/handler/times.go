package handler

import (
	"log"
	"net/http"

	"github.com/ivanauliaa/go-appoinment/src/domain"
	"github.com/ivanauliaa/go-appoinment/src/model"
	"github.com/ivanauliaa/go-appoinment/src/utils"
	"github.com/labstack/echo/v4"
)

type timesHandler struct {
	service domain.TimesService
}

func NewTimesHandler(s domain.TimesService) domain.TimesHandler {
	newHandler := timesHandler{
		service: s,
	}

	return &newHandler
}

func (h *timesHandler) PostTimeHandler(c echo.Context) error {
	payload := model.PostTimePayload{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(utils.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	requestHeader := model.RequestHeader{
		Authorization: c.Request().Header["Authorization"][0],
	}

	timeID, code, err := h.service.AddTime(payload, requestHeader)
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(utils.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(utils.ServerErrorResponse())
	}

	return c.JSON(utils.SuccessResponseWithData(timeID))
}
