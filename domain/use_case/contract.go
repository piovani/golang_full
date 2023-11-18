package usecase

import (
	"github.com/piovani/go_full/domain/entity"
	"github.com/piovani/go_full/dto"
)

type CreateStudentContract interface {
	Execute(dto dto.StudentInput) (*entity.Student, error)
}

type GetStudentsContract interface {
	Execute() (*[]entity.Student, error)
}

type GetStudentContract interface {
	Execute(id string) (*entity.Student, error)
}
