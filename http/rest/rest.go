package rest

import (
	"fmt"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	usecase "github.com/piovani/go_full/domain/use_case"
	"github.com/piovani/go_full/http/rest/controller"
	"github.com/piovani/go_full/infra/config"
	"github.com/piovani/go_full/infra/database/repository"
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
	// REPOSITORIES
	studentRepository := repository.NewStudentRepository()

	// CONTROLLERS
	healthController := controller.NewHealthController()
	studentController := controller.NewStudentController(
		usecase.NewCreateStudent(studentRepository),
		usecase.NewGetStudents(studentRepository),
	)

	r.echo.Use(middleware.Recover())
	r.echo.Use(middleware.Logger())

	r.echo.GET("/health", healthController.Health)

	r.echo.POST("/student", studentController.Create)
	r.echo.GET("/student", studentController.GetStudents)
}

func (r *Rest) start() error {
	return r.echo.Start(fmt.Sprintf(":%d", config.Env.ApiRestPort))
}
