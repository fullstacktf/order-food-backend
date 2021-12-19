package dtos

type UserToken struct {
	Token string `json:"token" bson:"token" validate:"required"`
}
