package custom_middleware

import (
	"github.com/amirhosseinmoayedi/Project-template/internall/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LoggerMiddleware() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogRemoteIP:  true,
		LogHost:      true,
		LogMethod:    true,
		LogURI:       true,
		LogUserAgent: true,
		LogStatus:    true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			log.Logger.WithFields(map[string]interface{}{
				"URI":        values.URI,
				"status":     values.Status,
				"remote_ip":  values.RemoteIP,
				"host":       values.Host,
				"method":     values.Method,
				"user_agent": values.UserAgent,
			}).Info("request")
			return nil
		},
	})
}
