package storage

import (
	"io"

	"github.com/google/uuid"
)

type File struct {
	ID      uuid.UUID `gorm:"column:id;primaryKey"`
	Kind    string    `gorm:"column:kind;size:255"`
	Path    string    `gorm:"column:path;size:255"`
	storage StorageContract
	file    io.Reader
}

func NewFie(f io.Reader) *File {
	return &File{
		ID:      uuid.New(),
		file:    f,
		storage: NewStorage(),
	}
}

func (f *File) Save() error {
	path, err := f.storage.Upload(f.file)
	if err != nil {
		return err
	}

	f.Kind = ".pdf"
	f.Path = path
	return nil
}

// func (f *File) SetPath(path string) {
// 	f.path = path
// }

// func (f *File) GetPath() string {
// 	return f.path
// }
