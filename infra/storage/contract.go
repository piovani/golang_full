package storage

import "io"

type StorageContract interface {
	Donwload(path string) (io.Reader, error)
	Upload(file io.Reader) (string, error)
}

type FileRepository interface {
	Save(file *File) error
	Find(file *File) error
}
