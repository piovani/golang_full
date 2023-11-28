package storage

import (
	"io"

	"github.com/google/uuid"
)

type File struct {
	ID     uuid.UUID `gorm:"column:id;primaryKey"`
	Name   string    `gorm:"column:name;size:255"`
	Kind   string    `gorm:"column:kind;size:255"`
	Path   string    `gorm:"column:path;size:255"`
	Reader io.Reader
}

func NewFie(name, kind string, reader io.Reader) *File {
	return &File{
		ID:     uuid.New(),
		Name:   name,
		Kind:   kind,
		Reader: reader,
	}
}

// func (f *File) Upload() error {
// 	path, err := NewStorage().Upload(f.reader)
// 	if err != nil {
// 		return err
// 	}
// 	f.Path = path
// 	return nil
// }

// func (f *File) Download() (io.Reader, error) {
// 	pathSplit := strings.Split(f.Path, "/")
// 	return NewStorage().Donwload(pathSplit[4])
// }
