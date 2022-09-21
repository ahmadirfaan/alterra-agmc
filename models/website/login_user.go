package models

type LoginUserRequest struct {
	Email    string `json:"email"  validate:"email,required"`
	Password string `json:"password"  validate:"required,min=8,max=100"`
}
