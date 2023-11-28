package usecase

import (
	"github.com/piovani/go_full/domain/entity"
	"github.com/piovani/go_full/dto"
	"github.com/piovani/go_full/infra/storage"
)

type CreateStudent struct {
	studentRepository entity.StudentRepository
	fileRepository    storage.FileRepository
}

func NewCreateStudent(sr entity.StudentRepository, fr storage.FileRepository) *CreateStudent {
	return &CreateStudent{
		studentRepository: sr,
		fileRepository:    fr,
	}
}

func (c *CreateStudent) Execute(dto dto.StudentInput) (student *entity.Student, err error) {
	if err = c.rules(dto); err != nil {
		return student, err
	}

	if err = dto.Document.Upload(); err != nil {
		return student, err
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

func (c *CreateStudent) rules(dto.StudentInput) error {
	return nil
}
