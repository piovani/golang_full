package usecase

import "github.com/piovani/go_full/domain/entity"

type GetStudents struct{}

func NewGetStudents() *GetStudents {
	return &GetStudents{}
}

func (s *GetStudents) Execute() ([]*entity.Student, error) {
	return entity.Students, nil
}
