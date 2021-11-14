package models

type ClientOrders struct {
	Current string   `json:"current"`
	History []string `json:"history"`
}
