package usecase

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/piovani/go_full/domain/entity"
)

type GetStudent struct {
	studentRepository entity.StudentRepository
}

func NewGetStudent(sr entity.StudentRepository) *GetStudent {
	return &GetStudent{
		studentRepository: sr,
	}
}

func (s *GetStudent) Execute(ID string) (*entity.Student, error) {
	var student entity.Student

	UUID, err := uuid.Parse(ID)
	if err != nil {
		return &student, err
	}

	student.ID = UUID
	if err = s.studentRepository.Find(&student); err != nil {
		return &student, err
	}

	if student.Name == "" {
		return &student, fmt.Errorf("student not found")
	}

	return &student, nil
}
