package repository

import (
	"comiditapp/api/dtos"
	"comiditapp/api/middlewares"
	"comiditapp/api/models"
	"comiditapp/api/services"
	"context"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
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

// check if exists any user with same email, cause emails must be unique
func (r *MongoUsersRepository) DoesUserExists(u models.User) bool {
	var result bson.M

	filter := bson.M{"email": u.Email}
	if err := r.users.FindOne(context.TODO(), filter).Decode(&result); err == nil {
		return true
	}

	return false
}

func (r *MongoUsersRepository) HavePermissions(u dtos.UpdateUser) bool {
	var dbUser models.User

	filter := bson.M{"email": u.Email}
	if err := r.users.FindOne(context.TODO(), filter).Decode(&dbUser); err != nil {
		return false
	}

	if err := middlewares.VerifyPassword(u.Pass, dbUser.Password); err != nil {
		return false
	}

	return true
}

// POST - http://localhost:3000/auth/register
func (r *MongoUsersRepository) Register(user models.User) error {
	_, err := r.users.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}

// POST - http://localhost:3000/auth/login
func (r *MongoUsersRepository) Login(context *gin.Context) (statusCode int, response interface{}) {

	var u dtos.UserLogin
	var dbUser models.User

	context.BindJSON(&u)

	filter := bson.M{"email": u.Email}
	if err := r.users.FindOne(context, filter).Decode(&dbUser); err != nil {
		return http.StatusNotFound, err.Error()
	}

	if isEqual := (u.Email == dbUser.Email); isEqual != true {
		return http.StatusUnauthorized, "Login failed, email or password are incorrect "
	}
	// Bcrypt se encarga de hashear la del user y compararla con la de db
	if err := middlewares.VerifyPassword(u.Password, dbUser.Password); err != nil {
		return http.StatusUnauthorized, "Login failed, email or password are incorrect "
	}

	token, err := services.SetUserCookie(context, dbUser)
	if err != nil {
		return http.StatusInternalServerError, "Internal server error"
	}

	return http.StatusOK, gin.H{"token": token, "user": services.UserToDomain(dbUser)}
}

// GET - http://localhost:3000/restaurants
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

// GET - http://localhost:3000/restaurants/:id
func (r *MongoUsersRepository) GetRestaurantById(context *gin.Context) (statusCode int, response interface{}) {
	var restaurant models.User

	id, err := primitive.ObjectIDFromHex(context.Param("id"))
	if err != nil {
		errorMessage := "Bad request, " + context.Param("id") + " is not a valid ID"
		return http.StatusBadRequest, errorMessage
	}

	filter := bson.M{"role": "restaurant", "id": id}
	if err := r.users.FindOne(context, filter).Decode(&restaurant); err != nil {
		return http.StatusNotFound, err.Error()
	}

	return http.StatusOK, restaurant
}

// GET - http://localhost:3000/clients
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

// GET - http://localhost:3000/clients/:id
func (r *MongoUsersRepository) GetClientById(context *gin.Context) (statusCode int, response interface{}) {
	var client models.User

	id, err := primitive.ObjectIDFromHex(context.Param("id"))
	if err != nil {
		errorMessage := "Bad request, " + context.Param("id") + " is not a valid ID"
		return http.StatusBadRequest, errorMessage
	}

	filter := bson.M{"role": "client", "id": id}
	if err := r.users.FindOne(context, filter).Decode(&client); err != nil {
		return http.StatusNotFound, err.Error()
	}

	return http.StatusOK, &client
}

// GET - http://localhost:3000/profile/:id
func (r *MongoUsersRepository) GetProfileById(context *gin.Context) (statusCode int, response interface{}) {
	token := context.Query("token")

	t, _ := jwt.Parse(token, nil)
	encodedId := t.Claims.(jwt.MapClaims)["id"]
	requesterId := fmt.Sprintf("%v", encodedId)

	if requesterId != context.Param("id") {
		return http.StatusUnauthorized, "Not enough permissions"
	}

	var user models.User

	id, err := primitive.ObjectIDFromHex(context.Param("id"))
	if err != nil {
		errorMessage := "Bad request, " + context.Param("id") + " is not a valid ID"
		return http.StatusBadRequest, errorMessage
	}

	filter := bson.M{"id": id}
	if err := r.users.FindOne(context, filter).Decode(&user); err != nil {
		return http.StatusNotFound, err.Error()
	}

	return http.StatusOK, &user
}

// GET - http://localhost:3000/restaurants/:id/products
func (r *MongoUsersRepository) GetRestaurantProducts(context *gin.Context) (statusCode int, response interface{}) {
	var restaurant models.User

	id, err := primitive.ObjectIDFromHex(context.Param("id"))
	if err != nil {
		errorMessage := "Bad request, " + context.Param("id") + " is not a valid ID"
		return http.StatusBadRequest, errorMessage
	}

	filter := bson.M{"role": "restaurant", "id": id}
	if err := r.users.FindOne(context, filter).Decode(&restaurant); err != nil {
		return http.StatusNotFound, err.Error()
	}

	products := restaurant.Menu

	return http.StatusOK, products
}

// PUT - http://localhost:3000/profile/:id
func (r *MongoUsersRepository) UpdateProfile(context *gin.Context) (statusCode int, response interface{}) {
	validate := validator.New()
	var newUser dtos.UpdateUser

	context.BindJSON(&newUser)

	if err := validate.Struct(newUser); err != nil {
		validatorError := err.(validator.ValidationErrors).Error()
		errorMessage := "Cannot update user, required fields not provided\n" + validatorError
		return http.StatusBadRequest, errorMessage
	}

	t, _ := jwt.Parse(newUser.Token, nil)
	encodedId := t.Claims.(jwt.MapClaims)["id"]
	requesterId := fmt.Sprintf("%v", encodedId)

	if requesterId != context.Param("id") {
		return http.StatusUnauthorized, "Not enough permissions"
	}

	if permissions := r.HavePermissions(newUser); permissions != true {
		return http.StatusUnauthorized, "Not enough permissions"
	}

	parsedMenu := []models.Product{}
	for _, product := range newUser.Menu {
		newProduct := models.Product{
			Id:       product.Id,
			Category: product.Category,
			Name:     product.Name,
			Price:    product.Price,
		}
		parsedMenu = append(parsedMenu, newProduct)
	}
	newUser.Menu = parsedMenu

	password, err := middlewares.HashPassword(newUser.Password)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	id, err := primitive.ObjectIDFromHex(context.Param("id"))
	if err != nil {
		errorMessage := "Bad request, " + context.Param("id") + " is not a valid ID"
		return http.StatusBadRequest, errorMessage
	}

	filter := bson.M{"id": id}
	update := bson.M{
		"$set": bson.M{"name": newUser.Name, "email": newUser.Email,
			"password": password, "phone": newUser.Phone,
			"address": newUser.Address, "menu": newUser.Menu},
	}

	if _, err := r.users.UpdateOne(context, filter, update); err != nil {
		return http.StatusBadRequest, err.Error()
	}

	var updatedUser models.User

	userFilter := bson.M{"id": id}
	if err := r.users.FindOne(context, userFilter).Decode(&updatedUser); err != nil {
		return http.StatusNotFound, err.Error()
	}

	return http.StatusOK, gin.H{"message": "Updated successfully", "user": services.UserToDomain(updatedUser)}
}

// DELETE - http://localhost:3000/profile/:id
func (r *MongoUsersRepository) DeleteAccount(context *gin.Context) (statusCode int, response interface{}) {
	token := context.Query("token")

	t, _ := jwt.Parse(token, nil)
	encodedId := t.Claims.(jwt.MapClaims)["id"]
	requesterId := fmt.Sprintf("%v", encodedId)

	if requesterId != context.Param("id") {
		return http.StatusUnauthorized, gin.H{"error": "Not enough permissions"}
	}

	id, err := primitive.ObjectIDFromHex(context.Param("id"))
	if err != nil {
		errorMessage := "Bad request, " + context.Param("id") + " is not a valid ID"
		return http.StatusBadRequest, gin.H{"error": errorMessage}
	}

	filter := bson.M{"id": id}
	result := r.users.FindOneAndDelete(context, filter)
	if result.Err() != nil {
		return http.StatusBadRequest, gin.H{"error": result.Err().Error()}
	}

	return http.StatusOK, id.Hex()
}

// GET - http://localhost:3000/products
func (r *MongoUsersRepository) FindProducts(context *gin.Context) (statusCode int, response interface{}) {
	token := context.Query("token")

	t, _ := jwt.Parse(token, nil)
	encodedId := t.Claims.(jwt.MapClaims)["id"]
	requesterId := fmt.Sprintf("%v", encodedId)

	id, err := primitive.ObjectIDFromHex(requesterId)
	if err != nil {
		return http.StatusInternalServerError, gin.H{"error": "Not enough permissions"}
	}

	var restaurant models.User

	filter := bson.M{"role": "restaurant", "id": id}
	if err := r.users.FindOne(context, filter).Decode(&restaurant); err != nil {
		return http.StatusNotFound, gin.H{"error": "Restaurant not found"}
	}

	products := restaurant.Menu

	return http.StatusOK, products
}

// POST - http://localhost:3000/products
func (r *MongoUsersRepository) CreateProduct(context *gin.Context) (statusCode int, response interface{}) {
	var prod dtos.CreateOrUpdateProduct

	context.BindJSON(&prod)

	validate := validator.New()
	if err := validate.Struct(prod); err != nil {
		validatorError := err.(validator.ValidationErrors).Error()
		errorMessage := "Cannot create product, required fields not provided\n" + validatorError
		return http.StatusBadRequest, gin.H{"error": errorMessage}
	}

	t, _ := jwt.Parse(prod.Token, nil)
	encodedId := t.Claims.(jwt.MapClaims)["id"]
	requesterId := fmt.Sprintf("%v", encodedId)

	newProduct := &models.Product{
		Id:       primitive.NewObjectID(),
		Category: prod.Category,
		Name:     prod.Name,
		Price:    prod.Price,
	}

	var restaurant models.User

	id, err := primitive.ObjectIDFromHex(requesterId)
	if err != nil {
		errorMessage := "Bad request, " + requesterId + " is not a valid ID"
		return http.StatusBadRequest, errorMessage
	}

	filterFind := bson.M{"role": "restaurant", "id": id}
	if err := r.users.FindOne(context, filterFind).Decode(&restaurant); err != nil {
		return http.StatusNotFound, gin.H{"error": "Restaurant not found"}
	}

	restaurant.Menu = append(restaurant.Menu, *newProduct)

	filter := bson.M{"id": bson.M{"$eq": restaurant.Id}}
	update := bson.M{
		"$set": bson.M{"menu": restaurant.Menu},
	}

	updateResult, err := r.users.UpdateOne(context, filter, update)
	if err != nil {
		return http.StatusBadRequest, gin.H{"error": err.Error()}
	}

	return http.StatusOK, gin.H{"success": updateResult}
}

// PUT - http://localhost:3000/products/:id
func (r *MongoUsersRepository) UpdateProduct(context *gin.Context) (statusCode int, response interface{}) {
	var prod dtos.CreateOrUpdateProduct

	context.BindJSON(&prod)

	validate := validator.New()
	if err := validate.Struct(prod); err != nil {
		validatorError := err.(validator.ValidationErrors).Error()
		errorMessage := "Cannot update product, required fields not provided\n" + validatorError
		return http.StatusBadRequest, gin.H{"error": errorMessage}
	}

	t, _ := jwt.Parse(prod.Token, nil)
	encodedId := t.Claims.(jwt.MapClaims)["id"]
	requesterId := fmt.Sprintf("%v", encodedId)

	prodId, err := primitive.ObjectIDFromHex(context.Param("id"))
	if err != nil {
		errorMessage := "Bad request, " + context.Param("id") + " is not a valid ID"
		return http.StatusBadRequest, gin.H{"error": errorMessage}
	}

	newProduct := &models.Product{
		Id:       prodId,
		Category: prod.Category,
		Name:     prod.Name,
		Price:    prod.Price,
	}

	var restaurant models.User

	id, err := primitive.ObjectIDFromHex(requesterId)
	if err != nil {
		errorMessage := "Bad request, " + requesterId + " is not a valid ID"
		return http.StatusBadRequest, errorMessage
	}

	filterFind := bson.M{"role": "restaurant", "id": id}
	if err := r.users.FindOne(context, filterFind).Decode(&restaurant); err != nil {
		return http.StatusNotFound, gin.H{"error": "Restaurant not found"}
	}

	// update of the specific product
	for i, product := range restaurant.Menu {
		if product.Id == prodId {
			restaurant.Menu[i] = *newProduct
			break
		}
	}

	filter := bson.M{"id": bson.M{"$eq": restaurant.Id}}
	update := bson.M{
		"$set": bson.M{"menu": restaurant.Menu},
	}

	if _, err := r.users.UpdateOne(context, filter, update); err != nil {
		return http.StatusBadRequest, gin.H{"error": err.Error()}
	}

	return http.StatusOK, gin.H{"success": "Restaurant " + restaurant.Id.Hex() + " updated"}
}
