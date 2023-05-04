package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HeartBeat(c echo.Context) error {
	return c.JSON(http.StatusOK, "Beating")
}
