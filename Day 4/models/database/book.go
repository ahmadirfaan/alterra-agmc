package database

import (
	"time"
)

type Book struct {
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Id        *int      `gorm:"autoIncrement;primary key" json:"-"`
	Title     string    `gorm:"type:text;not null" json:"title"`
	Writer    string    `gorm:"type:text;not null" json:"writer"`
	ISBN      string    `gorm:"type:text;not null" json:"isbn"`
}
