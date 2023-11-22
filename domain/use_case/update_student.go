package usecase

import (
	"time"

	"github.com/google/uuid"
	"github.com/piovani/go_full/domain/entity"
	"github.com/piovani/go_full/dto"
)

type UpdateStudent struct {
	studentRepository entity.StudentRepository
}

func NewUpdateStudent(sr entity.StudentRepository) *UpdateStudent {
	return &UpdateStudent{
		studentRepository: sr,
	}
}

func (u *UpdateStudent) Execute(dto dto.StudentInput) (*entity.Student, error) {
	UUID, err := uuid.Parse(dto.ID)
	if err != nil {
		return nil, err
	}

	student := entity.Student{ID: UUID}

	if err = u.studentRepository.Find(&student); err != nil {
		return &student, err
	}

	student.Name = dto.Name
	student.Age = dto.Age
	student.UpdatedAt = time.Now()

	if err = u.studentRepository.Save(&student); err != nil {
		return &student, err
	}

	return &student, nil
}
