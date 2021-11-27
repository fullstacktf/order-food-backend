package models

import (
	"comiditapp/api/enums"
)

type Order struct {
	Id           string            `json:"id" validate:"required"`
	RestaurantId string            `json:"restaurantId" validate:"required"`
	ClientId     string            `json:"clientId" validate:"required"`
	Status       enums.OrderStatus `json:"status"  validate:"required"`
	TotalPrice   float64           `json:"totalPrice" validate:"required"`
	Products     []ProductInfo     `json:"products" validate:"required"`
}
