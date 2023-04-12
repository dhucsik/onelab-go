package models

import "time"

type Record struct {
	ID           string     `json:"id"`
	UserID       string     `json:"user_id"`
	BookID       string     `json:"book_id"`
	DateBorrowed time.Time  `json:"date_borrowed"`
	DateReturned *time.Time `json:"date_returned"`
	DueDate      time.Time  `json:"due_date"`
}
