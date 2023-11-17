package usecase

import (
	"github.com/piovani/go_full/domain/entity"
)

type GetStudents struct {
	studentRepository entity.StudentRepository
}

func NewGetStudents(sr entity.StudentRepository) *GetStudents {
	return &GetStudents{
		studentRepository: sr,
	}
}

func (s *GetStudents) Execute() ([]*entity.Student, error) {
	var students []*entity.Student
	if err := s.studentRepository.All(students); err != nil {
		return students, err
	}

	return students, nil
}
