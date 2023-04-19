package models

import "time"

type BookRent struct {
	ID           string     `json:"id"`
	UserID       string     `json:"user_id"`
	BookID       string     `json:"book_id"`
	DateRented   time.Time  `json:"date_rented"`
	DateReturned *time.Time `json:"date_returned"`
	DueDate      time.Time  `json:"due_date"`
}
