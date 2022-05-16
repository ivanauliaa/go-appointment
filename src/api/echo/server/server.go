package server

import (
	"github.com/ivanauliaa/go-appoinment/src/api/echo/routes"
	"github.com/labstack/echo/v4"
)

func CreateServer() *echo.Echo {
	e := echo.New()
	routes.AssignRoutes(e)

	return e
}
