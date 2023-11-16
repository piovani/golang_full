package dto

type StudentInput struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type StudentOutput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
