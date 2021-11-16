package models

type Order struct {
	Id           string `json:"id"`
	Status       string `json:"status"`
	TotalPrice   int    `json:"totalPrice"`
	Products     []struct {
		ProductId string `json:"productId"`
		Quantity  string `json:"quantity"`
	}
}
