package storage

type FileRepository interface {
	Save(file *File) error
	Find(file *File) error
}
