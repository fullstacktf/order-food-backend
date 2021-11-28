package repository

import (
	"comiditapp/api/enums"
	"comiditapp/api/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var id1 = primitive.NewObjectID()
var id2 = primitive.NewObjectID()
var id3 = primitive.NewObjectID()
var id4 = primitive.NewObjectID()
var orders []models.Order = []models.Order{
	{Id: id1, RestaurantId: primitive.NewObjectID(), ClientId: primitive.NewObjectID(), Status: enums.Ordered, TotalPrice: 30.30, Products: []models.ProductInfo{{ProductId: primitive.NewObjectID(), Quantity: 1}}},
	{Id: id2, RestaurantId: primitive.NewObjectID(), ClientId: primitive.NewObjectID(), Status: enums.Preparing, TotalPrice: 45.30, Products: []models.ProductInfo{{ProductId: primitive.NewObjectID(), Quantity: 2}, {ProductId: primitive.NewObjectID(), Quantity: 2}}},
	{Id: id3, RestaurantId: primitive.NewObjectID(), ClientId: primitive.NewObjectID(), Status: enums.Preparing, TotalPrice: 45.30, Products: []models.ProductInfo{{ProductId: primitive.NewObjectID(), Quantity: 2}, {ProductId: primitive.NewObjectID(), Quantity: 2}}},
	{Id: id4, RestaurantId: primitive.NewObjectID(), ClientId: primitive.NewObjectID(), Status: enums.Preparing, TotalPrice: 45.30, Products: []models.ProductInfo{{ProductId: primitive.NewObjectID(), Quantity: 2}, {ProductId: primitive.NewObjectID(), Quantity: 2}}},
}

type MockedOrdersRepository struct {
	mock.Mock
}

// any_role methods
func (r *MockedOrdersRepository) GetOrders(context *gin.Context) []models.Order {
	// En principio estaria asi, ya veremos en un futuro si los filtros los hacemos a nivel de back o de front
	context.IndentedJSON(http.StatusOK, orders)

	// Son necesarias estas dos lineas en cada metodo ??
	// En caso de que si, que devolvemos en las funciones que se supone que devolverian void??
	args := r.Called(context)
	return args.Get(0).([]models.Order)
}

func (r *MockedOrdersRepository) GetOrderById(context *gin.Context) models.Order {
	result := []models.Order{}

	orderId, err := primitive.ObjectIDFromHex(context.Param("id"))

	if err != nil {
		errorMsg := "Bad Request, " + err.Error()

		http.Error(context.Writer, errorMsg, http.StatusBadRequest)
		panic(err)
	}

	for _, value := range orders {
		if value.Id == orderId {
			result = append(result, value)
			break
		}
	}

	if len(result) == 0 {
		errorMsg := "Cannot found order with ID " + orderId.Hex()
		http.Error(context.Writer, errorMsg, http.StatusNotFound)
		panic(errorMsg)
	}

	context.IndentedJSON(http.StatusOK, result)

	args := r.Called(context)
	return args.Get(0).(models.Order)
}

func (r *MockedOrdersRepository) CreateOrder(context *gin.Context) models.Order {

	var validate *validator.Validate = validator.New()
	var newOrder models.Order

	json.NewDecoder(context.Request.Body).Decode(&newOrder)

	err := validate.Struct(newOrder)
	if err != nil {
		errorMsg := "Cannot create order, required fields not provided...\n" + err.(validator.ValidationErrors).Error()

		http.Error(context.Writer, errorMsg, http.StatusBadRequest)
		panic(err)
	}

	for _, value := range orders {
		if value.Id == newOrder.Id {
			errorMsg := "Cannot create order, ID " + value.Id.Hex() + " already exists"
			http.Error(context.Writer, errorMsg, http.StatusBadRequest)
			panic(errorMsg)
		}
	}

	orders = append(orders, newOrder)
	context.IndentedJSON(http.StatusOK, newOrder)

	args := r.Called(context)
	return args.Get(0).(models.Order)
}

// restaurant_role methods
func (r *MockedOrdersRepository) UpdateClientOrder(context *gin.Context) models.Order {

	// Controlar que SOLO puede modificar el campo de status, si intenta cambiar otro tendremos que ignorarlo
	// Controlar que tiene rol de restaurante

	orderId, err := primitive.ObjectIDFromHex(context.Param("id"))

	if err != nil {
		errorMsg := "Bad Request, " + err.Error()

		http.Error(context.Writer, errorMsg, http.StatusBadRequest)
		panic(err)
	}

	var validate *validator.Validate = validator.New()
	var newOrder models.Order

	foundOrderIndex := -1
	for index, value := range orders {
		if value.Id == orderId {
			foundOrderIndex = index
		}
	}

	if foundOrderIndex == -1 {
		errorMsg := "Cannot found order with ID " + orderId.Hex()
		http.Error(context.Writer, errorMsg, http.StatusNotFound)
		panic(errorMsg)
	}

	json.NewDecoder(context.Request.Body).Decode(&newOrder)

	valErr := validate.Struct(newOrder)
	if valErr != nil {
		errorMsg := "Cannot update order, required fields not provided...\n" + valErr.(validator.ValidationErrors).Error()

		http.Error(context.Writer, errorMsg, http.StatusBadRequest)
		panic(valErr)
	}

	orders[foundOrderIndex] = newOrder
	context.IndentedJSON(http.StatusOK, newOrder)

	args := r.Called(context)
	return args.Get(0).(models.Order)
}

func (r *MockedOrdersRepository) GetClientsOrders(context *gin.Context) []models.Order {

	// Controlar que tiene rol de restaurante

	// Este id deberiamos sacarlos de la sesion del usuario, comprobando que es un restaurante y sacando su id

	// En este caso nunca va a encontrar el restaurante precisamente por lo anterior que digo,
	// pero una vez podamos enganchar el id, esto funcionará
	restaurantId := primitive.NewObjectID()

	result := []models.Order{}

	for _, value := range orders {
		if value.RestaurantId == restaurantId {
			result = append(result, value)
		}
	}

	if len(result) == 0 {
		errorMsg := "Cannot found orders for restaurant with ID " + restaurantId.Hex()
		http.Error(context.Writer, errorMsg, http.StatusNotFound)
		panic(errorMsg)
	}

	context.IndentedJSON(http.StatusOK, result)

	args := r.Called(context)
	return args.Get(0).([]models.Order)
}
