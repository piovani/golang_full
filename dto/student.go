package dto

import (
	"io"

	"github.com/google/uuid"
)

type StudentInput struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	File io.Reader
}

type StudentOutput struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Age  int       `json:"age"`
}
