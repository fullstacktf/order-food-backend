package models

import (
	"comiditapp/api/enums"
)

type Order struct {
	Id           string            `json:"id"`
	RestaurantId string            `json:"restaurantId"`
	ClientId     string            `json:"clientId"`
	Status       enums.OrderStatus `json:"status"`
	TotalPrice   float64           `json:"totalPrice"`
	Products     []ProductInfo
}
