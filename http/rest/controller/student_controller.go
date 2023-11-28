package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/piovani/go_full/domain/entity"
	usecase "github.com/piovani/go_full/domain/use_case"
	"github.com/piovani/go_full/dto"
	"github.com/piovani/go_full/infra/storage"
)

var (
	ErrorBadRequest = fmt.Errorf("bad request")
)

type StudentController struct {
	UseCaseCreate             usecase.CreateStudentContract
	UsecaseGetStudents        usecase.GetStudentsContract
	UsecaseGetStudent         usecase.GetStudentContract
	UseCaseUpdateStudent      usecase.UpdateStudentContract
	UseCaseDeleteStudent      usecase.DeleteStudentContract
	UseCaseGetDocumentStudent usecase.GetDocumentContract
}

func NewStudentController(
	ucc usecase.CreateStudentContract,
	ucgs usecase.GetStudentsContract,
	ucg usecase.GetStudentContract,
	ucus usecase.UpdateStudentContract,
	ucsd usecase.DeleteStudentContract,
	ucgd usecase.GetDocumentContract,
) *StudentController {
	return &StudentController{
		UseCaseCreate:             ucc,
		UsecaseGetStudents:        ucgs,
		UsecaseGetStudent:         ucg,
		UseCaseUpdateStudent:      ucus,
		UseCaseDeleteStudent:      ucsd,
		UseCaseGetDocumentStudent: ucgd,
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

func (s *StudentController) GetDocumentStudent(c echo.Context) error {
	ID := c.Param("id")

	document, err := s.UseCaseGetDocumentStudent.Execute(ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, s.getMessageErrorOut(err))
	}

	return c.JSON(http.StatusOK, s.getStudentDocumentOut(&document))
}

// AUX
func (s *StudentController) getStudentInput(c echo.Context) (dto dto.StudentInput, err error) {
	if err = c.Bind(&dto); err != nil {
		return dto, err
	}
	dto.ID = c.Param("id")

	if dto.Name == "" {
		dto.Name = c.FormValue("name")
		age := c.FormValue("age")
		dto.Age, err = strconv.Atoi(age)
		if err != nil {
			return dto, err
		}

		fileRequest, err := c.FormFile("file")
		if err != nil {
			return dto, err
		}

		header, err := fileRequest.Open()
		if err != nil {
			return dto, err
		}
		defer header.Close()

		dto.Document = *storage.NewFie(fileRequest.Filename, fileRequest.Header.Get("Content-Type"), header)
	}

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

func (s *StudentController) getStudentDocumentOut(document *storage.File) dto.StudentDocumentOutput {
	if document == nil {
		return dto.StudentDocumentOutput{}
	} else {
		return dto.StudentDocumentOutput{
			ID:   document.ID,
			Name: document.Name,
			Type: document.Kind,
			URL:  document.Path,
		}
	}
}

func (s *StudentController) getMessageErrorOut(err error) dto.MessageErrorOut {
	return dto.MessageErrorOut{
		Message: err.Error(),
	}
}
