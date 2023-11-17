package entity

import "github.com/piovani/go_full/infra/uuid"

var (
	Students = []*Student{}
)

type Student struct {
	ID   string `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name"`
	Age  int    `gorm:"column:age"`
}

func NewStudent(name string, age int) *Student {
	return &Student{
		ID:   uuid.New(),
		Name: name,
		Age:  age,
	}
}

type StudentRepository interface {
	Save(student *Student) error
	All(students []*Student) error
}
