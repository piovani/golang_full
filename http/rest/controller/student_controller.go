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
	UsecaseGetStudents usecase.GetStudentsContract
	UsecaseGetStudent  usecase.GetStudentContract
}

func NewStudentController(
	ucc usecase.CreateStudentContract,
	ucgs usecase.GetStudentsContract,
	ucg usecase.GetStudentContract,
) *StudentController {
	return &StudentController{
		UseCaseCreate:      ucc,
		UsecaseGetStudents: ucgs,
		UsecaseGetStudent:  ucg,
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

	return c.JSON(http.StatusCreated, s.getStudentsOut(students))
}

func (s *StudentController) GetStudent(c echo.Context) error {
	ID := c.Param("id")
	student, err := s.UsecaseGetStudent.Execute(ID)
	if err != nil {
		return c.String(http.StatusInternalServerError, "bad request")
	}

	return c.JSON(http.StatusCreated, s.getStudentOut(student))
}

func (s *StudentController) getStudentOut(student *entity.Student) dto.StudentOutput {
	return dto.StudentOutput{
		ID:   student.ID,
		Name: student.Name,
		Age:  student.Age,
	}
}

func (s *StudentController) getStudentsOut(students *[]entity.Student) (out []dto.StudentOutput) {
	newStudents := *students
	for i := 0; i < len(newStudents); i++ {
		out = append(out, s.getStudentOut(&newStudents[i]))
	}
	return out
}
