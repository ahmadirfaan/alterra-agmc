package model

import "time"

type Book struct {
	ID        uint   `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	Writer    string `json:"writer,omitempty"`
	ISBN      string `json:"isbn,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

var BookData = []Book{
	{ID: 1, Title: "Judul Buku Satu", Writer: "Dr. Who 1", ISBN: "1-234 -5678-9101112-13", CreatedAt: time.Now().Format("2006-01-02 15:04:05")},
	{ID: 2, Title: "Judul Buku Dua", Writer: "Dr. Who 2", ISBN: "1-234 -5678-9101112-14", CreatedAt: time.Now().Format("2006-01-02 15:04:05")},
	{ID: 3, Title: "Judul Buku Tiga", Writer: "Dr. Who 3", ISBN: "1-234 -5678-9101112-15", CreatedAt: time.Now().Format("2006-01-02 15:04:05")},
}
