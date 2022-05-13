package auth

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:                  &Claims{},
		SigningKey:              []byte(os.Getenv("ACCESS_TOKEN_KEY")),
		ErrorHandlerWithContext: JWTErrorChecker,
	})
}
