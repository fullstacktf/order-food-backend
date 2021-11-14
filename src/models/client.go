package models

type Client struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    int    `json:"phone"`
	Name     struct {
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"name"`
	Orders ClientOrders `json:"orders"`
}
