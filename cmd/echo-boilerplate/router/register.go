package router

import (
	healthcheck "github.com/N-HAL-2020/echo-boilerplate/cmd/echo-boilerplate/handler/health-check"
	"github.com/labstack/echo/v4"
)

func RegisterRouter(e *echo.Echo) {
	healthHandler := healthcheck.NewHandler()
	healthRouter := NewHealthCheckRouter(healthHandler)

	healthRouter.Register(e)
}
