package handler

import (
	"log"
	"net/http"

	"github.com/ivanauliaa/go-appoinment/src/domain"
	"github.com/ivanauliaa/go-appoinment/src/model"
	"github.com/ivanauliaa/go-appoinment/src/utils"
	"github.com/labstack/echo/v4"
)

type urlsHandler struct {
	service domain.URLsService
}

func NewURLsHandler(s domain.URLsService) domain.URLsHandler {
	newHandler := urlsHandler{
		service: s,
	}

	return &newHandler
}

func (h *urlsHandler) PostURLHandler(c echo.Context) error {
	payload := model.URL{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(utils.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	requestHeader := model.RequestHeader{
		Authorization: c.Request().Header["Authorization"][0],
	}

	appointmentURL, code, err := h.service.AddURL(payload, requestHeader)
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(utils.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(utils.ServerErrorResponse())
	}

	return c.JSON(utils.SuccessResponseWithData(appointmentURL))
}
