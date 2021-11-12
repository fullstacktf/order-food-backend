package models

type Restaurant struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
	Phone    int    `json:"phone"`
	Name     string `json:"name"`
	Address  struct {
		Street  string `json:"street"`
		ZipCode int    `json:"zipCode"`
		Country string `json:"country"`
		City    string `json:"city"`
	} `json:"address"`
	Orders struct {
		Current []interface{} `json:"current"`
		History []interface{} `json:"history"`
	} `json:"orders"`
	Menu struct {
		Products []interface{} `json:"products"`
	} `json:"menu"`
}
