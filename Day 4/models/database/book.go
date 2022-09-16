package database

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Id        *int           `gorm:"autoIncrement;primary key" json:"-"`
	Title     string         `gorm:"type:text;not null" json:"title"`
	Writer    string         `gorm:"type:text;not null" json:"writer"`
	ISBN      string         `gorm:"type:text;not null" json:"isbn"`
}
