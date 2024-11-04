package models

type Book struct {
	ID     string  `json:"id,omitempty"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}
