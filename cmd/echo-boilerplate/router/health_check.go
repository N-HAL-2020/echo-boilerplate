package router

import (
	healthcheck "github.com/N-HAL-2020/echo-boilerplate/cmd/echo-boilerplate/handler/health-check"
	"github.com/labstack/echo/v4"
)

type HealthCheckRouter struct {
	Handler healthcheck.Handler
}

func NewHealthCheckRouter(handler healthcheck.Handler) *HealthCheckRouter {
	return &HealthCheckRouter{Handler: handler}
}

func (r *HealthCheckRouter) Register(e *echo.Echo) {
	e.GET("/health-check", r.Handler.HealthCheck)
}
