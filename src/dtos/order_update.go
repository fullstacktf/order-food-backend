package dtos

import "comiditapp/api/enums"

type OrderUpdate struct {
	Status enums.OrderStatus `json:"status" bson:"status"  validate:"required"`
}
