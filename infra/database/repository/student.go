package repository

import (
	domain_entity "github.com/piovani/go_full/domain/entity"
	"github.com/piovani/go_full/infra/database"
)

type StudentRepository struct {
	database *database.Database
}

func NewStudentRepository() *StudentRepository {
	return &StudentRepository{
		database: database.NewDatabase(),
	}
}

func (r *StudentRepository) Save(s *domain_entity.Student) error {
	db, err := r.database.Open()
	if err != nil {
		return err
	}

	return db.Create(s).Error
}

func (r *StudentRepository) All(s *[]domain_entity.Student) error {
	db, err := r.database.Open()
	if err != nil {
		return err
	}

	return db.Find(s).Error
}

func (r *StudentRepository) Find(s *domain_entity.Student) error {
	db, err := r.database.Open()
	if err != nil {
		return err
	}

	return db.Find(s).Error
}
