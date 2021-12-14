package dtos

import (
	"comiditapp/api/models"
)

type OrderCreate struct {
	TotalPrice float64              `json:"totalPrice" bson:"totalPrice" validate:"required"` // Este tampoco deberia ponerse, ya que deberia calcularse automatico
	Products   []models.ProductInfo `json:"products,omitempty" bson:"products,omitempty" validate:"required,dive,required"`
}
