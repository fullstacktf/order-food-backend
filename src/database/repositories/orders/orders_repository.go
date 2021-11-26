package orders_repository

import (
	"comiditapp/api/models"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrdersRepository interface {
	createOrder(ctx context.Context, order models.Order)      // Endpoint -> /restaurants/:id/orders
	findOrderById(ctx context.Context, id primitive.ObjectID) // Endpoint -> /profile/orders/:id
	UpdateOrder(ctx context.Context, order models.Order)      // Endpoint -> /orders/:id
}
