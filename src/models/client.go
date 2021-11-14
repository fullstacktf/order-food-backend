package models

type Client struct {
	Id       string       `json:"id"`
	Email    string       `json:"email"`
	Password string       `json:"password"`
	Phone    int          `json:"phone"`
	Name     ClientName   `json:"name"`
	Orders   ClientOrders `json:"orders"`
}
