package models

type CreateBookRequest struct {
	Title  string `json:"title"  validate:"required"`
	Writer string `json:"writer"  validate:"required"`
	ISBN   string `json:"isbn"  validate:"required"`
}
