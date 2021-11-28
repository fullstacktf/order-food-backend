package models

import (
	"comiditapp/api/enums"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Id       primitive.ObjectID    `json:"id" bson:"id"`
	Category enums.ProductCategory `json:"type" bson:"type"`
	Name     string                `json:"name" bson:"name"`
	Price    float64               `json:"price" bson:"price"`
}
