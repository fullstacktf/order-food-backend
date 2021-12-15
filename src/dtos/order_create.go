package dtos

import (
	"comiditapp/api/models"
)

type OrderCreate struct {
	Products []models.ProductInfo `json:"products,omitempty" bson:"products,omitempty" validate:"required,dive,required"`
}
