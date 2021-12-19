package dtos

import "comiditapp/api/enums"

type OrderUpdate struct {
	Token  string            `json:"token" bson:"token" validate:"required"`
	Status enums.OrderStatus `json:"status" bson:"status"  validate:"required"`
}
