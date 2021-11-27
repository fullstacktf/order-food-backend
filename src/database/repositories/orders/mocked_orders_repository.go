package repository

import (
	"comiditapp/api/enums"
	"comiditapp/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

var orders []models.Order = []models.Order{
	{Id: "0", RestaurantId: "1", ClientId: "1", Status: enums.Ordered, TotalPrice: 30.30, Products: []models.ProductInfo{{ProductId: "1", Quantity: 1}}},
	{Id: "1", RestaurantId: "1", ClientId: "2", Status: enums.Preparing, TotalPrice: 45.30, Products: []models.ProductInfo{{ProductId: "3", Quantity: 2}, {ProductId: "1", Quantity: 2}}},
}

type MockedOrdersRepository struct {
	mock.Mock
}

// any_role methods
func (r *MockedOrdersRepository) FindOrders(context *gin.Context) []models.Order {
	context.IndentedJSON(http.StatusOK, orders)

	// Son necesarias estas dos lineas en cada metodo ??
	// En caso de que si, que devolvemos en las funciones que se supone que devolverian void??
	args := r.Called(context)
	return args.Get(0).([]models.Order)
}

func (r *MockedOrdersRepository) GetOrderById(context *gin.Context) models.Order {
	args := r.Called(context)
	return args.Get(0).(models.Order)
}
func (r *MockedOrdersRepository) CreateOrder(context *gin.Context) models.Order {
	args := r.Called(context)
	return args.Get(0).(models.Order)
}

// restaurant_role methods
func (r *MockedOrdersRepository) UpdateClientOrder(context *gin.Context) models.Order {
	args := r.Called(context)
	return args.Get(0).(models.Order)
}
func (r *MockedOrdersRepository) FindClientOrders(context *gin.Context) []models.Order {
	args := r.Called(context)
	return args.Get(0).([]models.Order)
}
