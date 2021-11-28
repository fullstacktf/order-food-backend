package models

import "comiditapp/api/enums"

type Product struct {
	Id       string                `json:"id" bson:"id"`
	Category enums.ProductCategory `json:"type" bson:"type"`
	Name     string                `json:"name" bson:"name"`
	Price    float64               `json:"price" bson:"price"`
}
