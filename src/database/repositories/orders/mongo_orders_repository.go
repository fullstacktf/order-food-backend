package repository

import (
	"comiditapp/api/dtos"
	"comiditapp/api/enums"
	"comiditapp/api/models"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoOrdersRepository struct {
	orders *mongo.Collection
}

func NewMongoOrdersRepository(db *mongo.Database) *MongoOrdersRepository {
	return &MongoOrdersRepository{orders: db.Collection("orders")}
}

// any_role methods

// GET - http://localhost:3000/profile/orders
func (r *MongoOrdersRepository) FindOrders(context *gin.Context) (statusCode int, response interface{}) {
	// Checking permissions
	var userToken dtos.UserToken

	context.BindJSON(&userToken)

	validate := validator.New()
	if err := validate.Struct(userToken); err != nil {
		validatorError := err.(validator.ValidationErrors).Error()
		errorMessage := "Cannot get orders, required fields not provided\n" + validatorError
		return http.StatusBadRequest, gin.H{"error": errorMessage}
	}

	t, _ := jwt.Parse(userToken.Token, nil)
	encodedId := t.Claims.(jwt.MapClaims)["id"]
	requesterId := fmt.Sprintf("%v", encodedId)

	id, err := primitive.ObjectIDFromHex(requesterId)
	if err != nil {
		return http.StatusUnauthorized, gin.H{"error": "Not enough permissions"}
	}

	filter := bson.M{"clientId": id}
	foundOrders, err := r.orders.Find(context, filter)
	if err != nil {
		return http.StatusConflict, gin.H{"error": err.Error()}
	}

	orders := []*models.Order{}
	if err := foundOrders.All(context, &orders); err != nil {
		return http.StatusConflict, gin.H{"error": err.Error()}
	}

	return http.StatusOK, orders
}

// GET - http://localhost:3000/profile/orders/:id
func (r *MongoOrdersRepository) GetOrderById(context *gin.Context) (statusCode int, response interface{}) {
	// Checking permissions
	var userToken dtos.UserToken

	context.BindJSON(&userToken)

	validate := validator.New()
	if err := validate.Struct(userToken); err != nil {
		validatorError := err.(validator.ValidationErrors).Error()
		errorMessage := "Cannot get orders, required fields not provided\n" + validatorError
		return http.StatusBadRequest, gin.H{"error": errorMessage}
	}

	t, _ := jwt.Parse(userToken.Token, nil)
	encodedId := t.Claims.(jwt.MapClaims)["id"]
	requesterId := fmt.Sprintf("%v", encodedId)

	clientId, err := primitive.ObjectIDFromHex(requesterId)
	if err != nil {
		return http.StatusUnauthorized, gin.H{"error": "Not enough permissions"}
	}

	id, err := primitive.ObjectIDFromHex(context.Param("id"))
	if err != nil {
		errorMessage := "Bad request, " + context.Param("id") + " is not a valid ID"
		return http.StatusBadRequest, gin.H{"error": errorMessage}
	}

	var foundOrder models.Order

	filter := bson.M{"clientId": clientId, "id": id}
	if err := r.orders.FindOne(context, filter).Decode(&foundOrder); err != nil {
		return http.StatusNotFound, gin.H{"error": "Order not found"}
	}

	return http.StatusOK, foundOrder
}

// POST - http://localhost:3000/restaurants/:id/order
func (r *MongoOrdersRepository) CreateOrder(context *gin.Context) (statusCode int, response interface{}) {
	validate := validator.New()
	var newOrder dtos.OrderCreate

	context.BindJSON(&newOrder)

	if err := validate.Struct(newOrder); err != nil {
		validatorError := err.(validator.ValidationErrors).Error()
		errorMessage := "Cannot create order, required fields not provided\n" + validatorError
		return http.StatusBadRequest, gin.H{"error": errorMessage}
	}

	t, _ := jwt.Parse(newOrder.Token, nil)
	encodedId := t.Claims.(jwt.MapClaims)["id"]
	requesterId := fmt.Sprintf("%v", encodedId)

	clientId, err := primitive.ObjectIDFromHex(requesterId)
	if err != nil {
		return http.StatusUnauthorized, gin.H{"error": "Not enough permissions"}
	}

	restaurantId, err := primitive.ObjectIDFromHex(context.Param("id"))
	if err != nil {
		errorMessage := "Bad request, " + context.Param("id") + " is not a valid ID"
		return http.StatusBadRequest, gin.H{"error": errorMessage}
	}

	var totalPrice float64 = 0.0

	for _, orderProduct := range newOrder.Products {
		totalPrice += orderProduct.Price * float64(orderProduct.Quantity)
	}

	id := primitive.NewObjectID()
	order := &models.Order{
		Id:           id,
		ClientId:     clientId,
		RestaurantId: restaurantId,
		Status:       enums.Ordered,
		TotalPrice:   totalPrice,
		Products:     newOrder.Products,
	}

	if _, err := r.orders.InsertOne(context, order); err != nil {
		return http.StatusBadRequest, err.Error()
	}

	return http.StatusCreated, gin.H{"success": "Order " + id.Hex() + " created"}
}

// restaurant_role methods

// PUT - http://localhost:3000/orders/:id
func (r *MongoOrdersRepository) UpdateClientOrder(context *gin.Context) (statusCode int, response interface{}) {
	validate := validator.New()
	var newOrder dtos.OrderUpdate

	context.BindJSON(&newOrder)

	if err := validate.Struct(newOrder); err != nil {
		validatorError := err.(validator.ValidationErrors).Error()
		errorMessage := "Cannot update order, required fields not provided\n" + validatorError
		return http.StatusBadRequest, gin.H{"error": errorMessage}
	}

	t, _ := jwt.Parse(newOrder.Token, nil)
	encodedId := t.Claims.(jwt.MapClaims)["id"]
	requesterId := fmt.Sprintf("%v", encodedId)

	restaurantId, err := primitive.ObjectIDFromHex(requesterId)
	if err != nil {
		return http.StatusUnauthorized, gin.H{"error": "Not enough permissions"}
	}

	orderId, err := primitive.ObjectIDFromHex(context.Param("id"))
	if err != nil {
		errorMessage := "Bad request, " + context.Param("id") + " is not a valid ID"
		return http.StatusBadRequest, gin.H{"error": errorMessage}
	}

	filter := bson.M{"id": orderId, "restaurantId": restaurantId}
	update := bson.M{"$set": bson.M{"status": newOrder.Status}}

	result, err := r.orders.UpdateOne(context, filter, update)
	if result.MatchedCount == 0 {
		return http.StatusUnauthorized, gin.H{"error": "Not enough permissions"}
	}
	if err != nil {
		return http.StatusInternalServerError, gin.H{"error": err.Error()}
	}

	return http.StatusOK, gin.H{"success": "Order " + orderId.Hex() + " updated"}
}

// GET - http://localhost:3000/orders
func (r *MongoOrdersRepository) FindClientOrders(context *gin.Context) (statusCode int, response interface{}) {
	// Checking permissions
	var userToken dtos.UserToken

	context.BindJSON(&userToken)

	validate := validator.New()
	if err := validate.Struct(userToken); err != nil {
		validatorError := err.(validator.ValidationErrors).Error()
		errorMessage := "Cannot get orders, required fields not provided\n" + validatorError
		return http.StatusBadRequest, gin.H{"error": errorMessage}
	}

	t, _ := jwt.Parse(userToken.Token, nil)
	encodedId := t.Claims.(jwt.MapClaims)["id"]
	requesterId := fmt.Sprintf("%v", encodedId)

	restaurantId, err := primitive.ObjectIDFromHex(requesterId)
	if err != nil {
		return http.StatusUnauthorized, gin.H{"error": "Not enough permissions"}
	}

	filter := bson.M{"restaurantId": restaurantId}
	foundOrders, err := r.orders.Find(context, filter)
	if err != nil {
		return http.StatusConflict, gin.H{"error": err.Error()}
	}

	orders := []*models.Order{}
	if err := foundOrders.All(context, &orders); err != nil {
		return http.StatusConflict, gin.H{"error": err.Error()}
	}

	return http.StatusOK, orders
}
