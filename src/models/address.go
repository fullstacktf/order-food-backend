package models

type Address struct {
	Street  string `json:"street"`
	ZipCode int    `json:"zipCode"`
	Country string `json:"country"`
	City    string `json:"city"`
}
