package models

type CreateBookRequest struct {
	Title  string `json:"title"`
	Writer string `json:"writer"`
	ISBN   string `json:"isbn"`
}
