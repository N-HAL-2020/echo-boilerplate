package main

import (
	"fmt"

	"github.com/N-HAL-2020/echo-boilerplate/cmd/echo-boilerplate/config"
	"github.com/N-HAL-2020/echo-boilerplate/cmd/echo-boilerplate/middleware"
	"github.com/N-HAL-2020/echo-boilerplate/cmd/echo-boilerplate/router"
	"github.com/N-HAL-2020/echo-boilerplate/cmd/echo-boilerplate/validator"
	"github.com/labstack/echo/v4"
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

	validator.ApplyValidator(e)
	middleware.ApplyMiddleware(e)
	router.RegisterRouter(e)

	address := fmt.Sprintf(":%s", config.Port)
	e.Logger.Fatal(e.Start(address))
}
