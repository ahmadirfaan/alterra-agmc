package models

type CreateUserRequest struct {
	Name     string `json:"name"  validate:"required"`
	Email    string `json:"email"  validate:"email,required"`
	Password string `json:"password"  validate:"required,min=8,max=100"`
}
