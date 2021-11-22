package models

import "comiditapp/api/src/enums"

type Order struct {
	Id           string            `json:"id"`
	RestaurantId string            `json:"restaurantId"`
	ClientId     string            `json:"clientId"`
	Status       enums.OrderStatus `json:"status"`
	TotalPrice   int               `json:"totalPrice"`
	Products     []struct {
		ProductId string `json:"productId"`
		Quantity  string `json:"quantity"`
	}
}
