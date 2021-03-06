package models

import (
	"comiditapp/api/enums"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Id       primitive.ObjectID    `json:"id" bson:"id"`
	Category enums.ProductCategory `json:"category" bson:"category" validate:"required"`
	Name     string                `json:"name" bson:"name" validate:"required"`
	Price    float64               `json:"price" bson:"price" validate:"required"`
}
