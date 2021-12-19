package dtos

import (
	"comiditapp/api/models"
)

type OrderCreate struct {
	Token    string               `json:"token" bson:"token" validate:"required"`
	Products []models.ProductInfo `json:"products,omitempty" bson:"products,omitempty" validate:"required,dive,required"`
}
