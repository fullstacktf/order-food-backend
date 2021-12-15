package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProductInfo struct {
	ProductId primitive.ObjectID `json:"productId" bson:"productId"`
	Quantity  int                `json:"quantity" bson:"quantity" validate:"required"`
	Price     float64            `json:"price" bson:"price" validate:"required"`
}
