package repository

import (
	"github.com/piovani/go_full/infra/database"
	"github.com/piovani/go_full/infra/storage"
)

type FileRepository struct {
	database *database.Database
}

func NewFileRepository() *FileRepository {
	return &FileRepository{
		database: database.NewDatabase(),
	}
}

func (r *FileRepository) Save(f *storage.File) error {
	db, err := r.database.Open()
	if err != nil {
		return err
	}

	return db.Save(f).Error
}

func (r *FileRepository) Find(f *storage.File) error {
	db, err := r.database.Open()
	if err != nil {
		return err
	}

	return db.Find(f).Error
}
