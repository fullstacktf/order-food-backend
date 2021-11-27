package repository

import (
	"comiditapp/api/models"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UsersRepository interface {
	createUser(ctx context.Context, user models.User) // Endpoint -> /auth/signup
	updateUser(ctx context.Context, user models.User) // Endpoint -> /profile
	findUserById(ctx context.Context, id primitive.ObjectID)
}
