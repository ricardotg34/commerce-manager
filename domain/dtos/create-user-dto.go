package dtos

type CreateUserDTO struct {
	Name string `json:"name" binding:"required"`
}
