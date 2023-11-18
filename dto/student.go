package dto

import "github.com/google/uuid"

type StudentInput struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type StudentOutput struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Age  int       `json:"age"`
}
