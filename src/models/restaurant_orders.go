package models

type RestaurantOrders struct {
	Current []string `json:"current"`
	History []string `json:"history"`
}
