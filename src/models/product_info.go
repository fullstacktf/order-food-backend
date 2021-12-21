package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProductInfo struct {
	ProductId primitive.ObjectID `json:"productId" bson:"productId"`
	Name      string             `json:"name" bson:"name" validate:"required"`
	Quantity  int                `json:"quantity" bson:"quantity" validate:"required"`
	Price     float64            `json:"price" bson:"price" validate:"required"`
}
