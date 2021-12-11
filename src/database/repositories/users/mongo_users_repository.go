package repository

import (
	"comiditapp/api/middlewares"
	"comiditapp/api/models"
	"fmt"
	"net/http"
	"time"

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

// any_role methods

// POST - http://localhost:3000/auth/signup
func (r *MongoUsersRepository) SignUpUser(context *gin.Context) (statusCode int, response interface{}) {
	var validate *validator.Validate = validator.New()
	var newUser models.User

	context.BindJSON(&newUser)

	err := validate.Struct(newUser)
	if err != nil {
		validatorError := err.(validator.ValidationErrors).Error()
		errorMessage := "Cannot create user, required fields not provided\n" + validatorError
		return http.StatusBadRequest, errorMessage
	}

	parsedMenu := []models.Product{}
	for _, product := range newUser.Menu {
		newProduct := models.Product{
			Id:       primitive.NewObjectID(),
			Category: product.Category,
			Name:     product.Name,
			Price:    product.Price,
		}
		parsedMenu = append(parsedMenu, newProduct)
	}

	newId := primitive.NewObjectID()
	newPassword, err := middlewares.HashPassword(newUser.HashedPassword)
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}

	user := &models.User{
		Id:             newId,
		Role:           newUser.Role,
		Name:           newUser.Name,
		Email:          newUser.Email,
		HashedPassword: newPassword,
		Phone:          newUser.Phone,
		Address:        newUser.Address,
		Menu:           parsedMenu,
	}

	if _, err := r.users.InsertOne(context, user); err != nil {
		return http.StatusBadRequest, err.Error()
	}

	// Tras crear el usuario se le genera un jwt y se setea en sus cookies
	expirationTime := time.Now().Add(time.Hour * 8760)
	token, err := middlewares.GenerateJWT(newUser.Email, newUser.Id.Hex(), string(newUser.Role), expirationTime.Unix())
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}

	c := &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: expirationTime,
	}
	http.SetCookie(context.Writer, c)
	context.Request.Header.Add("Set-Cookie", c.String())

	return http.StatusCreated, newId.Hex()
}

// POST - http://localhost:3000/auth/signin
func (r *MongoUsersRepository) SignInUser(context *gin.Context) (statusCode int, response interface{}) {

	var u models.User
	var dbUser models.User

	context.BindJSON(&u)

	filter := bson.M{"email": u.Email}
	if err := r.users.FindOne(context, filter).Decode(&dbUser); err != nil {
		return http.StatusNotFound, err.Error()
	}

	// Bcrypt se encarga de hashear la del user y compararla con la de db
	if err := middlewares.VerifyPassword(u.HashedPassword, dbUser.HashedPassword); err != nil {
		return http.StatusUnauthorized, "Login failed, username or password are incorrect"
	}

	expirationTime := time.Now().Add(time.Hour * 8760)
	token, err := middlewares.GenerateJWT(dbUser.Email, dbUser.Id.Hex(), string(dbUser.Role), expirationTime.Unix())
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}

	c := &http.Cookie{
		Name:    "token",
		Value:   token,
		Path:    "/",
		Expires: expirationTime,
	}
	http.SetCookie(context.Writer, c)
	context.Request.Header.Add("Set-Cookie", c.String())

	return http.StatusOK, token
}

// POST - http://localhost:3000/auth/signOutUser
func (r *MongoUsersRepository) SignOutUser(context *gin.Context) (statusCode int, response interface{}) {
	c := &http.Cookie{
		Name:    "token",
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0),
	}

	http.SetCookie(context.Writer, c)
	context.Request.Header.Del("Set-Cookie")

	return http.StatusOK, gin.H{"message": "Successfully logged out"}
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
	var validate *validator.Validate = validator.New()
	var newUser models.User

	context.BindJSON(&newUser)

	c, err := context.Cookie("token")
	if err != nil {
		return http.StatusUnauthorized, "Not enough permissions"
	}

	// Checking permissions
	t, _ := jwt.Parse(c, nil)
	encodedId := t.Claims.(jwt.MapClaims)["id"]
	requesterId := fmt.Sprintf("%v", encodedId)

	if requesterId != context.Param("id") {
		return http.StatusUnauthorized, "Not enough permissions"
	}

	if err := validate.Struct(newUser); err != nil {
		validatorError := err.(validator.ValidationErrors).Error()
		errorMessage := "Cannot update user, required fields not provided\n" + validatorError
		return http.StatusBadRequest, errorMessage
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

	id, err := primitive.ObjectIDFromHex(context.Param("id"))
	if err != nil {
		errorMessage := "Bad request, " + context.Param("id") + " is not a valid ID"
		return http.StatusBadRequest, errorMessage
	}

	filter := bson.M{"id": bson.M{"$eq": id}}
	update := bson.M{
		"$set": bson.M{"role": newUser.Role, "name": newUser.Name, "email": newUser.Email,
			"hashedPassword": newUser.HashedPassword, "phone": newUser.Phone,
			"address": newUser.Address, "menu": parsedMenu},
	}

	if _, err := r.users.UpdateOne(context, filter, update); err != nil {
		return http.StatusBadRequest, err.Error()
	}

	return http.StatusOK, id.Hex()
}

// DELETE - http://localhost:3000/profile/:id
func (r *MongoUsersRepository) DeleteAccount(context *gin.Context) (statusCode int, response interface{}) {
	id, err := primitive.ObjectIDFromHex(context.Param("id"))
	if err != nil {
		errorMessage := "Bad request, " + context.Param("id") + " is not a valid ID"
		return http.StatusBadRequest, errorMessage
	}

	filter := bson.M{"id": bson.M{"$eq": id}}
	result := r.users.FindOneAndDelete(context, filter)
	if result.Err() != nil {
		return http.StatusBadRequest, result.Err().Error()
	}

	return http.StatusOK, id.Hex()
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
