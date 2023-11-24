package entity

import (
	"time"

	"github.com/google/uuid"
)

var (
	Students = []*Student{}
)

type Student struct {
	ID         uuid.UUID `gorm:"column:id;primaryKey"`
	Name       string    `gorm:"column:name;size:255"`
	Age        int       `gorm:"column:age"`
	DocumentID uuid.UUID `gorm:"column:document_id"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
	DeletedAt  time.Time `gorm:"column:deleted_at;default:null;index"`
}

func NewStudent(name string, age int, documentID uuid.UUID) *Student {
	now := time.Now()
	return &Student{
		ID:         uuid.New(),
		Name:       name,
		Age:        age,
		DocumentID: documentID,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}

type StudentRepository interface {
	Save(student *Student) error
	All(students *[]Student) error
	Find(student *Student) error
}
