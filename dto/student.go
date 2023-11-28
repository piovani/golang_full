package dto

import (
	"github.com/google/uuid"
	"github.com/piovani/go_full/infra/storage"
)

type StudentInput struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Document storage.File
}

type StudentOutput struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Age  int       `json:"age"`
}

type StudentDocumentOutput struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Type string    `json:"type"`
	URL  string    `json:"url"`
}
