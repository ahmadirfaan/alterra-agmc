package database

import (
	"time"
)

type User struct {
	Id        *int      `gorm:"autoIncrement;primary key" json:"-"`
	Name      string    `gorm:"type:text;not null" json:"name"`
	Email     string    `gorm:"uniqueIndex;type:varchar(500);not null" json:"email"`
	Password  string    `gorm:"type:text;not null" json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
