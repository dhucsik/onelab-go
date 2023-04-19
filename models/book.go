package models

type Book struct {
	ID        string  `json:"id"`
	Title     string  `json:"title"`
	Author    string  `json:"author"`
	RentPrice float64 `json:"rent_price"`
}

type BookUserIncome struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Author      string     `json:"author"`
	Users       []UserResp `json:"users"`
	TotalIncome float64    `json:"total_income"`
}
