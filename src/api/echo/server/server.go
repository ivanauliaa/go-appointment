package server

import (
	"github.com/ivanauliaa/go-appoinment/src/api/echo/routes"
	"github.com/ivanauliaa/go-appoinment/src/api/echo/validator"
	"github.com/labstack/echo/v4"
)

func CreateServer() *echo.Echo {
	e := echo.New()
	routes.AssignRoutes(e)
	validator.AssignValidator(e)

	return e
}
