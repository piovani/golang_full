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

	file := storage.NewFie(dto.File)
	if err = file.Save(); err != nil {
		return student, err
	}

	if err = c.fileRepository.Save(file); err != nil {
		return student, err
	}

	student = entity.NewStudent(dto.Name, dto.Age, file.ID)
	if err = c.studentRepository.Save(student); err != nil {
		return student, err
	}

	return student, nil
}

func (c *CreateStudent) rules(dto.StudentInput) error {
	return nil
}
