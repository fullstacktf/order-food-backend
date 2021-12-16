package services

import (
	"comiditapp/api/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateMenu(menu []models.Product) []models.Product {
	parsedMenu := []models.Product{}

	for _, product := range menu {
		newProduct := models.Product{
			Id:       primitive.NewObjectID(),
			Category: product.Category,
			Name:     product.Name,
			Price:    product.Price,
		}
		parsedMenu = append(parsedMenu, newProduct)
	}

	return parsedMenu
}
