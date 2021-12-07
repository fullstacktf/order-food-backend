package models

import (
	"comiditapp/api/enums"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	Id           primitive.ObjectID `json:"id" bson:"id"`
	RestaurantId primitive.ObjectID `json:"restaurantId" bson:"restaurantId"`
	ClientId     primitive.ObjectID `json:"clientId" bson:"clientId" validate:"required"` // cuando tengamos JWT quitaremos el validate de aqui
	Status       enums.OrderStatus  `json:"status" bson:"status"  validate:"required"`
	TotalPrice   float64            `json:"totalPrice" bson:"totalPrice" validate:"required"`
	Products     []ProductInfo      `json:"products,omitempty" bson:"products,omitempty" validate:"required,dive,required"`
}
