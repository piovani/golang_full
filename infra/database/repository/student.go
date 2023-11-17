package repository

import (
	"fmt"

	"github.com/piovani/go_full/domain/entity"
	"github.com/piovani/go_full/infra/database"
)

type Student struct {
	ID   string `gorm:"primaryKey;column:id"`
	Name string `gorm:"column:name"`
	Age  int    `gorm:"column:age"`
}

type StudentRepository struct {
	database *database.Database
}

func NewStudentRepository() *StudentRepository {
	return &StudentRepository{
		database: database.NewDatabase(),
	}
}

func (r *StudentRepository) Save(s *entity.Student) error {
	db, _ := r.database.Open()

	result := db.Create(r.getStudentDB(s))
	fmt.Println(result)

	return nil
}

func (r *StudentRepository) getStudentDB(s *entity.Student) Student {
	return Student{
		ID:   s.ID,
		Name: s.Name,
		Age:  s.Age,
	}
}
