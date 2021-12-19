package dtos

import (
	"comiditapp/api/enums"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateOrUpdateProduct struct {
	Token    string                `json:"token" bson:"token" validate:"required"`
	Id       primitive.ObjectID    `json:"id" bson:"id"`
	Category enums.ProductCategory `json:"category" bson:"category" validate:"required"`
	Name     string                `json:"name" bson:"name" validate:"required"`
	Price    float64               `json:"price" bson:"price" validate:"required"`
}
