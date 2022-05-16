package handler

import (
	"log"
	"net/http"

	"github.com/ivanauliaa/go-appoinment/src/domain"
	"github.com/ivanauliaa/go-appoinment/src/model"
	"github.com/ivanauliaa/go-appoinment/src/utils"

	"github.com/labstack/echo/v4"
)

type authenticationsHandler struct {
	service domain.AuthenticationsService
}

func NewAuthenticationsHandler(s domain.AuthenticationsService) domain.AuthenticationsHandler {
	newHandler := authenticationsHandler{
		service: s,
	}

	return &newHandler
}

func (h *authenticationsHandler) LoginHandler(c echo.Context) error {
	payload := model.User{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(utils.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	token, code, err := h.service.Login(payload)
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(utils.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(utils.ServerErrorResponse())
	}

	return c.JSON(utils.SuccessResponseWithData(token))
}

func (h *authenticationsHandler) RegisterHandler(c echo.Context) error {
	payload := model.User{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(utils.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(utils.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	code, err := h.service.Register(payload)
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(utils.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(utils.ServerErrorResponse())
	}

	return c.JSON(utils.SuccessResponse())
}

func (h *authenticationsHandler) PutAuthenticationHandler(c echo.Context) error {
	payload := model.RefreshTokenPayload{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(utils.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	accessToken, code, err := h.service.UpdateAccessToken(payload)
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(utils.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(utils.ServerErrorResponse())
	}

	return c.JSON(utils.SuccessResponseWithData(accessToken))
}

func (h *authenticationsHandler) DeleteAuthenticationHandler(c echo.Context) error {
	payload := model.RefreshTokenPayload{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(utils.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	code, err := h.service.DeleteRefreshToken(payload)
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(utils.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(utils.ServerErrorResponse())
	}

	return c.JSON(utils.SuccessResponse())
}
