package main

import (
	"fmt"
	"net/http"

	"github.com/N-HAL-2020/echo-boilerplate/cmd/echo-boilerplate/config"
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

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	address := fmt.Sprintf(":%s", config.Port)
	e.Logger.Fatal(e.Start(address))
}
