package usecase

import (
	"github.com/piovani/go_full/domain/entity"
	"github.com/piovani/go_full/dto"
	"github.com/piovani/go_full/infra/storage"
)

type CreateStudent struct {
	storage           storage.StorageContract
	studentRepository entity.StudentRepository
	fileRepository    storage.FileRepository
}

func NewCreateStudent(
	s storage.StorageContract,
	sr entity.StudentRepository,
	fr storage.FileRepository,
) *CreateStudent {
	return &CreateStudent{
		storage:           s,
		studentRepository: sr,
		fileRepository:    fr,
	}
}

func (c *CreateStudent) Execute(dto dto.StudentInput) (student *entity.Student, err error) {
	if err = c.storage.Upload(&dto.Document); err != nil {
		return nil, err
	}

	if err = c.fileRepository.Save(&dto.Document); err != nil {
		return student, err
	}

	student = entity.NewStudent(dto.Name, dto.Age, dto.Document.ID)
	if err = c.studentRepository.Save(student); err != nil {
		return student, err
	}

	return student, nil
}
