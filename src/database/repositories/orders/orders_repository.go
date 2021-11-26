package orders_repository

import (
	"comiditapp/api/models"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrdersRepository interface {
	createOrder(ctx context.Context, user models.User)      // Endpoint -> /restaurants/:id/orders
	findOrderById(ctx context.Context, user models.User)    // Endpoint -> /profile/orders/:id
	UpdateOrder(ctx context.Context, id primitive.ObjectID) // Endpoint -> /orders/:id
}
