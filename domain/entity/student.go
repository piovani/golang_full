package entity

import "github.com/piovani/go_full/infra/uuid"

var (
	Students = []*Student{}
)

type Student struct {
	ID   string
	Name string
	Age  int
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
}
