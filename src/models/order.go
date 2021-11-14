package models

type Order struct {
	Id           string `json:"id"`
	ClientId     string `json:"clientId"`
	RestaurantId string `json:"restaurantId"`
	Status       string `json:"status"`
	TotalPrice   int    `json:"totalPrice"`
	Products     []struct {
		ProductId string `json:"productId"`
		Quantity  string `json:"quantity"`
	}
}
