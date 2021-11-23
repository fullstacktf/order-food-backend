package models

import "comiditapp/api/enums"

type Product struct {
	Id       string                `json:"id"`
	Category enums.ProductCategory `json:"type"`
	Name     string                `json:"name"`
	Price    float64               `json:"price"`
}
