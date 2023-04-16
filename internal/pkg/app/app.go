package app

import (
	"fmt"
	"log"

	"github.com/YakovAkk/TestMiddleware/internal/app/endpoint"
	"github.com/YakovAkk/TestMiddleware/internal/app/middleware"
	"github.com/YakovAkk/TestMiddleware/internal/app/service"
	"github.com/labstack/echo/v4"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.s = service.New()
	a.e = endpoint.New(a.s)

	a.echo = echo.New()

	a.echo.Use(middleware.RoleCheck)

	a.echo.GET("/status", a.e.Status)

	return a, nil
}

func (app *App) Run() error {
	fmt.Println("Server is running...")

	err := app.echo.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
