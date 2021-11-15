package models

type Client struct {
	Id       string `json:"id"`
	Role 	 string `json:"role"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    int    `json:"phone"`
	Orders struct {
		Current []Order `json:"current"`
		History []Order `json:"history"`
	} `json:"orders"`
	Address  []Address  `json:"address"`
	Menu     []Product  `json:"menu"`
}
