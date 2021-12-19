package repository

import (
	"comiditapp/api/dtos"
	"comiditapp/api/middlewares"
	"comiditapp/api/models"
	"comiditapp/api/services"
	"context"
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

func (r *MongoUsersRepository) DoesUserExists(u models.User) bool {
	// check if exists any user with same email, cause emails must be unique
	var result bson.M

	filter := bson.M{"email": u.Email}
	if err := r.users.FindOne(context.TODO(), filter).Decode(&result); err == nil {
		return true
	}
	return false
}

// POST - http://localhost:3000/auth/signup
func (r *MongoUsersRepository) SignUpUser(user models.User) error {
	_, err := r.users.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}

// POST - http://localhost:3000/auth/signin
func (r *MongoUsersRepository) SignInUser(context *gin.Context) (statusCode int, response interface{}) {

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

	expirationTime := time.Now().Add(time.Hour * 8760)
	token, err := services.GenerateJWT(dbUser.Email, dbUser.Id.Hex(), string(dbUser.Role), expirationTime.Unix())
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

	return http.StatusOK, gin.H{"token": token, "user": services.UserToDomain(dbUser)}
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

	expirationTime := time.Now().Add(time.Hour * 8760)
	token, err := services.GenerateJWT(restaurants[0].Email, restaurants[0].Id.Hex(), string(restaurants[0].Role), expirationTime.Unix())
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}

	c := &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		Expires:  expirationTime,
		SameSite: http.SameSiteLaxMode,
		Secure:   true,
	}

	// Con esta cookie ya podemos limitar las acciones segun quien tenga cookie o no
	http.SetCookie(context.Writer, c)

	// Con esta engancho la cookie a los headers
	context.Writer.Header().Add("Set-Cookie", c.String())

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
	validate := validator.New()
	var newUser models.User

	context.BindJSON(&newUser)

	// Checking permissions
	c, err := context.Cookie("token")
	if err != nil {
		return http.StatusUnauthorized, "Not enough permissions"
	}

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
			"password": newUser.Password, "phone": newUser.Phone,
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
		return http.StatusBadRequest, gin.H{"error": errorMessage}
	}

	// Checking permissions
	c, err := context.Cookie("token")
	if err != nil {
		return http.StatusUnauthorized, gin.H{"error": "Not enough permissions"}
	}

	t, _ := jwt.Parse(c, nil)
	encodedId := t.Claims.(jwt.MapClaims)["id"]
	requesterId := fmt.Sprintf("%v", encodedId)

	if requesterId != context.Param("id") {
		return http.StatusUnauthorized, gin.H{"error": "Not enough permissions"}
	}

	filter := bson.M{"id": bson.M{"$eq": id}}
	result := r.users.FindOneAndDelete(context, filter)
	if result.Err() != nil {
		return http.StatusBadRequest, gin.H{"error": result.Err().Error()}
	}

	return http.StatusOK, id.Hex()
}

// restaurant_role methods
// GET - http://localhost:3000/products
func (r *MongoUsersRepository) FindProducts(context *gin.Context) (statusCode int, response interface{}) {
	var restaurant models.User

	// Checking permissions
	c, err := context.Cookie("token")
	if err != nil {
		return http.StatusInternalServerError, gin.H{"error": "Not logged in"}
	}

	t, _ := jwt.Parse(c, nil)
	encodedId := t.Claims.(jwt.MapClaims)["id"]
	requesterId := fmt.Sprintf("%v", encodedId)

	id, err := primitive.ObjectIDFromHex(requesterId)
	if err != nil {
		return http.StatusInternalServerError, gin.H{"error": "Not enough permissions"}
	}

	filter := bson.M{"role": "restaurant", "id": id}
	if err := r.users.FindOne(context, filter).Decode(&restaurant); err != nil {
		return http.StatusNotFound, gin.H{"error": "Restaurant not found"}
	}

	products := restaurant.Menu

	return http.StatusOK, products
}

// POST - http://localhost:3000/products
func (r *MongoUsersRepository) CreateProduct(context *gin.Context) (statusCode int, response interface{}) {
	var prod models.Product

	context.BindJSON(&prod)

	// Checking permissions
	c, err := context.Cookie("token")
	if err != nil {
		return http.StatusUnauthorized, "Not enough permissions"
	}

	t, _ := jwt.Parse(c, nil)
	encodedId := t.Claims.(jwt.MapClaims)["id"]
	requesterId := fmt.Sprintf("%v", encodedId)

	validate := validator.New()
	if err := validate.Struct(prod); err != nil {
		validatorError := err.(validator.ValidationErrors).Error()
		errorMessage := "Cannot create product, required fields not provided\n" + validatorError
		return http.StatusBadRequest, gin.H{"error": errorMessage}
	}

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
	var prod models.Product

	context.BindJSON(&prod)

	// Checking permissions
	c, err := context.Cookie("token")
	if err != nil {
		return http.StatusUnauthorized, "Not enough permissions"
	}

	t, _ := jwt.Parse(c, nil)
	encodedId := t.Claims.(jwt.MapClaims)["id"]
	requesterId := fmt.Sprintf("%v", encodedId)

	validate := validator.New()
	if err := validate.Struct(prod); err != nil {
		validatorError := err.(validator.ValidationErrors).Error()
		errorMessage := "Cannot update product, required fields not provided\n" + validatorError
		return http.StatusBadRequest, gin.H{"error": errorMessage}
	}

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
