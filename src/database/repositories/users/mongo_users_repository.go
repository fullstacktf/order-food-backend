package repository

import (
	"comiditapp/api/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUsersRepository struct {
	db *mongo.Database
}

func NewMongoUsersRepository(db *mongo.Database) *MongoUsersRepository {
	return &MongoUsersRepository{db: db}
}

// any_role methods
func (r *MongoUsersRepository) SignUpUser(context *gin.Context) *models.User
func (r *MongoUsersRepository) SignInUser(context *gin.Context) *models.User
func (r *MongoUsersRepository) FindRestaurants(context *gin.Context) *[]models.User
func (r *MongoUsersRepository) GetRestaurantById(context *gin.Context) *models.User
func (r *MongoUsersRepository) FindClients(context *gin.Context) *[]models.User
func (r *MongoUsersRepository) GetClientById(context *gin.Context) *models.User
func (r *MongoUsersRepository) GetRestaurantProducts(context *gin.Context) *[]models.Product
func (r *MongoUsersRepository) UpdateProfile(context *gin.Context) *models.User

// restaurant_role methods
func (r *MongoUsersRepository) FindProducts(context *gin.Context) *[]models.Product
func (r *MongoUsersRepository) CreateProduct(context *gin.Context) *models.Product
func (r *MongoUsersRepository) UpdateProduct(context *gin.Context) *models.Product
