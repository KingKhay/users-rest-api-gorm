package dto

type UpdateUserDTO struct {
	ID    uint
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}
