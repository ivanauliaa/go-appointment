package routes

import "github.com/labstack/echo/v4"

func AssignRoutes(e *echo.Echo) {
	// TODO: assign e dengan routes" dari tiap" entity
	UsersRoutes(e)
	AuthenticationsRoutes(e)
	AppointmentsRoutes(e)
	DatesRoutes(e)
	TimesRoutes(e)
	URLsRoutes(e)
}
