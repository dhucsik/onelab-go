package models

type Book struct {
	ID     string `param:"id" json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Price  uint   `json:"price"`
}
