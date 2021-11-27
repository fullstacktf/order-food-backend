package repository

import (
	"comiditapp/api/models"

	"github.com/gin-gonic/gin"
)

type OrdersRepository interface {
	// any_role methods
	FindOrders(context *gin.Context) []models.Order // Endpoint -> /profile/orders
	GetOrderById(context *gin.Context) models.Order // Endpoint -> /profile/orders/:id
	CreateOrder(context *gin.Context) models.Order  // Endpoint -> /restaurants/:id/orders

	// restaurant_role methods
	UpdateClientOrder(context *gin.Context) models.Order  // Endpoint -> /orders/:id
	FindClientOrders(context *gin.Context) []models.Order // Endpoint -> /orders
}
