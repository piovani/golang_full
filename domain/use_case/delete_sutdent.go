package usecase

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/piovani/go_full/domain/entity"
)

type DeleteStudent struct {
	studentRepository entity.StudentRepository
}

func NewDeleteStudent(sr entity.StudentRepository) *DeleteStudent {
	return &DeleteStudent{
		studentRepository: sr,
	}
}

func (d *DeleteStudent) Execute(ID string) error {
	UUID, err := uuid.Parse(ID)
	if err != nil {
		return err
	}

	student := entity.Student{ID: UUID}

	if d.studentRepository.Find(&student); err != nil {
		return err
	}

	if student.Name == "" {
		return fmt.Errorf("student not found")
	}

	if d.studentRepository.Delete(&student); err != nil {
		return err
	}

	return nil
}
