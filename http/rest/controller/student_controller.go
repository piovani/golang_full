package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/piovani/go_full/domain/entity"
	usecase "github.com/piovani/go_full/domain/use_case"
	"github.com/piovani/go_full/dto"
)

var (
	ErrorBadRequest = fmt.Errorf("bad request")
)

type StudentController struct {
	UseCaseCreate        usecase.CreateStudentContract
	UsecaseGetStudents   usecase.GetStudentsContract
	UsecaseGetStudent    usecase.GetStudentContract
	UseCaseUpdateStudent usecase.UpdateStudentContract
	UseCaseDeleteStudent usecase.DeleteStudentContract
}

func NewStudentController(
	ucc usecase.CreateStudentContract,
	ucgs usecase.GetStudentsContract,
	ucg usecase.GetStudentContract,
	ucus usecase.UpdateStudentContract,
	ucsd usecase.DeleteStudentContract,
) *StudentController {
	return &StudentController{
		UseCaseCreate:        ucc,
		UsecaseGetStudents:   ucgs,
		UsecaseGetStudent:    ucg,
		UseCaseUpdateStudent: ucus,
		UseCaseDeleteStudent: ucsd,
	}
}

func (s *StudentController) Create(c echo.Context) error {
	dto, err := s.getStudentInput(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, s.getMessageErrorOut(ErrorBadRequest))
	}

	student, err := s.UseCaseCreate.Execute(dto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, s.getMessageErrorOut(err))
	}

	return c.JSON(http.StatusCreated, s.getStudentOut(student))
}

func (s *StudentController) GetStudents(c echo.Context) error {
	students, err := s.UsecaseGetStudents.Execute()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, s.getMessageErrorOut(err))
	}

	return c.JSON(http.StatusOK, s.getStudentsOut(students))
}

func (s *StudentController) GetStudent(c echo.Context) error {
	ID := c.Param("id")
	student, err := s.UsecaseGetStudent.Execute(ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, s.getMessageErrorOut(err))
	}

	return c.JSON(http.StatusOK, s.getStudentOut(student))
}

func (s *StudentController) Update(c echo.Context) error {
	dto, err := s.getStudentInput(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, s.getMessageErrorOut(ErrorBadRequest))
	}

	student, err := s.UseCaseUpdateStudent.Execute(dto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, s.getMessageErrorOut(err))
	}

	return c.JSON(http.StatusOK, s.getStudentOut(student))
}

func (s *StudentController) Delete(c echo.Context) error {
	ID := c.Param("id")
	if err := s.UseCaseDeleteStudent.Execute(ID); err != nil {
		return c.JSON(http.StatusInternalServerError, s.getMessageErrorOut(err))
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "delete student successfully"})
}

// AUX
func (s *StudentController) getStudentInput(c echo.Context) (dto dto.StudentInput, err error) {
	if err = c.Bind(&dto); err != nil {
		return dto, err
	}
	dto.ID = c.Param("id")
	return dto, err
}

func (s *StudentController) getStudentOut(student *entity.Student) dto.StudentOutput {
	if student == nil {
		return dto.StudentOutput{}
	} else {
		return dto.StudentOutput{
			ID:   student.ID,
			Name: student.Name,
			Age:  student.Age,
		}
	}
}

func (s *StudentController) getStudentsOut(students *[]entity.Student) (out []dto.StudentOutput) {
	newStudents := *students
	for i := 0; i < len(newStudents); i++ {
		out = append(out, s.getStudentOut(&newStudents[i]))
	}
	return out
}

func (s *StudentController) getMessageErrorOut(err error) dto.MessageErrorOut {
	return dto.MessageErrorOut{
		Message: err.Error(),
	}
}
