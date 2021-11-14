package models

type Restaurant struct {
	Id       string           `json:"id"`
	Email    string           `json:"email"`
	Password string           `json:"password"`
	Username string           `json:"username"`
	Phone    int              `json:"phone"`
	Name     string           `json:"name"`
	Address  Address          `json:"address"`
	Orders   RestaurantOrders `json:"orders"`
	Menu     []Product        `json:"menu"`
}
