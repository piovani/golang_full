package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/piovani/go_full/domain/entity"
	usecase "github.com/piovani/go_full/domain/use_case"
	"github.com/piovani/go_full/dto"
)

type StudentController struct {
	UseCaseCreate usecase.CreateStudentContract
}

func NewStudentController(useCaseCreate usecase.CreateStudentContract) *StudentController {
	return &StudentController{
		UseCaseCreate: useCaseCreate,
	}
}

func (s *StudentController) Create(c echo.Context) error {
	var studentDTO dto.StudentInput
	if c.Bind(&studentDTO) != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	student, err := s.UseCaseCreate.Execute(studentDTO)
	if err != nil {
		return c.String(http.StatusInternalServerError, "bad request")
	}

	return c.JSON(http.StatusCreated, s.getStudentOut(student))
}

func (s *StudentController) getStudentOut(student *entity.Student) dto.StudentOutput {
	return dto.StudentOutput{
		Name: student.Name,
		Age:  student.Age,
	}
}
