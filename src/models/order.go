package models

import (
	"comiditapp/api/enums"
)

type Order struct {
	Id           string            `json:"id" bson:"id" validate:"required"`
	RestaurantId string            `json:"restaurantId" bson:"restaurantId" validate:"required"`
	ClientId     string            `json:"clientId" bson:"clientId" validate:"required"`
	Status       enums.OrderStatus `json:"status" bson:"status"  validate:"required"`
	TotalPrice   float64           `json:"totalPrice" bson:"totalPrice" validate:"required"`
	Products     []ProductInfo     `json:"products,omitempty" bson:"products,omitempty" validate:"required"`
}
