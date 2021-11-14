package models

type Product struct {
	Id    string  `json:"id"`
	Type  string  `json:"type"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
