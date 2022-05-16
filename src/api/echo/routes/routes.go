package routes

import "github.com/labstack/echo/v4"

func AssignRoutes(e *echo.Echo) {
	UsersRoutes(e)
	AuthenticationsRoutes(e)
	AppointmentsRoutes(e)
	DatesRoutes(e)
	TimesRoutes(e)
	URLsRoutes(e)
}
