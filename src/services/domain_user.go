package services

import (
	"comiditapp/api/dtos"
	"comiditapp/api/models"
)

func UserToDomain(user models.User) dtos.DomainUser {
	return dtos.DomainUser{
		Id:      user.Id,
		Role:    user.Role,
		Name:    user.Name,
		Email:   user.Email,
		Phone:   user.Phone,
		Address: user.Address,
		Menu:    user.Menu,
	}
}
