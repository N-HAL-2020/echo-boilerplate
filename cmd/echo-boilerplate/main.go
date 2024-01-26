package main

import (
	"fmt"

	"github.com/N-HAL-2020/echo-boilerplate/cmd/echo-boilerplate/config"
	healthcheck "github.com/N-HAL-2020/echo-boilerplate/cmd/echo-boilerplate/handler/health-check"
	"github.com/N-HAL-2020/echo-boilerplate/cmd/echo-boilerplate/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	run()
}

func run() {
	config, err := config.Load()
	if err != nil {
		panic(err)
	}

	e := echo.New()

	applyMiddleware(e)
	registerRouter(e)

	address := fmt.Sprintf(":%s", config.Port)
	e.Logger.Fatal(e.Start(address))
}

func applyMiddleware(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}

func registerRouter(e *echo.Echo) {
	healthHandler := healthcheck.NewHandler()
	healthRouter := router.NewHealthCheckRouter(healthHandler)
	healthRouter.Register(e)
}
