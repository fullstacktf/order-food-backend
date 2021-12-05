package repository

import (
	"comiditapp/api/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
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
// A este método probablemente sea al que tengamos que añadir cosas del JWT
func (r *MongoUsersRepository) SignUpUser(context *gin.Context) (statusCode int, response interface{}) {

	var validate *validator.Validate = validator.New()
	var newUser models.User

	json.NewDecoder(context.Request.Body).Decode(&newUser)

	err := validate.Struct(newUser)
	if err != nil {
		validatorError := err.(validator.ValidationErrors).Error()
		errorMessage := "Cannot create user, required fields not provided\n" + validatorError
		return http.StatusBadRequest, errorMessage
	}

	newId := primitive.NewObjectID()
	user := &models.User{
		Id:             newId,
		Role:           newUser.Role,
		Name:           newUser.Name,
		Email:          newUser.Email,
		HashedPassword: newUser.HashedPassword,
		Phone:          newUser.Phone,
		Address:        newUser.Address,
		Menu:           newUser.Menu,
	}

	if _, err := r.users.InsertOne(context, user); err != nil {
		return http.StatusBadRequest, err.Error()
	}

	message := "User " + newId.Hex() + " updated succesfully"
	return http.StatusCreated, message
}

// TODO: Implement JWT auth
func (r *MongoUsersRepository) SignInUser(context *gin.Context) (statusCode int, response interface{}) {
	return 0, &models.User{}
}
func (r *MongoUsersRepository) FindRestaurants(context *gin.Context) (statusCode int, response interface{}) {

	foundRestaurants, err := r.users.Find(context, bson.M{"role": "restaurant"})
	if err != nil {
		return http.StatusConflict, err.Error()
	}

	restaurants := []*models.User{}
	if err := foundRestaurants.All(context, &restaurants); err != nil {
		return http.StatusConflict, err.Error()
	}

	return http.StatusOK, restaurants
}
func (r *MongoUsersRepository) GetRestaurantById(context *gin.Context) (statusCode int, response interface{}) {

	var restaurant models.User

	// De el filtro aplicado el que no funciona es el de id
	filter := bson.M{"role": "restaurant", "id": context.Param("id")}
	if err := r.users.FindOne(context, filter).Decode(&restaurant); err != nil {
		return http.StatusNotFound, err.Error()
	}

	return http.StatusOK, restaurant
}

func (r *MongoUsersRepository) FindClients(context *gin.Context) (statusCode int, response interface{}) {

	foundRestaurants, err := r.users.Find(context, bson.M{"role": "client"})
	if err != nil {
		return http.StatusConflict, err.Error()
	}

	clients := []*models.User{}
	if err := foundRestaurants.All(context, &clients); err != nil {
		return http.StatusConflict, err.Error()
	}

	return http.StatusOK, clients
}

func (r *MongoUsersRepository) GetClientById(context *gin.Context) (statusCode int, response interface{}) {
	var client models.User

	// De el filtro aplicado el que no funciona es el de id
	filter := bson.M{"role": "client", "id": context.Param("id")}
	if err := r.users.FindOne(context, filter).Decode(&client); err != nil {
		return http.StatusNotFound, err.Error()
	}

	return http.StatusOK, &client
}

// Hasta que no funcione el buscar un elemento por id poco podemos hacer con estos endpoints
func (r *MongoUsersRepository) GetRestaurantProducts(context *gin.Context) (statusCode int, response interface{}) {
	return 0, &[]models.Product{}
}
func (r *MongoUsersRepository) UpdateProfile(context *gin.Context) (statusCode int, response interface{}) {
	return 0, &models.User{}
}

// restaurant_role methods
// Para estos necesitamos comprobar que el restaurante tiene permisos y sacar su id de la sesion para poder efectuar las acciones
func (r *MongoUsersRepository) FindProducts(context *gin.Context) (statusCode int, response interface{}) {
	return 0, &[]models.Product{}
}
func (r *MongoUsersRepository) CreateProduct(context *gin.Context) (statusCode int, response interface{}) {
	return 0, &models.Product{}
}
func (r *MongoUsersRepository) UpdateProduct(context *gin.Context) (statusCode int, response interface{}) {
	return 0, &models.Product{}
}
