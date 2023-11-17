package usecase

import (
	"github.com/piovani/go_full/domain/entity"
	"github.com/piovani/go_full/dto"
)

type CreateStudent struct {
	studentRepository entity.StudentRepository
}

func NewCreateStudent(sr entity.StudentRepository) *CreateStudent {
	return &CreateStudent{
		studentRepository: sr,
	}
}

func (c *CreateStudent) Execute(dto dto.StudentInput) (student *entity.Student, err error) {
	if err = c.rules(dto); err != nil {
		return student, err
	}

	student = entity.NewStudent(dto.Name, dto.Age)
	entity.Students = append(entity.Students, student)
	c.studentRepository.Save(student)

	return student, nil
}

func (c *CreateStudent) rules(dto.StudentInput) error {
	return nil
}
