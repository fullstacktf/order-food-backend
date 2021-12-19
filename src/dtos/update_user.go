package dtos

import (
	"comiditapp/api/models"
)

type UpdateUser struct {
	Name     string           `json:"name" bson:"name" validate:"required"`
	Email    string           `json:"email" bson:"email" validate:"required,email"`
	Password string           `json:"password" bson:"password" validate:"required"`
	Pass     string           `json:"pass" bson:"pass" validate:"required"`
	Phone    int              `json:"phone" bson:"phone" validate:"required"`
	Address  []models.Address `json:"address" bson:"address" validate:"required"`
	Menu     []models.Product `json:"menu,omitempty" bson:"menu,omitempty"`
}
