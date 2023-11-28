package usecase

import (
	"github.com/google/uuid"
	"github.com/piovani/go_full/domain/entity"
	"github.com/piovani/go_full/infra/storage"
)

type GetDocumentStudent struct {
	studentRepository entity.StudentRepository
	fileRepository    storage.FileRepository
}

func NewGetDocumentStudent(sr entity.StudentRepository, fr storage.FileRepository) *GetDocumentStudent {
	return &GetDocumentStudent{
		studentRepository: sr,
		fileRepository:    fr,
	}
}

func (g *GetDocumentStudent) Execute(ID string) (file storage.File, err error) {
	UUID, err := uuid.Parse(ID)
	if err != nil {
		return file, err
	}

	student := entity.Student{ID: UUID}
	if err = g.studentRepository.Find(&student); err != nil {
		return file, err
	}

	file.ID = student.DocumentID
	if err = g.fileRepository.Find(&file); err != nil {
		return file, err
	}

	return file, nil
}
