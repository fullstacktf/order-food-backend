package models

type ProductInfo struct {
	ProductId string `json:"productId" bson:"productId"`
	Quantity  int    `json:"quantity" bson:"quantity"`
}
