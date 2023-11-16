package rest

import (
	"fmt"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/piovani/go_full/http/rest/controller"
	"github.com/piovani/go_full/infra/config"
)

type Rest struct {
	echo *echo.Echo
}

func NewRest() *Rest {
	return &Rest{
		echo: echo.New(),
	}
}

func (r *Rest) Execute() error {
	r.getRoutes()

	return r.start()
}

func (r *Rest) getRoutes() {
	healthController := controller.NewHealthController()

	r.echo.Use(middleware.Recover())
	r.echo.Use(middleware.Logger())

	r.echo.GET("/health", healthController.Health)
}

func (r *Rest) start() error {
	return r.echo.Start(fmt.Sprintf(":%d", config.Env.ApiRestPort))
}
