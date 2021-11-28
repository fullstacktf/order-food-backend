package models

type Address struct {
	Street  string `json:"street" bson:"street"`
	ZipCode int    `json:"zipCode" bson:"zipCode"`
	Country string `json:"country" bson:"country"`
	City    string `json:"city" bson:"city"`
}
