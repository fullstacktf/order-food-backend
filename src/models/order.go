package models

import (
	"comiditapp/api/enums"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	Id           primitive.ObjectID `json:"id"`
	RestaurantId primitive.ObjectID `json:"restaurantId"`
	ClientId     primitive.ObjectID `json:"clientId"`
	Status       enums.OrderStatus  `json:"status"`
	TotalPrice   float64            `json:"totalPrice"`
	Products     []ProductInfo
}
