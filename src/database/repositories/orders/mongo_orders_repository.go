package repository

import (
	"comiditapp/api/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoOrdersRepository struct {
	db *mongo.Database
}

func NewMongoOrdersRepository(db *mongo.Database) *MongoOrdersRepository {
	return &MongoOrdersRepository{db: db}
}

// any_role methods
func (r *MongoOrdersRepository) FindOrders(context *gin.Context) *[]models.Order {
	return &[]models.Order{}
}
func (r *MongoOrdersRepository) GetOrderById(context *gin.Context) *models.Order {
	return &models.Order{}
}
func (r *MongoOrdersRepository) CreateOrder(context *gin.Context) *models.Order {
	return &models.Order{}
}

// restaurant_role methods
func (r *MongoOrdersRepository) UpdateClientOrder(context *gin.Context) *models.Order {
	return &models.Order{}
}
func (r *MongoOrdersRepository) FindClientOrders(context *gin.Context) *[]models.Order {
	return &[]models.Order{}
}
