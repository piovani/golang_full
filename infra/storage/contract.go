package storage

import "io"

type StorageContract interface {
	// Donwload(path string) (io.Reader, error)
	Upload(file io.Reader) (string, error)
}
