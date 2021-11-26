package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProductInfo struct {
	ProductId primitive.ObjectID `json:"productId"`
	Quantity  int                `json:"quantity"`
}
