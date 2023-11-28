package usecase

import (
	"github.com/piovani/go_full/domain/entity"
	"github.com/piovani/go_full/dto"
	"github.com/piovani/go_full/infra/storage"
)

type CreateStudentContract interface {
	Execute(dto dto.StudentInput) (*entity.Student, error)
}

type GetStudentsContract interface {
	Execute() (*[]entity.Student, error)
}

type GetStudentContract interface {
	Execute(ID string) (*entity.Student, error)
}

type UpdateStudentContract interface {
	Execute(dto dto.StudentInput) (*entity.Student, error)
}

type DeleteStudentContract interface {
	Execute(ID string) error
}

type GetDocumentContract interface {
	Execute(ID string) (storage.File, error)
}
