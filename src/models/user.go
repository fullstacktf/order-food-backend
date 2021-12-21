package models

import (
	"comiditapp/api/enums"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `json:"id" bson:"id"`
	Role     enums.Role         `json:"role" bson:"role" validate:"required"`
	Name     string             `json:"name" bson:"name" validate:"required"`
	Email    string             `json:"email" bson:"email" validate:"required,email"`
	Password string             `json:"password" bson:"password" validate:"required"`
	Phone    int                `json:"phone" bson:"phone" validate:"required"`
	Address  []Address          `json:"address" bson:"address" validate:"required"`
	Menu     []Product          `json:"menu" bson:"menu"`
}
