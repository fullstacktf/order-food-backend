package dtos

import (
	"comiditapp/api/enums"
	"comiditapp/api/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DomainUser struct {
	Id      primitive.ObjectID `json:"id" bson:"id"`
	Role    enums.Role         `json:"role" bson:"role"`
	Name    string             `json:"name" bson:"name"`
	Email   string             `json:"email" bson:"email"`
	Phone   int                `json:"phone" bson:"phone"`
	Address []models.Address   `json:"address" bson:"address"`
	Menu    []models.Product   `json:"menu" bson:"menu"`
}
