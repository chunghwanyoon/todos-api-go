package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HealthCheck() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "OK")
	}
}
