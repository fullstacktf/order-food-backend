package repository

import (
	"comiditapp/api/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUsersRepository struct {
	users *mongo.Collection
}

func NewMongoUsersRepository(db *mongo.Database) *MongoUsersRepository {
	return &MongoUsersRepository{users: db.Collection("users")}
}

// any_role methods
func (r *MongoUsersRepository) SignUpUser(context *gin.Context) *mongo.InsertOneResult {

	var validate *validator.Validate = validator.New()
	var newUser models.User

	json.NewDecoder(context.Request.Body).Decode(&newUser)

	err := validate.Struct(newUser)

	if err != nil {
		errorMsg := "Cannot create user, required fields not provided\n" + err.(validator.ValidationErrors).Error()

		http.Error(context.Writer, errorMsg, http.StatusBadRequest)
		panic(errorMsg)
	}

	user := &models.User{
		Id:             primitive.NewObjectID(),
		Role:           newUser.Role,
		Name:           newUser.Name,
		Email:          newUser.Email,
		HashedPassword: newUser.HashedPassword,
		Phone:          newUser.Phone,
		Address:        newUser.Address,
		Menu:           newUser.Menu,
	}

	createdUserId, err := r.users.InsertOne(context, user)
	if err != nil {
		http.Error(context.Writer, err.Error(), http.StatusConflict)
		panic(err)
	}

	// No se si deberiamos usar el context aqui dentro o delegarlo al handler
	// context.IndentedJSON(http.StatusCreated, createdUserId)

	return createdUserId
}

func (r *MongoUsersRepository) SignInUser(context *gin.Context) *models.User {
	return &models.User{}
}
func (r *MongoUsersRepository) FindRestaurants(context *gin.Context) *[]models.User {
	return &[]models.User{}
}
func (r *MongoUsersRepository) GetRestaurantById(context *gin.Context) *models.User {
	return &models.User{}
}
func (r *MongoUsersRepository) FindClients(context *gin.Context) *[]models.User {
	return &[]models.User{}
}
func (r *MongoUsersRepository) GetClientById(context *gin.Context) *models.User {
	return &models.User{}
}
func (r *MongoUsersRepository) GetRestaurantProducts(context *gin.Context) *[]models.Product {
	return &[]models.Product{}
}
func (r *MongoUsersRepository) UpdateProfile(context *gin.Context) *models.User {
	return &models.User{}
}

// restaurant_role methods
func (r *MongoUsersRepository) FindProducts(context *gin.Context) *[]models.Product {
	return &[]models.Product{}
}
func (r *MongoUsersRepository) CreateProduct(context *gin.Context) *models.Product {
	return &models.Product{}
}
func (r *MongoUsersRepository) UpdateProduct(context *gin.Context) *models.Product {
	return &models.Product{}
}
