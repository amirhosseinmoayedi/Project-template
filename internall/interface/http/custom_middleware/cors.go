package custom_middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CorsMiddleware() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "x-CSRF-Token"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Link"},
		MaxAge:           300,
	})
}
