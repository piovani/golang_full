package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/piovani/go_full/domain/entity"
	usecase "github.com/piovani/go_full/domain/use_case"
	"github.com/piovani/go_full/dto"
)

type StudentController struct {
	UseCaseCreate      usecase.CreateStudentContract
	UsecaseGetStudents usecase.GetStudentContract
}

func NewStudentController(ucc usecase.CreateStudentContract, ucgs usecase.GetStudentContract) *StudentController {
	return &StudentController{
		UseCaseCreate:      ucc,
		UsecaseGetStudents: ucgs,
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

func (s *StudentController) GetStudents(c echo.Context) error {
	students, err := s.UsecaseGetStudents.Execute()
	if err != nil {
		return c.String(http.StatusInternalServerError, "bad request")
	}

	return c.JSON(http.StatusCreated, students)
}

func (s *StudentController) getStudentOut(student *entity.Student) dto.StudentOutput {
	return dto.StudentOutput{
		ID:   student.ID,
		Name: student.Name,
		Age:  student.Age,
	}
}

func (s *StudentController) getStudentsOut(students []*entity.Student) (out []dto.StudentOutput) {
	for _, student := range students {
		out = append(out, s.getStudentOut(student))
	}
	return out
}
