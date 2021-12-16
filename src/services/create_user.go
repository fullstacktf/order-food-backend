package services

import (
	"comiditapp/api/middlewares"
	"comiditapp/api/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(user models.User) (*models.User, error) {
	id := primitive.NewObjectID()

	password, err := middlewares.HashPassword(user.Password)
	if err != nil {
		return &models.User{}, err
	}

	newUser := &models.User{
		Id:       id,
		Role:     user.Role,
		Name:     user.Name,
		Email:    user.Email,
		Password: password,
		Phone:    user.Phone,
		Address:  user.Address,
		Menu:     user.Menu,
	}

	return newUser, nil
}
